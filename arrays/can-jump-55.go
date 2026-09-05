package arrays

// Ein Array oder eine Liste durchqueren, indem jede Zahl darin als Sprung aufgefasst wird. Beginn: auf Index 0. Ziel: auf Index n-1 (Ende der Liste). Frage: kommt man auf den letzten Index, springt man darüber, oder kommt man auf einen Endlos-Loop?
// Invariante: Nach Verarbeitung der bisherigen Indizes ist jumpReach der größte Index, der über irgendeinen bisherigen Weg erreichbar ist.

func CanJump(nums []int) bool {
	if len(nums) == 0 || (nums[0] == 0 && !(len(nums) == 1)) {
		return false
	}

	furthestReach := nums[0]

	for i := 0; i <= len(nums)-1; i++ {
		if furthestReach >= len(nums)-1 {
			return true
		}

		if i > furthestReach {
			return false
		}

		if nums[i]+i > furthestReach {
			furthestReach = nums[i] + i
		}
	}

	return false
}

// GPT-generated solution, compressed
func moreOptimalCanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	furthestReach := 0

	for i, jumpLength := range nums {

		// wir sind bei einer Position angekommen, die wir nicht hätten erreichen können
		if i > furthestReach {
			return false
		}

		// wenn eine Position die Reichweite erhöht, muss diese neue maximale Reichweite gespeichert werden
		if i+jumpLength > furthestReach {
			furthestReach = i + jumpLength
		}

		if furthestReach > (len(nums) - 1) {
			return true
		}
	}

	return false
}

func ___CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	jumpReach := 0

	for i := 0; i <= len(nums)-1 && i <= jumpReach; i++ {
		// schauen, welche indexe erreichbar sind
		if nums[i]+i > jumpReach {
			jumpReach = nums[i] + i
		}
	}

	if jumpReach >= len(nums)-1 {
		return true
	}

	return false
}

// Immer nur die geringste Menge springen, außer:
// negative Zahl oder 0, dann versuchen, darüber zu springen.
// DFS-Backtracking Ansatz: möglich, aber nicht optimal
func __CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	i := 0

	// Abbrechen, sobald man über die Boundary oder am Ziel ist
	for i < len(nums) && i != len(nums) {
		// Überprüfen, ob man am Ziel ist
		if i == len(nums)-1 {
			return true
		}

		// diese sequenz überspringt Zahlen, die das Fortkommen aufhalten würden
		j := 2
		if nums[i+1] <= 0 {

			for nums[i+1] <= 0 && j <= len(nums) {
				// Boundary check für diese Schleife
				if i+j >= len(nums)-1 {
					return false
				}

				// Ist dieser Index groß genug, um von da aus weiterzumachen?
				if nums[i+j] > 0 {
					i += j
					break
				}
			}
		} else {
			i += 1
		}
	}

	return false
}

func _CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	i := 0

	// Abbrechen, sobald man über die Boundary oder am Ziel ist
	for i < len(nums) && i != len(nums) {
		// Überprüfen, ob man am Ziel ist
		if i == len(nums)-1 {
			return true
		}

		// Überprüfen, ob der nächste Index noch in der Boundaries ist
		if i+nums[i] >= len(nums) {
			return false
		}

		// wenn eine negative Zahl genauso groß wie die aktuelle positive ist
		var targetValue = i + nums[i]
		if nums[i]+nums[targetValue] == 0 || nums[i] == 0 {
			return false
		}

		i += nums[i]
	}

	return false
}
