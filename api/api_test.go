package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	util "volumefi.com/volume_fi/utilities"
)

func TestCalculateFlightPath(test *testing.T) {
	req := util.Request{}
	req.FlightList = append(req.FlightList, []string{"A", "B"})
	req.FlightList = append(req.FlightList, []string{"B", "C"})
	resp, err := calculateFlightPath(req)
	if err != nil {
		test.Fail()
	}
	if resp.FlightPath[0] != "A" && resp.FlightPath[1] != "B" && resp.FlightPath[2] != "C" {
		test.Fail()
	}
}

func TestCalculateFlightPath2(test *testing.T) {
	req := util.Request{}
	req.FlightList = append(req.FlightList, []string{"A", "B"})
	resp, err := calculateFlightPath(req)
	if err != nil {
		test.Fail()
	}
	if resp.FlightPath[0] != "A" && resp.FlightPath[1] != "B" {
		test.Fail()
	}
}

func TestCalculateFlightPathBadInput(test *testing.T) {
	req := util.Request{}
	_, err := calculateFlightPath(req)
	if err == nil {
		test.Fail()
	}
}

func testEndpoint(test *testing.T, method string, endpoint string, request []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, endpoint, bytes.NewBuffer(request))
	writer := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleCalculateFlightTracker)
	handler.ServeHTTP(writer, req)

	log.Printf("writer resp code [%d], writer resp [%s]", writer.Code, writer.Body.String())
	return writer
}

func TestCalculateEndpoint(test *testing.T) {
	writer := testEndpoint(test, "POST", "/calculate", []byte(`{
		"flight_list": [["A","B"],["B","C"]]}`))

	if http.StatusOK != writer.Code {
		test.Error("expected", http.StatusOK, "got", writer.Code)
	}
	body := writer.Body.String()
	if len(body) == 0 {
		test.Fail()
	} else {
		resp := util.Response{}
		json.Unmarshal(writer.Body.Bytes(), &resp)
		if len(resp.FlightPath) != 3 {
			test.Fail()
		}
	}
}

func TestCalculateEndpointWrongInput(test *testing.T) {
	writer := testEndpoint(test, "POST", "/calculate", []byte(`{
		"flight_list": []}`))
	if http.StatusBadRequest != writer.Code {
		test.Error("expected", http.StatusBadRequest, "got", writer.Code)
	}

	writer = testEndpoint(test, "POST", "/calculate", []byte(`{
		"flight_list": [[]]}`))
	if http.StatusInternalServerError != writer.Code {
		test.Error("expected", http.StatusInternalServerError, "got", writer.Code)
	}
}
