package alldata

import (
	"skillbox-diploma/models"
	"sort"
)

func (handler *Handler) obtainMMSData() []models.MMSData {
	mmsSortedByProvider := handler.mms.Read()
	sort.Slice(mmsSortedByProvider, func(i, j int) bool {
		mmsA := mmsSortedByProvider[i]
		mmsB := mmsSortedByProvider[j]

		return mmsA.Provider < mmsB.Provider
	})

	mmsSortedByCountry := handler.mms.Read()
	sort.Slice(mmsSortedByCountry, func(i, j int) bool {
		mmsA := mmsSortedByCountry[i]
		mmsB := mmsSortedByCountry[j]

		return mmsA.Country < mmsB.Country
	})

	return []models.MMSData{mmsSortedByProvider, mmsSortedByCountry}
}
