package alldata

import "skillbox-diploma/models"

const (
	supportIsNotOverloaded        = 9
	supportIsModeratelyOverloaded = 16
	supportNotLoaded              = 1
	supportIsModeratelyLoaded     = 2
	supportIsOverLoaded           = 3
	averageTimeInMinutesPerTicket = 3
)

func (handler *Handler) obtainSupportData() []int {
	support := handler.support.Read()
	rateOfLoadSupport := 0
	responseTimeForANewTicket := 0

	arrayWithTickets := handler.moveNumbersOfTicketToAnotherArray(support)
	averageCountOfTickets := handler.calculateAverageOfOpenTickets(arrayWithTickets)
	summaOfOpenTickets := handler.calculateSummaOfOpenTickets(arrayWithTickets)

	if averageCountOfTickets < supportIsNotOverloaded {
		rateOfLoadSupport = supportNotLoaded
	} else if averageCountOfTickets >= supportIsNotOverloaded &&
		averageCountOfTickets < supportIsModeratelyOverloaded {
		rateOfLoadSupport = supportIsModeratelyLoaded
	} else if averageCountOfTickets > supportIsModeratelyOverloaded {
		rateOfLoadSupport = supportIsOverLoaded
	}

	responseTimeForANewTicket = summaOfOpenTickets * averageTimeInMinutesPerTicket

	return []int{rateOfLoadSupport, responseTimeForANewTicket}
}

func (handler *Handler) moveNumbersOfTicketToAnotherArray(support models.SupportData) []int {
	lenOfSupport := len(support)
	arrayWithMovedTickets := []int{}

	for index := 0; index < lenOfSupport; index++ {
		elementOfSupport := support[index]
		numberOfOpenTickets := elementOfSupport.ActiveTickets

		arrayWithMovedTickets = append(arrayWithMovedTickets, numberOfOpenTickets)
	}

	return arrayWithMovedTickets
}

func (handler *Handler) calculateAverageOfOpenTickets(arrayWithTickets []int) int {
	lenOfArrayWithTickets := len(arrayWithTickets)
	summaOfTickets := handler.calculateSummaOfOpenTickets(arrayWithTickets)

	averageOfOpenTickets := summaOfTickets / lenOfArrayWithTickets

	return averageOfOpenTickets
}

func (handler *Handler) calculateSummaOfOpenTickets(arrayWithTickets []int) int {
	lenOfArrayWithTickets := len(arrayWithTickets)
	summaOfOpenTickets := 0

	for index := 0; index < lenOfArrayWithTickets; index++ {
		elementOfArray := arrayWithTickets[index]

		summaOfOpenTickets += elementOfArray
	}

	return summaOfOpenTickets
}
