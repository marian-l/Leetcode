package arrays

/*
Problem:
Maximiere prices[sell] - prices[buy], wobei buy < sell gelten muss.

Muster:
Ein Durchlauf mit laufendem Minimum und bisher bestem Ergebnis.

Kernidee:
Für einen festen Verkaufstag ist nur der günstigste Kaufpreis
aus allen vorherigen Tagen relevant. Die möglichen früheren
Kauftage werden dadurch in einem einzigen Wert zusammengefasst.

Invariante:
- cheapestBuyPrice ist der günstigste Preis vor dem aktuellen Tag.
- maxProfit ist der größte Gewinn aller bisher geprüften Verkäufe.

Erkenntnis:
Wenn eine innere Schleife immer wieder denselben bereits gelesenen
Bereich durchsucht, prüfen, ob sich dessen relevante Information
durch Minimum, Maximum, Summe oder Häufigkeit zusammenfassen lässt.
*/

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	
	cheapestBuyPrice := 8999999999999999999
	maxProfit := 0
	
	for day := 0; day < len(prices); day++ {
		if cheapestBuyPrice > prices[day] {
			cheapestBuyPrice = prices[day]
		} else if prices[day] - cheapestBuyPrice > maxProfit {
			maxProfit = prices[day] - cheapestBuyPrice
		}
	}

	return maxProfit
}

// funktioniert, ist aber zu langsam, skaliert nicht!
func _MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0

	// buy muss immer kleiner seine als sell
	for buy := 0; buy < len(prices); buy++ {
		for sell := buy + 1; sell < len(prices); sell++ {
			profit := prices[sell] - prices[buy]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
