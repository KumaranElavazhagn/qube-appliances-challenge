package service

import (
	"net/http"
	"qubeChallenge/dto"
	"qubeChallenge/errs"
	mockResponse "qubeChallenge/mock"
	"strings"
)

type service struct{}

type Service interface {
	Appliances(deviceStatus string, downloadStatus string) (*dto.AppliancesResponse, *errs.AppError)
	Appliance(applianceId string) (*dto.ApplianceResponse, *errs.AppError)
}

func (r service) Appliances(deviceStatus string, downloadStatus string) (*dto.AppliancesResponse, *errs.AppError) {
	// Initialize filteredAppliances with the full list of appliances
	filteredAppliances := mockResponse.Response.Appliances

	// Apply filter if either deviceStatus or downloadStatus is provided
	if deviceStatus != "" || downloadStatus != "" {
		var filtered []dto.Appliances
		for _, appliance := range filteredAppliances {
			if (deviceStatus == "" || strings.EqualFold(appliance.DeviceStatus, deviceStatus)) &&
				(downloadStatus == "" || strings.EqualFold(appliance.DownloadStatus, downloadStatus)) {
				filtered = append(filtered, appliance)
			}
		}
		filteredAppliances = filtered
	}

	// Check if any appliances are found after filtering
	if len(filteredAppliances) == 0 {
		return nil, errs.GenerateErrorResponse(http.StatusNotFound, "Not Found", []errs.Errors{{
			Code:    "AA001S",
			Message: "Appliances not found",
		}})
	}

	// Return the filtered appliances
	return &dto.AppliancesResponse{
		Appliances: filteredAppliances,
	}, nil
}

func (r service) Appliance(applianceId string) (*dto.ApplianceResponse, *errs.AppError) {
	for _, appliance := range mockResponse.Response.Appliances {
		if strings.EqualFold(applianceId, appliance.SerialNo) {
			response := dto.ApplianceResponse{
				SerialNo:                 appliance.SerialNo,
				TheatreName:              appliance.TheatreName,
				Location:                 appliance.Location,
				IspPaymentResponsibility: "Qube",
				Bandwidth:                appliance.Bandwidth,
				AvgBandwidth:             appliance.AvgBandwidth,
				PlanStartDate:            "1st Oct",
				BillingCycle:             "monthly",
				DeviceStatus:             appliance.DeviceStatus,
				DownloadStatus:           appliance.DownloadStatus,
				OsVersion:                appliance.OsVersion,
				Storage:                  "828 GB",
			}
			return &response, nil
		}
	}
	// Appliance not found
	return nil, errs.GenerateErrorResponse(http.StatusNotFound, "Not Found", []errs.Errors{{
		Code:    "A001S",
		Message: "Appliance not found",
	}})

}

func NewService() Service {
	return service{}
}
