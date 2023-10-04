package dto

type PatientInfo struct {
	Name             string `json:"patientName"`
	Address          string `json:"address"`
	ActiveInsurance  bool   `json:"activeInsurance"`
	Age              int    `json:"age"`
	TIN              int    `json:"tin"`
	PowerofAttorney  string `json:"powerOfAttorney"`
	MedicineID       int    `json:"medicineId"`
	MedicalTrialName string `json:"medicalTrialName"`
	MedicalRiskLevel string `json:"medicalRiskLevel"`
}
