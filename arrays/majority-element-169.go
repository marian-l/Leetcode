package arrays

func MajorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Speichert für jeden Zahlenwert, wie oft er in nums vorkommt.
	// Nicht vorhandene Schlüssel liefern bei map[int]int automatisch 0.
	counts := make(map[int]int)

	// Idee: Beim Loop über das Array den Wert am aktuellen Index (zum Beispiel 3 in [3,2,3]) als Schlüssel nutzen und dort inkrementieren
	for read := 0; read < len(counts); read++ {
		counts[nums[read]] = counts[nums[read]] + 1
	}

	result := -1
	tmp := 0

	// Zum Auslesen braucht es die Schlüssel aus dem Array
	for i, v := range counts {
		if v == 0 {
			continue
		}

		if v > tmp {
			tmp = v
			result = i
		}
	}

	return result
}
