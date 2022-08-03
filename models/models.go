package models

import "time"

type Trains []Train

type Train struct {
	TrainID            int       `json:"trainId"`
	DepartureStationID int       `json:"departureStationId"`
	ArrivalStationID   int       `json:"arrivalStationId"`
	Price              float32   `json:"price"`
	ArrivalTime        time.Time `json:"arrivalTime"`
	DepartureTime      time.Time `json:"departureTime"`
}

type Tmp struct {
	TrainID            int     `json:"trainId"`
	DepartureStationID int     `json:"departureStationId"`
	ArrivalStationID   int     `json:"arrivalStationId"`
	Price              float32 `json:"price"`
	ArrivalTime        string  `json:"arrivalTime"`
	DepartureTime      string  `json:"departureTime"`
}
