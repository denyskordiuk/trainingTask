package main

import (
	"fmt"
	"log"
	"technicalTask/tools"
)

func main() {
	var departureStation, arrivalStation, criteria string
	fmt.Println("Please enter departure station:\t")
	fmt.Scanf("%s\n", &departureStation)
	fmt.Println("Please enter arrival station:\t")
	fmt.Scanf("%s\n", &arrivalStation)
	fmt.Println("Please enter criteria:\t")
	fmt.Scanf("%s\n", &criteria)

	result, err := tools.FindTrains(departureStation, arrivalStation, criteria)
	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		for i := 0; i < 3; i++ {
			fmt.Printf("\nTrainID: %d, DepartureStationID: %d, ArrivalStationID: %d, Price: %f, ArrivalTime: %v, DepartureTime: %v \n\n", result[i].TrainID,
				result[i].DepartureStationID, result[i].ArrivalStationID, result[i].Price, result[i].ArrivalTime, result[i].DepartureTime)
		}
	}

}
