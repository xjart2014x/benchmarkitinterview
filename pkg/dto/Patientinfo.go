package dto

type PatientInfo struct {
	Name            string `json:"patientName"`
	Address         string `json:"address"`
	ActiveInsurance bool   `json:"activeInsurance"`
	Age             int    `json:"age"`
	TIN             int    `json:"tin"`
	PowerofAttorney string `json:"powerOfAttorney"`
}
