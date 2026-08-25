package arrays

/*
Problem:
Ein sortiertes Array in-place komprimieren. Jeder Wert darf höchstens zweimal vorkommen. Nur nums[:k] ist relevant.

Muster:
- Stable In-place Compaction
- Read-Write-Pointer
- Begrenzte Häufigkeit in sortierten Daten

Invariante:
nums[:write] enthält das korrekte Ergebnis für alle bisher von read untersuchten Elemente.
*/
func RemoveDuplicatesFromSortedArrayII(nums []int) int {
	write := 0

	if len(nums) == 0 {
		return write
	}

	// Um für den eigentlichen Algorithmus das Inbounds-Verhalten am Start zu garantieren, muss nums mindestens zwei Elemente besitzen
	if len(nums) == 1 {
		return 1
	}

	// 1,1,1,2,2,3
	// Jeder Wert muss einmal gelesen werden - nicht mehrfach!
	for read := 0; read < len(nums); read++ {
		// 1 Sicherheitsbedingung, sodass anfangs kein out-of-bounds-Zugriff passiert
		// 2 wenn dann das element dasselbe wie zwei zuvor ist: nicht übernehmen
		if read >= 2 && nums[read] == nums[write-2] {
			continue
		}
		nums[write] = nums[read]
		write++
	}

	return write
}
