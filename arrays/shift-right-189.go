package arrays

import "fmt"

/*
Problem:
Den Wert finden, der häufiger als n/2 vorkommt.

Muster:
- Cyclic Array Transformation
- Index Mapping mit zusätzlichem Slice
- "Vollständige Operation k-mal wiederholen" -> wahrscheinlich vermeidbare O(n*k)-Lösung

Kernidee:
- zuerst normalisieren, um unnötige Rotationen zu sparen -> ab einer vollständigen Rotation wurde unnötige Arbeit verrichtet
*/
func RotateShiftRight(nums []int, amount int) {
	if amount == 0 || len(nums) < 2 {
		return
	}

	// normalisieren: bei 3 Elementen sind 9 Rotationen dasselbe wie 0 Rotationen
	amount %= len(nums)

	// Dieses Array hält das Ergebnis temporär
	array := make([]int, len(nums))

	// Anstatt so oft zu rotieren, wie die Aufgabe besagt, wird jedes Element in einem Durchlauf um "k" rotiert.
	for index, value := range nums {
		// Für jeden Wert in nums errechnen wir zuerst den Zielindex
		target := (index + amount) % len(nums) // (Position 1 und Rotation 2000) normalisiert durch Reste-Teilung der Gesamtmenge
		array[target] = value
	}

	copy(nums, array)
}

// Zu langsam! O(n)
func _RotateShiftRight(nums []int, amount int) {
	if amount == 0 || len(nums) < 2 {
		return
	}

	for range amount {
		tmp := nums[len(nums)-1]

		// von hinten nach vorne, weil sonst das erste element immer wieder kopiert würde
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}

		nums[0] = tmp
	}

	fmt.Print("")
}
