package vbb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Vbb struct {
	Key string
}

type JourneyDetailRef struct {
	Ref string `json:"ref"`
}

type Product struct {
	Name         string `json:"name"`
	Num          string `json:"num"`
	Line         string `json:"line"`
	CatOut       string `json:"catOut"`
	CatIn        string `json:"catIn"`
	CatCode      string `json:"catCode"`
	CatOutS      string `json:"catOutS"`
	CatOutL      string `json:"catOutL"`
	OperatorCode string `json:"operatorCode"`
	Operator     string `json:"operator"`
	Admin        string `json:"admin"`
}

type Departure struct {
	JourneyDetailRef JourneyDetailRef `json:"JourneyDetailRef"`
	Product          Product          `json:"Product"`
	Name             string           `json:"name"`
	Type             string           `json:"type"`
	Stop             string           `json:"stop"`
	StopId           string           `json:"stopid"`
	StopExtId        string           `json:"stopExtId"`
	PrognosisType    string           `json:"prognosisType"`
	Time             string           `json:"time"`
	Date             string           `json:"date"`
	RtTime           string           `json:"rtTime"`
	RtDate           string           `json:"rtDate"`
	Direction        string           `json:"direction"`
	TrainNumber      string           `json:"trainNumber"`
	TrainCategory    string           `json:"trainCategory"`
}

type DepartureBoard struct {
	Departures []Departure `json:"Departure"`
}

func (v *Vbb) GetDepartureBoard(depRequest *Departure) (*DepartureBoard, error) {
	parameters := url.Values{}

	parameters.Add("accessId", v.Key)
	parameters.Add("id", depRequest.StopId)
	parameters.Add("direction", depRequest.Direction)
	parameters.Add("format", "json")

	Url := url.URL{
		Scheme:   "http",
		Host:     "demo.hafas.de",
		Path:     "/openapi/vbb-proxy/departureBoard",
		RawQuery: parameters.Encode(),
	}

	res, err := http.Get(Url.String())

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var depBoard DepartureBoard

	err = json.Unmarshal(body, &depBoard)

	if err != nil {
		return nil, err
	}

	return &depBoard, nil
}
