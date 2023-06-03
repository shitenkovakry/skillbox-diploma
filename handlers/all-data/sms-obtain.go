package alldata

import (
	"sort"

	"skillbox-diploma/models"
)

func (handler *Handler) obtainSMSData() []models.SMSData {
	smsSortedByProvider := handler.sms.Read()
	sort.Slice(smsSortedByProvider, func(i, j int) bool {
		smsA := smsSortedByProvider[i]
		smsB := smsSortedByProvider[j]

		return smsA.Provider < smsB.Provider
	})

	smsSortedByCountry := handler.sms.Read()
	sort.Slice(smsSortedByCountry, func(i, j int) bool {
		smsA := smsSortedByCountry[i]
		smsB := smsSortedByCountry[j]

		return smsA.Country < smsB.Country
	})

	return []models.SMSData{smsSortedByProvider, smsSortedByCountry}
}
