package main

func createStatementData(invoice []*Invoice) StatementData {
	data.Customer = invoice[0].Customer
	// as javascript map()
	invoice[0].mapToEnrichPerformance()
	data.Performances = invoice[0].Performances
	data.TotalAmount = totalAmount(data.Performances)
	data.TotalCredits = totalVolumeCredits(data.Performances)

	return data
}

func (i *Invoice) mapToEnrichPerformance() {
	for k, v := range i.Performances {
		// v is local temp variable, not actual Performance
		i.Performances[k].Play = playFor(v)
		i.Performances[k].Play.Amount = amountFor(i.Performances[k])
		i.Performances[k].Play.Credits = volumeCreditsFor(i.Performances[k])
	}
	return
}

func playFor(aPerformance Performance) Play {
	return plays[aPerformance.PlayID]
}
