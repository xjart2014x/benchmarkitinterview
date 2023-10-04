package dto

type MedicineTrial struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	RiskLevel         string        `json:"riskLevel"`
	NumberOfSuccesses int           `json:"numberOfSuccesses"`
	NumberOfFailures  int           `json:"numberOfFailures"`
	SimilarMedicines  []string      `json:"similarMedicines"`
	TrialSites        *[]TrialSites `json:"trialSites"`
}

type TrialSites struct {
	Name           string `json:"name"`
	State          string `json:"state"`
	BuildingNumber int    `json:"buildingNumber"`
}
