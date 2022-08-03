package tools

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"technicalTask/models"
	"time"
)

const (
	file       = "data/data.json"
	timeLayout = "15:04:05"
)

func FindTrains(depStation, arrStation, criteria string) (models.Trains, error) {
	err := validInputChecker(depStation, arrStation, criteria)
	if err != nil {
		return nil, err
	}

	trains, err := getTrains(file, depStation, arrStation)
	if err != nil {
		return nil, err
	}

	sortedTrains, err := sort(trains, criteria)
	if err != nil {
		return nil, err
	}

	return sortedTrains, nil
}

func validInputChecker(dep, arr, criteria string) error {
	if dep == "" {
		return errors.New("empty departure station")
	}
	if arr == "" {
		return errors.New("empty arrival station")
	}
	num, err := strconv.Atoi(dep)
	if (err != nil) || (num <= 0) {
		return errors.New("bad departure station input")
	}
	num, err = strconv.Atoi(arr)
	if (err != nil) || (num <= 0) {
		return errors.New("bad arrival station input")
	}

	switch criteria {
	case "price", "arrival-time", "departure-time":
		return nil
	default:
		return errors.New("unsupported criteria")
	}
}

func getTrains(s, dep, ar string) (models.Trains, error) {
	var tmpTrains []models.Tmp
	var trains models.Trains

	sliceByte, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(sliceByte, &tmpTrains)
	if err != nil {
		return nil, err
	}

	d, err := strconv.Atoi(dep)
	if err != nil {
		return nil, err
	}

	a, err := strconv.Atoi(ar)
	if err != nil {
		return nil, err
	}

	for i, _ := range tmpTrains {
		var train models.Train

		if tmpTrains[i].DepartureStationID != d {
			continue
		}

		if tmpTrains[i].ArrivalStationID != a {
			continue
		}
		train.TrainID = tmpTrains[i].TrainID
		train.DepartureStationID = tmpTrains[i].DepartureStationID
		train.ArrivalStationID = tmpTrains[i].ArrivalStationID
		train.Price = tmpTrains[i].Price

		tArrival, err := timeParser(tmpTrains[i].ArrivalTime)
		if err != nil {
			return nil, err
		}
		train.ArrivalTime = tArrival

		tDeparture, err := timeParser(tmpTrains[i].DepartureTime)
		if err != nil {
			return nil, err
		}
		train.DepartureTime = tDeparture

		trains = append(trains, train)
	}

	return trains, nil
}

func timeParser(s string) (time.Time, error) {
	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func sort(trains models.Trains, criteria string) (models.Trains, error) {
	switch criteria {
	case "price":
		for i := 1; i < len(trains); i++ {
			for j := i; j > 0; j-- {
				if trains[j-1].Price < trains[j].Price {
					break
				}
				trains[j-1], trains[j] = trains[j], trains[j-1]
			}
		}
		return trains, nil
	case "arrival-time":
		for i := 1; i < len(trains); i++ {
			for j := i; j > 0; j-- {
				if trains[j-1].ArrivalTime.Before(trains[j].ArrivalTime) {
					break
				}
				trains[j-1], trains[j] = trains[j], trains[j-1]
			}
		}
		return trains, nil
	case "departure-time":
		for i := 1; i < len(trains); i++ {
			for j := i; j > 0; j-- {
				if trains[j-1].DepartureTime.Before(trains[j].DepartureTime) {
					break
				}
				trains[j-1], trains[j] = trains[j], trains[j-1]
			}
		}
		return trains, nil
	}

	return nil, nil
}
