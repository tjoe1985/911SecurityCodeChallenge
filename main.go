package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

type Sensor struct {
	ID         string  `json:"id"`
	GUID       string  `json:"guid"`
	IsActive   bool    `json:"isActive"`
	Registered string  `json:"registered"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Range      int     `json:"range"`
}
type Storm struct {
	Name      string  `json:"name"`
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
}

// RequestData contains One slice with sensor and another one with Storms
type RequestData struct {
	Sensor []Sensor `json:"sensor"`
	Storms []Storm  `json:"storms"`
}

func main() {
	http.HandleFunc("/findoverlabs", findOverlabsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
findOverlabsHandler takes a request with a json containing an array
of sensors and storms. Unmarshalling the content in to requestData
and passing it to checkSensorStormOverlap to get a string array with
all overlap events and returning and logging them for the user.
*/
func findOverlabsHandler(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData
	// Parse the request body into the requestData struct
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// pass incoming data to checkSensorStomrOverlap
	overlaps := checkSensorStormOverlap(requestData.Sensor, requestData.Storms)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(overlaps)
}

/*
	checkSensorStormOverlap takes a slice of Sensors and Storms and checks

for overlaps. It will ignore inactive sensors but could be improved by adding
the flag checkActiveOnly that way engineers sending information have the choice
to use it as needed.
*/
func checkSensorStormOverlap(sensors []Sensor, storms []Storm) []string {

	var overlappingSensors []string

	for _, sensor := range sensors {

		if !sensor.IsActive {
			fmt.Println("Sensor not active moving on")
			continue
		}

		for _, storm := range storms {

			if IsWithinRadius(sensor.Latitude, sensor.Longitude, storm.Latitude, storm.Longitude, storm.Radius) {
				log.Println("Sensor ID: " + sensor.ID + " Overlapping with Storm:" + storm.Name + " Duration will be added soon")
				overlappingSensors = append(overlappingSensors, "Sensor ID: "+sensor.ID+" Overlapping Storm:"+storm.Name+" Duration will be added soon")
			}
		}
	}

	return overlappingSensors
}

// IsWithinRadius checks if a coordinate is within the specified radius of another coordinate.
func IsWithinRadius(SensorLatitude, SensorLongitude, StormLatitude, StormLongitude, radius float64) bool {
	distance := CalculateDistance(SensorLatitude, SensorLongitude, StormLatitude, StormLongitude)
	return distance <= radius
}

// CalculateDistance calculates the distance between two coordinates using the Haversine formula.
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert coordinates from degrees to radians
	lat1Rad := degreesToRadians(lat1)
	lon1Rad := degreesToRadians(lon1)
	lat2Rad := degreesToRadians(lat2)
	lon2Rad := degreesToRadians(lon2)

	// Haversine formula
	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	earthRadius := 6371.0 // Radius of the Earth in kilometers
	distance := earthRadius * c

	return distance
}

// degreesToRadians converts degrees to radians.
func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
