package handler

import (
	"encoding/json"
	"net/http"
	Service "qubeChallenge/Service"
	"qubeChallenge/errs"
	"regexp"

	"github.com/gorilla/mux"
)

type Handlers struct {
	Service Service.Service
}

func (s *Handlers) Appliances(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	deviceStatus := queryParams.Get("deviceStatus")
	downloadStatus := queryParams.Get("downloadStatus")
	var Errors []errs.Errors

	if deviceStatus != "" && !regexp.MustCompile(`(?i)^(Online|Offline)$`).MatchString(deviceStatus) {
		Errors = append(Errors, errs.Errors{
			Code:    "AA001H",
			Message: "Invalid parameter, deviceStatus should be either 'Online' or 'Offline'",
		})
	}

	if downloadStatus != "" &&
		!regexp.MustCompile(`(?i)^(Failed|Stalled|Archived|Cancelled|Scheduled|Unarchiving|Downloading|Downloaded)$`).MatchString(downloadStatus) {
		Errors = append(Errors, errs.Errors{
			Code: "AA002H",
			Message: `Invalid parameter, downloadStatus should be one of the following: 'Failed', 'Stalled', 
			'Archived', 'Cancelled', 'Scheduled', 'Unarchiving', 'Downloading', or 'Downloaded'`,
		})
	}

	if len(Errors) > 0 {
		writeResponse(w, http.StatusBadRequest, errs.AppError{
			HTTPStatus: http.StatusBadRequest,
			HTTPCode:   "BadRequest",
			Errors:     Errors,
		})
		return
	}

	AppliancesResponse, err := s.Service.Appliances(deviceStatus, downloadStatus)
	if err != nil {
		writeResponse(w, err.HTTPStatus, err)
		return
	}

	writeResponse(w, http.StatusOK, AppliancesResponse)
}

func (s *Handlers) Appliance(w http.ResponseWriter, r *http.Request) {

	applianceId := mux.Vars(r)["appliance-id"]

	if applianceId == "" {
		writeResponse(w, http.StatusBadRequest, errs.AppError{
			HTTPStatus: http.StatusBadRequest,
			HTTPCode:   "BadRequest",
			Errors: []errs.Errors{{
				Code:    "A001H",
				Message: "Missing parameter, appliance-id",
			}},
		})
		return
	}

	AppliancesResponse, err := s.Service.Appliance(applianceId)
	if err != nil {
		writeResponse(w, err.HTTPStatus, err)
		return
	}

	writeResponse(w, http.StatusOK, AppliancesResponse)
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Check if the status code allows a response body
	if statusCode != http.StatusNoContent {
		// Marshal data to JSON
		jsonResponse, err := json.Marshal(data)
		if err != nil {
			// Handle JSON marshaling error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Write the response body
		_, err = w.Write(jsonResponse)
		if err != nil {
			// Handle write error
			return
		}
	}
}
