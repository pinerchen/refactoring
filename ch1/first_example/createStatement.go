package main

func createStatementData(invoice []*Invoice) StatementData {
	Data.Customer = invoice[0].Customer
	// as javascript map()
	invoice[0].mapToEnrichPerformance()
	Data.Performances = invoice[0].Performances
	Data.TotalAmount = totalAmount(Data.Performances)
	Data.TotalCredits = totalVolumeCredits(Data.Performances)

	return Data
}

func (i *Invoice) mapToEnrichPerformance() {
	for k, v := range i.Performances {
		// v is local temp variable, not actual Performance
		var calculator = NewPerformanceCalculator(v, playFor(v))
		i.Performances[k].Play = calculator.play
		i.Performances[k].Play.Amount = calculator.amount()
		i.Performances[k].Play.Credits = volumeCreditsFor(i.Performances[k])
	}
	return
}

func playFor(aPerformance Performance) Play {
	return Plays[aPerformance.PlayID]
}
