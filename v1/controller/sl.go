package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	// "errors"
	// "os"
	// "path/filepath"

	"github.com/insektionen/IN-TV/api"
	// "github.com/spf13/viper"
)

type SLController struct{}

func NewSLController() api.Controller {
	return &SLController{}
}

// HandleGet implements api.Controller.
func (c *SLController) HandleGet(r *http.Request) (api.ViewModel, error) {
	err, json_response := sl_api_get()
	if err != nil {
		fmt.Println("Error in requesting data from sl api: ", err)
		return nil, err
	}

	err, sl_data_simple := simplify_sl_data(json_response)
	if err != nil {
		fmt.Println("Error in simplifying json response: ", err)
		return nil, err
	}

	err = fix_sl_name_field(&sl_data_simple)
	if err != nil {
		fmt.Println("Error with fixing name field: ", err)
		return nil, err
	}

	err = fix_sl_direction_field(&sl_data_simple)
	if err != nil {
		fmt.Println("Error with fixing direction field: ", err)
		return nil, err
	}

	err = calculate_time_till_departure(&sl_data_simple)
	if err != nil {
		fmt.Println("Error with calculating time till departure: ", err)
		return nil, err
	}
	
	// Marshal the struct into JSON.
    jsonBytes, err := json.Marshal(sl_data_simple)
    if err != nil {
        fmt.Println("Error:", err)
        return nil, err
    }
	
	// return api.JSONView(sl_data_simple), nil
	return api.JSONView(jsonBytes), nil
	// fmt.Println(sl_data_simple)
}

// HandleDelete implements api.Controller.
func (*SLController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	panic("unimplemented")
}

// HandlePost implements api.Controller.
func (*SLController) HandlePost(r *http.Request) (api.ViewModel, error) {
	panic("unimplemented")
}

// HandlePut implements api.Controller.
func (*SLController) HandlePut(r *http.Request) (api.ViewModel, error) {
	panic("unimplemented")
}







// The struct for the simplified json data from SL API
type sl_data_t struct {
	Name                string        `json:"name"`
	Time                string        `json:"time"`
	Date                string        `json:"date"`
	Direction           string        `json:"direction"`
	RTTime              string        `json:"rtTime"`
	Time_till_departure time.Duration `json:"time_till_departure"`
}

func sl_api_get() (error, []byte) {
	// Define the URL of the API you want to request
	apiUrl := "https://api.resrobot.se/v2.1/departureBoard?"
	apiUrl += "&format=" + "json"
	apiUrl += "&id=" + "740012883"
	apiUrl += "&maxJourneys=" + "3"
	apiUrl += "&accessId=" + "ea7df8ff-ab5d-4a50-91a2-d5c4a520e56d"

	// Send a GET request to the API
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return err, nil
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status code:", response.StatusCode)
		return err, nil
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		return err, nil
	}

	return nil, body
}

func simplify_sl_data(data_in []byte) (error, []sl_data_t) {
	// Unmarshal the JSON data into a map
	var data map[string]interface{}
	err := json.Unmarshal(data_in, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return err, nil
	}

	var sl_data []sl_data_t // Create an array for the important sl info

	// Getting past the Departure nesting
	departure, ok := data["Departure"].([]interface{})
	if !ok {
		fmt.Println("Error: 'Departure' is not an array")
		return err, nil
	}

	// Loop through the departure array
	for i, entity := range departure {

		entityBytes, err := json.Marshal(entity)
		if err != nil {
			fmt.Println("Error:", err)
			return err, nil
		}

		// Unmarshaling into temproary sl_data instance
		var sl sl_data_t
		err = json.Unmarshal(entityBytes, &sl)
		if err != nil {
			fmt.Println("Error: 'entity' is not a valid JSON object")
			return err, nil
		}

		// Ensure sl_data has enough capacity
		if i >= len(sl_data) {
			sl_data = append(sl_data, sl)
		} else {
			sl_data[i] = sl
		}

		sl_data = append(sl_data, sl) // Fill in the array of sl_data

		// fmt.Printf("Departure[%d]: %+v\n", i, sl) // For debugging
	}
	return nil, sl_data
}

func fix_sl_name_field(sl_data_in *[]sl_data_t) error {
	for i := range *sl_data_in {
		current_name := (*sl_data_in)[i].Name              // variable to simplify the current name
		if strings.Contains(current_name, "-Tunnelbana") { // Check if text is Tunnelbana
			(*sl_data_in)[i].Name = strings.Replace(current_name, "Länstrafik -", "", 1) // Take out first word of the name
		}
		if strings.Contains(current_name, "Buss") { // Check if text is Buss
			(*sl_data_in)[i].Name = strings.Replace(current_name, "Länstrafik - ", "", 1) // Take out first word of the name
		}
	}
	return nil
}

func fix_sl_direction_field(sl_data_in *[]sl_data_t) error {
	for i := range *sl_data_in {
		(*sl_data_in)[i].Direction = fix_direction_case((*sl_data_in)[i].Direction, " (Stockholm kn)")
		(*sl_data_in)[i].Direction = fix_direction_case((*sl_data_in)[i].Direction, " T-bana")
		(*sl_data_in)[i].Direction = fix_direction_case((*sl_data_in)[i].Direction, " (Järfälla kn)")
		(*sl_data_in)[i].Direction = fix_direction_case((*sl_data_in)[i].Direction, " (Danderyd kn)")
	}
	return nil
}

func fix_direction_case(current_name, to_fix string) string {
	if strings.Contains(current_name, to_fix) { // Check if text is Buss
		return strings.Replace(current_name, to_fix, "", 1) // Take out first word of the name
	}
	return current_name
}

func calculate_time_till_departure(sl_data *[]sl_data_t) error {
	loc, err := time.LoadLocation("Europe/Stockholm") // Setting the location for timezone
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	current_time := time.Now().In(loc) // Get the current time in Stockholm

	layout := "2006-01-02 15:04:05" // Specify the correct time format
	var departure_time string

	// Parse the time string with the specified layout depending on if real time update is available
	for i := range *sl_data {
		if (*sl_data)[i].RTTime != "" {
			departure_time = (*sl_data)[i].Date + " " + (*sl_data)[i].RTTime
		} else {
			departure_time = (*sl_data)[i].Date + " " + (*sl_data)[i].Time
		}

		// parsed_departure_time, err := time.Parse(layout, departure_time)
		parsed_departure_time, err := time.ParseInLocation(layout, departure_time, loc)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		(*sl_data)[i].Time_till_departure = parsed_departure_time.Sub(current_time)
		// TODO: change format so only hours and minutes get stored.

		// For debugging
		// fmt.Print("Parsed Time:", parsed_departure_time)
		// fmt.Print(" Current Time:", current_time)
		// fmt.Println(" Time Difference:", (*sl_data)[i].Time_till_departure)
	}
	return nil
}
