package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	util "volumefi.com/volume_fi/utilities"
)

// HandleDefault : Default API handler
/*
	Author		: Ravneet Singh
	Function 	: HandleDefault - Serves the default message for the default handler
	Input(s) 	: http.ResponseWriter, *http.Request (http.ResponseWriter to serve the response and *http.Request to process the request)
	Output(s) 	: None
	Example(s) 	: HandleDefault(response http.ResponseWriter, request *http.Request)
*/
func HandleDefault(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response, `{"message":"The server is up. Use /calculate endpoint for flight tracking"}`)
}

// HandleCalculateFlightTracker : API handler to handle /calculate endpoint
/*
	Author		: Ravneet Singh
	Function 	: HandleCalculateFlightTracker - Calculates the complete flight path based on JSON input reequest
	Input(s) 	: http.ResponseWriter, *http.Request (http.ResponseWriter to serve the response and *http.Request to process the request)
	Output(s) 	: None
	Example(s) 	: HandleCalculateFlightTracker(response http.ResponseWriter, request *http.Request)
*/
func HandleCalculateFlightTracker(response http.ResponseWriter, request *http.Request) {
	log.Println("HandleCalculateFlightTracker() begins")
	response.Header().Set("Content-Type", "application/json")
	if request.Method != "POST" {
		http.Error(response, `{"error":"Wrong HTTP method specified"}`, http.StatusMethodNotAllowed)
		fmt.Fprintf(response, util.ErrorMessageWrongMethod)
		return
	}
	flightList := util.Request{}
	err := json.NewDecoder(request.Body).Decode(&flightList)
	if err != nil || len(flightList.FlightList) == 0 {
		fmt.Fprintf(response, util.ErrorMessageBadrequest)
		return
	}

	prettyRequest, _ := json.MarshalIndent(flightList, "", "    ")
	log.Printf("Read request is: \n%s\n", string(prettyRequest))

	result, err := calculateFlightPath(flightList)
	if err != nil || (len(result.FlightPath) == 0) {
		log.Println("Something went wrong. ", err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, util.ErrorMessageServerError)
	} else {
		jsonResponse, err := json.Marshal(result)
		if err != nil {
			log.Println("Something went wrong. ", err.Error())
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(response, util.ErrorMessageServerError)
		} else {
			log.Println("Request processed successfully!")
			response.WriteHeader(http.StatusOK)
			fmt.Fprintf(response, "%s", string(jsonResponse[:]))
		}
	}

	log.Println("HandleCalculateFlightTracker() ends")
}

// calculateFlightPath : Actual algorithm to track flight path. I generally keep it in separate file. ket here for sake of easiness to track.
/*
	Author		: Ravneet Singh
	Function 	: calculateFlightPath - Calculates the actual flight path and returns a resonse and error if occured.
	Input(s) 	: http.ResponseWriter, *http.Request (http.ResponseWriter to serve the response and *http.Request to process the request)
	Output(s) 	: util.Response, error ([]string, error)
	Example(s) 	: calculateFlightPath(response http.ResponseWriter, request *http.Request)
*/
func calculateFlightPath(flightsList util.Request) (util.Response, error) {
	result := util.Response{}
	if len(flightsList.FlightList) == 0 {
		return result, fmt.Errorf("Incomplete or bad input")
	}

	// Get count of each key(airlort) to fetch start and end of the flight path assuming first entry as source and second entry as destination
	// Also store the source and destination as a key value pair to track the path

	overallString := ""
	flights := flightsList.FlightList
	flightMap := make(map[string]string)
	for index := 0; index < len(flightsList.FlightList); index++ {
		path := flights[index]
		flightMap[path[0]] = path[1]
		overallString += "," + path[0] + "," + path[1]
	}

	// Since each destination would be a source only once, and final destination would never be a source for another flight
	// we just need to find, which source occurred only once
	source := ""
	for key := range flightMap {
		keyCount := strings.Count(overallString, key)
		if keyCount == 1 {
			// We found the source. Store it and move on
			source = key
		}
	}

	// Now that we have the source, use it to track the flight path via using subsequent values for each key as next key
	nextKey := source
	for index := 0; index < len(flightMap); index++ {
		result.FlightPath = append(result.FlightPath, nextKey)
		nextKey = flightMap[nextKey]
	}
	result.FlightPath = append(result.FlightPath, nextKey) // Add the final destination

	return result, nil
}
