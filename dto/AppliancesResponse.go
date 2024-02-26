package dto

type AppliancesResponse struct {
	Appliances []Appliances `json:"appliances"`
}

type Appliances struct {
	SerialNo       string   `json:"serialNo"`
	TheatreName    string   `json:"theatreName"`
	Location       Location `json:"location"`
	Bandwidth      string   `json:"bandwidth"`
	AvgBandwidth   string   `json:"avgBandwidth"`
	DeviceStatus   string   `json:"deviceStatus"`
	DownloadStatus string   `json:"downloadStatus"`
	OsVersion      string   `json:"osVersion"`
}

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}
