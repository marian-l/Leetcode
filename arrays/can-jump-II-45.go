package arrays

// [2,3,1,1,4]

func Jump(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	// Minimale Kosten
	jumps := 0

	if len(nums) == 1 {
		return jumps
	}

	// Das Ende der aktuellen Untersuchungsiteration / Untersuchungsbreite / Breitenlevel im Baum
	currentEnd := 0
	// Das Ende für die nächste Iteration und das Maß, ob wir fertig sind
	farthestReach := 0

	for i, jumpLength := range nums {
		// wir sind bei einer Position angekommen, die wir nicht hätten erreichen können (direkt nach der Indexinkrementation)
		if i > farthestReach {
			return -1
		}

		// wenn eine Position die Reichweite erhöht, muss diese neue maximale Reichweite gespeichert werden
		if i+jumpLength > farthestReach {
			farthestReach = i + jumpLength
		}

		// Das Ende einer Breite entspricht am Anfang 0, da wir immer auf Feld eins starten. 
		// Danach ist das Ende einer Breite die größte Distanz, die wir in der Untersuchung gefunden haben. 
		// Die Sprünge müssen inkrementiert werden und die nächste Iteration beginnt
		if i == currentEnd {
			// Am Ende einer Breite muss gesprungen werden, um die gefundene, maximale Distanz auch zurückzulegen
			jumps++

			// Am Ende einer Breite muss die Grenze der aktuellen Breite auf die gefundene Höchstgrenze gesetzt werden, damit die nächste Breite untersucht werden kann
			currentEnd = farthestReach

			// Am Ende einer Breite muss untersucht werden, ob das Ziel erreicht wurde
			if currentEnd >= len(nums)-1 {
				return jumps
			}
		}
	}

	return jumps
}
