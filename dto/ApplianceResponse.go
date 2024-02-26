package dto

type ApplianceResponse struct {
	SerialNo                 string   `json:"serialNo"`
	TheatreName              string   `json:"theatreName"`
	Location                 Location `json:"location"`
	IspPaymentResponsibility string   `json:"ispPaymentResponsibility"`
	Bandwidth                string   `json:"bandwidth"`
	AvgBandwidth             string   `json:"avgBandwidth"`
	PlanStartDate            string   `json:"planStartDate"`
	BillingCycle             string   `json:"billingCycle"`
	DeviceStatus             string   `json:"deviceStatus"`
	DownloadStatus           string   `json:"downloadStatus"`
	OsVersion                string   `json:"osVersion"`
	Storage                  string   `json:"storage"`
}
