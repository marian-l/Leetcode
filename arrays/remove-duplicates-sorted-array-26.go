package arrays

func RemoveDuplicatesFromSortedArray(nums []int) int {
	write := 0

	if len(nums) == 0 {
		return write
	}

	write = 1

	// Um für den eigentlichen Algorithmus das Inbounds-Verhalten am Start zu garantieren, muss nums mindestens zwei Elemente besitzen
	if len(nums) == 1 {
		return write
	}

	// Jeder Wert muss einmal gelesen werden - nicht mehrfach!
	for read := 0; read < len(nums); read++ {
		// Die gelesene Zahl wird mit der letzten geschriebenen Zahl verglichen. Das ist nur in einem sortierten Array gültig, da die Zahlen in einer vorhersehbaren Reihenfolge drankommen
		if nums[read] != nums[write-1] {
			nums[write] = nums[read]
			// Der Schreibkopf wandert ausschließlich weiter, wenn geschrieben wurde - alle anderen Bewegungen sind nicht nötig
			write++
		}
	}

	return write
}

func _RemoveDuplicatesFromSortedArray(nums []int) int {
	unique := 0

	if len(nums) == 0 {
		return unique
	}

	unique = 1

	// Um für den eigentlichen Algorithmus das Inbounds-Verhalten am Start zu garantieren, muss nums mindestens zwei Elemente besitzen
	if len(nums) == 1 {
		return unique
	}

	// um das array zu durchqueren
	for i := 0; i < len(nums)-1; i++ {

		// so lange shiften, bis eine neue Zahl auftaucht
		for nums[i] == nums[i+1] {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}

			// nach jedem shiften überprüfen, ob jetzt eine neue zahl vorliegt
			if nums[i] != nums[i+1] {
				unique++
			}
		}
	}
	// Fehler: Durch Shift Left werden hier keine Nullen erzeugt, wie es eigentlich sein sollte. Dadurch entsteht ein Endlosloop
	// Gleichzeitig ist aber ein Shift-Left mit Nullen auffüllen auch nur dann richtig, wenn das Array aufsteigend sortiert ist
	// Workaround: ans Ende des Arrays einen Delimiter wie Int.Max oder -1 stellen?

	return unique
}

/////////////// Array-Delimiter
//	activeLength := len(nums)
//
// 	for i := 1; i < activeLength; {
// 		if nums[i] == nums[i-1] {
// 			copy(nums[i:activeLength-1], nums[i+1:activeLength])
// 			activeLength--
// 		} else {
// 			i++
// 		}
// 	}
//
// 	return activeLength
