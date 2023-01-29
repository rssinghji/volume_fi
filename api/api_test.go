package api

import (
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
