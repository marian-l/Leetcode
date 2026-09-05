package arrays

import "sort"

// Jeder Eintrag markiert die Anzahl von Zitationen für das jeweilige Paper einer Wissenschaftlerin.
// Errechne den H-Index

// H-Index ist der Wert höchste Wert H an Zitationen pro Paper, den eine Anzahl von H Paper hat
// 10 Paper, 1 Paper 1x zitiert, 9 Paper 9x zitiert: H-Index von 9

// Wahrscheinlich macht es Sinn, von hinten nach vorne zu suchen, damit das erste valide Ergebnis direkt das finale Ergebnis ist (suche zuerst den größten)

// Der H-Index kann nie höher sein als die Länge der Liste!

// Wenn man einmal eine Frequency-Map machen würde,
// deren Index die Anzahl Zitationen und
// deren Wert die Anzahl von Paper ist,

// Bei einer Liste von 10 Papern könnte man dann die Anzahl von Papern (Value) immer weiter aufsummieren, bis sie größer-gleich der Anzahl von Zitationen (Key) ist.

// Gibt es einen einfacheren Weg, die Vergangenheit zusammenzufassen?
// Maximal gesetzte Menge von Zitationen pro Paper ist 1000, maximale gesetzte Länge von Papern ist 5000

func HIndex(citations []int) int {
	listLength := len(citations)

	if listLength == 0 {
		return 0
	}

	// sortieren (absteigend), weil es einfach ist und die Map dann sortiert aufgebaut wird
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	// absteigend sortiertes slice (dynamische Länge möglich)
	citationPaperAmount := make([]int, citations[listLength-1]+1)

	// Frequency Map erstellen
	for _, citationAmount := range citations {
		citationPaperAmount[citationAmount] += 1
	}

	clear(citations)

	// Wie viele Paper haben bereits eine höhere Zitationsanzahl als die aktuelle Zitationsanzahl?
	currentPapersSum := 0

	// von hinten nach vorne in der Zitationsanzahl aufsummieren, bis die Paperanzahl übereinstimmt.
	for citationIndex := len(citationPaperAmount) - 1; citationIndex >= 0; citationIndex-- {
		if citationPaperAmount[citationIndex] == 0 {
			continue
		}

		// Ist die Anzahl von Papern, die einen Wert >= der aktuellen Zitationsanzahl haben, höher oder gleich der Anzahl von Zitationen pro Paper, dann haben wir es gefunden
		if (citationPaperAmount[citationIndex] + currentPapersSum) >= citationIndex {
			return max(citationIndex, currentPapersSum)
		}

		currentPapersSum += citationPaperAmount[citationIndex]
	}

	return len(citations)
}

// Reverse Ansätze
// citationPaperAmount = slices.Reverse(citationPaperAmount)
// for i := range citationPaperAmount
// citationPaperAmount[len(slice)-1-i]

//
//
//

// Hier habe ich den H-Index falsch verstanden:

// H-Index: Größter Wert H, wobei H sowohl die Anzahl der Veröffentlichungen (Länge der Liste) ist als auch die größte Menge von Zitationen, die JEDES Paper bekommen hat.
// 10 Paper, wovon eines 8 Zitationen hat und alle anderen 9: Der H-Index ist 8

// Es macht keinen Unterschied, von wo man sucht, weil man das Paper mit der geringsten Zitationsanzahl braucht, da es alle anderen zurückhält.

// Wahrscheinlich macht es aber auch Sinn, einmal durch die Liste zu loopen und den geringsten Wert ausfindig zu machen. Wenn dieser kleiner-gleich der Listenlänge ist, dann
func _HIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

	// Laut Aufgabe ist 1000 das gesetzte Maximum
	lowestCitations := 1001

	// Wir suchen eigentlich nur die Begrenzung durch den niedrigsten Wert von Zitationen
	for _, currentPaperCitations := range citations {
		if currentPaperCitations < lowestCitations {
			lowestCitations = currentPaperCitations
		}
	}

	// Wir geben den Wert zurück, der kleiner ist
	return min(len(citations), lowestCitations)
}
