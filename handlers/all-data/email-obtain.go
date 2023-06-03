package alldata

import (
	"skillbox-diploma/models"
	"sort"
)

func (handler *Handler) obtainEmailData() []models.EmailData {
	emailSortedByProvider := handler.email.Read()
	sort.Slice(emailSortedByProvider, func(i, j int) bool {
		emailA := emailSortedByProvider[i]
		emailB := emailSortedByProvider[j]

		return emailA.Provider < emailB.Provider
	})

	emailSortedByCountry := handler.email.Read()
	sort.Slice(emailSortedByCountry, func(i, j int) bool {
		emailA := emailSortedByCountry[i]
		emailB := emailSortedByCountry[j]

		return emailA.Country < emailB.Country
	})

	return []models.EmailData{emailSortedByProvider, emailSortedByCountry}
}
