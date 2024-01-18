package main

import (
	"fmt"
	"log"

	"github.com/Sk4ine/humanStruct/internal/humanStruct"
)

func main() {

	humanSlice := make([]humanStruct.IHuman, 0)

	humanSlice = append(humanSlice, createNewHuman(25, "Grigory"))
	humanSlice = append(humanSlice, createNewHuman(14, "Arkasha"))
	humanSlice = append(humanSlice, createNewHuman(105, "Zina"))
	humanSlice = append(humanSlice, createNewHuman(41, "Roman"))

	avgAge := humanStruct.CalculateAverageAge(humanSlice)
	fmt.Println(avgAge)

	fmt.Println(humanStruct.TryAdd(&humanSlice, createNewHuman(33, "Olga")))
}

func createNewHuman(age int, name string) humanStruct.IHuman {
	newHuman, err := humanStruct.NewHumanStruct(age, name)
	if err != nil {
		log.Fatal(err)
	}
	return newHuman
}
