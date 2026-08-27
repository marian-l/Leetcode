package arrays

/*
Idee: Aktuellen und nächsten Tag vergleichen, wenn verkauft werden soll. Es soll jeder Profit mitgenommen werden. Wenn Folgetag günstiger: Heute verkaufen (wenn möglich) und am Folgetag kaufen.
Wenn Folgetag teurer: heute kaufen (wenn möglich) und am Folgetag verkaufen. 
Immer nur tagesweise handeln: unbegrenzte Trades

Start: Wenn heute günstiger als morgen: kaufen und morgen verkaufen. auf maxProfit aufrechnen
Bool behalten, um festzustellen, ob Stock gehalten wird
*/

func MaxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	cheapestBuyPrice := prices[0]

	return maxProfit
}
