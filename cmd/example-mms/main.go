package main

import (
	"fmt"
	"skillbox-diploma/datasources/mms"
)

func main() {
	mms := mms.New("http://localhost:8383/mms")

	mmsData := mms.Load()

	fmt.Println(mmsData)
}
