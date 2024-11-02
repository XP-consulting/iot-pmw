package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

// User represents the structure of the JSON data
type json_package struct {
	BaseName    string  `json:"bn"`
	BaseTime    float64 `json:"bt"`
	Name        string  `json:"n"`
	BaseVersion int16   `json:"bver"`
	Exception   string  `json:"x"`
	Error       string  `json:"e"`
	SignalEvent string  `json:"s"`
	Severity    int16   `json:"se"`
	Description string  `json:"d"`
	Payload     string  `json:"p"`
}

// parseJSON parses JSON data into a User struct
func parseJSON(data []byte) ([]json_package, error) {
	var user []json_package
	// var user json_package
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return []json_package{}, err
	}
	//return user, nil
	return user, nil
}

// Unused - function to parse unix epoch to human readable utc time
func parseEpoch(epoch float64) time.Time {
	tm := time.Unix(int64(epoch), 0).UTC()
	return tm
}

func main() {
	// Sample JSON data
	//jsonData := `
	//[{"bn": "urn:dev:name:twg0043:", "bt": 176627612.2, "n": "/dev/sda0", "x": "Storage full", "se": 255, "d": "Disk 100.0% full"},
	//{"n": "wlan0", "e": "Disconnected", "se": 100, "d": "Connection lost"},
	//{"n": "ufw", "x": "Disabled", "se": 100, "d": "Firewall disabled"}]
	//`
	//Sample JSON data
	jsonData2 := `[{"bn": "urn:dev:mac:00170d451f62:", "bt": 176627612.2, "n": "sensor/gps", "x": "Sensor malfuction", "se": 200, "d": "GPS sensor not connected"},
  {"n": "co2/0", "e": "Bad sensor reading", "se": 200, "d": "Negative value -4 ppm"},
  {"n": "temperature/0", "e": "Bad sensor reading", "se": 200, "d": "Outside sensing range(-40 to +85): 6387.4"},
  {"n": "battery", "e": "Low battery level", "se": 100, "d": "Battery at 12/%/ capacity"}
]`

	// Parse the JSON data
	user, err := parseJSON([]byte(jsonData2))
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	for _, element := range user {
		values := reflect.ValueOf(element)
		types := values.Type()
		for i := 0; i < values.NumField(); i++ {
			fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
		}
	}

}
