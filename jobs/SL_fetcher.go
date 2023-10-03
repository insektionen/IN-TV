package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/insektionen/IN-TV/v1/viewmodels"
)

// Define a custom type for the enum
type Interval int

// Define constants to represent enum values
const (
	High   Interval = 2
	Medium Interval = 5
	Low    Interval = 10
	None   Interval = 0
)

type timeSpan struct {
	startHour int
	stopHour  int
	interval  time.Duration
}

func SetTimerInterval(now time.Time, currentSpan timeSpan, t *time.Ticker) {
	if currentSpan.startHour < now.Hour() && currentSpan.stopHour > now.Hour() {
		if currentSpan.interval == 0 {
			t.Stop()
		} else {
			t.Reset(currentSpan.interval * time.Minute)
		}
	}
}

func FetchSLTimetable(exit <-chan bool) {

	timespan1 := timeSpan{3, 8, time.Duration(None)}
	timespan2 := timeSpan{8, 11, time.Duration(Low)}
	timespan3 := timeSpan{11, 18, time.Duration(High)}
	timespan4 := timeSpan{18, 3, time.Duration(Medium)}


	t := time.NewTicker(time.Duration(Medium) * time.Minute)

	// requestLimit := (11-8)*(60/10) + (18-11)*(60/2) + (24+3-18)*(60/10) = 282
	// 282 * 31 = 8742

	go func() {
		for {
			now := time.Now()
			SetTimerInterval(now, timespan1, t)
			SetTimerInterval(now, timespan2, t)
			SetTimerInterval(now, timespan3, t)
			SetTimerInterval(now, timespan4, t)

			time.Sleep(1 * time.Minute)
		}
	}()

	for {
		select {
		case <-t.C: // Run doFetchSL according to configuration
			doFetchSL()
		case <-exit: // Exit when program is terminated
			break
		}
	}
}

var SLDataSimple []*viewmodels.SLData

func doFetchSL() {
	// Define the URL of the API you want to request
	apiUrl := "https://api.resrobot.se/v2.1/departureBoard"

	query := url.Values{}
	query.Set("format", "json")
	query.Set("id", "740012883")
	query.Set("maxJourneys", "5")
	query.Set("accessId", "ea7df8ff-ab5d-4a50-91a2-d5c4a520e56d")

	urlPath := fmt.Sprintf("%s?%s", apiUrl, query.Encode())

	// Send a GET request to the API
	response, err := http.Get(urlPath)
	if err != nil {
		log.Println("Error making API request:", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		log.Println("API request failed with status code:", response.StatusCode)
		return
	}

	// Read the response body
	SLJSONData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading API response:", err)
		return
	}

	// Simplify the data from the SL JSON
	SLDataSimple, err := simplifySLData(SLJSONData)
	if err != nil {
		log.Println("Error simplifying SL data:", err)
		return
	}

	// Fixing the naming field to have only the buss or train
	err = FixSLNameField(SLDataSimple)
	if err != nil {
		log.Println("Error fixing name field:", err)
		return
	}

	// Fixing the direction field to have only the place name
	err = FixSLDirectionField(SLDataSimple)
	if err != nil {
		log.Println("Error fixing direction field:", err)
		return
	}

	// Calculating the TimeTillDeparture field
	err = CalcTimeTillDeparture(SLDataSimple)
	if err != nil {
		log.Println("Error calculating TimeTillDeparture:", err)
		return
	}

}

func simplifySLData(DataIn []byte) ([]*viewmodels.SLData, error) {
	// Unmarshal the JSON data into a map
	var data map[string]interface{}
	err := json.Unmarshal(DataIn, &data)
	if err != nil {
		return nil, fmt.Errorf("JSON error: %v", err)
	}

	var SLDataSimple []*viewmodels.SLData // Create an array for the important sl info

	// Getting past the Departure nesting
	departure, ok := data["Departure"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("JSON error: %v", err)
	}

	// Loop through the departure array
	for _, entity := range departure {
		entityBytes, err := json.Marshal(entity)
		if err != nil {
			return nil, fmt.Errorf("JSON error: %v", err)
		}

		// Unmarshaling into temproary SLData instance
		var SLInstance *viewmodels.SLData
		err = json.Unmarshal(entityBytes, &SLInstance)
		if err != nil {
			return nil, fmt.Errorf("JSON error: %v", err)
		}

		SLDataSimple = append(SLDataSimple, SLInstance) // Fill in the array of SLData
	}
	return SLDataSimple, nil
}

func FixSLNameField(SLDataIn []*viewmodels.SLData) error {
	for i := range SLDataIn {
		if strings.Contains(SLDataIn[i].Name, "-Tunnelbana") { // Check if text is Tunnelbana
			SLDataIn[i].Name = strings.Replace(SLDataIn[i].Name, "Länstrafik -", "", 1) // Take out first word of the name
		}
		if strings.Contains(SLDataIn[i].Name, "Buss") { // Check if text is Buss
			SLDataIn[i].Name = strings.Replace(SLDataIn[i].Name, "Länstrafik - ", "", 1) // Take out first word of the name
		}
	}
	return nil
}

func FixSLDirectionField(SLDataIn []*viewmodels.SLData) error {
	for i := range SLDataIn {
		FixDirectionCase(SLDataIn[i], " (Stockholm kn)")
		FixDirectionCase(SLDataIn[i], " T-bana")
		FixDirectionCase(SLDataIn[i], " (Järfälla kn)")
		FixDirectionCase(SLDataIn[i], " (Danderyd kn)")
	}
	return nil
}

func FixDirectionCase(SLInstance *viewmodels.SLData, TakeOut string) {
	if strings.Contains(SLInstance.Direction, TakeOut) { // Check if text is Buss
		SLInstance.Direction = strings.Replace(SLInstance.Direction, TakeOut, "", 1) // Take out first word of the name
	}
}

func CalcTimeTillDeparture(SLData []*viewmodels.SLData) error {
	loc, err := time.LoadLocation("Europe/Stockholm") // Setting the location for timezone
	if err != nil {
		return fmt.Errorf("Location error: %v", err)
	}
	CurrentTime := time.Now().In(loc) // Get the current time in Stockholm

	layout := "2006-01-02 15:04:05" // Specify the correct time format
	var DepartureTime string

	// Parse the time string with the specified layout depending on if real time update is available
	for i := range SLData {
		if SLData[i].RTTime != "" {
			DepartureTime = SLData[i].Date + " " + SLData[i].RTTime
		} else {
			DepartureTime = SLData[i].Date + " " + SLData[i].Time
		}

		ParsedDepartureTime, err := time.ParseInLocation(layout, DepartureTime, loc)
		if err != nil {
			return fmt.Errorf("Parsing time error: %v", err)
		}

		SLData[i].TimeTillDeparture = int(ParsedDepartureTime.Sub(CurrentTime).Minutes())
	}
	return nil
}
