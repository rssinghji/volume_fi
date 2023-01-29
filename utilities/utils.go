package utilities

// Request : struct to parse for API request
type Request struct {
	FlightList [][]string `json:"flight_list"`
}

// Response : struct to serve API response
type Response struct {
	FlightPath []string `json:"flight_path"`
}

// Constants for API handling
const (
	ErrorMessageWrongMethod = `{"error":"Wrong HTTP method."}`
	ErrorMessageBadrequest  = `{"error":"Bad Request."}`
	ErrorMessageServerError = `{"error":"Something went wrong."}`
)
