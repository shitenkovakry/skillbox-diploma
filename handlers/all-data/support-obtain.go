package alldata

func (handler *Handler) obtainSupportData() []int {
	support := handler.support.Read()
	resultOfSupport := []int{}
	supportNotLoaded := 1
	supportIsModeratelyLoaded := 2
	supportIsOverLoaded := 3
	averageTimeInMinutesPerTicket := 3

	for index := 0; index < len(support); index++ {
		elementOfSupport := support[index]
		numberOfOpenTickets := elementOfSupport.ActiveTickets

		if elementOfSupport.ActiveTickets < 9 {
			resultOfSupport = append(resultOfSupport, supportNotLoaded)
		} else if elementOfSupport.ActiveTickets >= 9 && elementOfSupport.ActiveTickets < 16 {
			resultOfSupport = append(resultOfSupport, supportIsModeratelyLoaded)
		} else if elementOfSupport.ActiveTickets > 16 {
			resultOfSupport = append(resultOfSupport, supportIsOverLoaded)
		}

		responseTimeForANewTicket := numberOfOpenTickets * averageTimeInMinutesPerTicket

		resultOfSupport = append(resultOfSupport, responseTimeForANewTicket)

	}

	return resultOfSupport
}
