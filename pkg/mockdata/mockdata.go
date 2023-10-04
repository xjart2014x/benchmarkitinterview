// mocking up data for interview

package mockdata

import "github.com/xjart2014x/benchmarkitinterview/pkg/dto"

var Patients = []dto.PatientInfo{
	{
		Name:            "John Jones",
		Address:         "1234 Street Dr",
		ActiveInsurance: true,
		Age:             29,
		TIN:             938274839,
		PowerofAttorney: "Ron Jones",
		MedicineID:      1,
	},
	{
		Name:            "Lenny Norwood",
		Address:         "9283 Place Dr",
		ActiveInsurance: true,
		Age:             51,
		TIN:             384930293,
		PowerofAttorney: "Benny Norwood",
		MedicineID:      2,
	},
	{
		Name:            "Karen Michaels",
		Address:         "7283 Area Dr",
		ActiveInsurance: false,
		Age:             34,
		TIN:             284920485,
		PowerofAttorney: "Shawn Michaels",
		MedicineID:      3,
	},
}

var TrialInfo = []dto.MedicineTrial{
	{
		ID:                1,
		Name:              "Miracle Pill for COVID",
		RiskLevel:         "low",
		NumberOfSuccesses: 98,
		NumberOfFailures:  2,
		SimilarMedicines: []string{
			"miracle covid pill beta",
			"miracle covid pill pre release",
		},
		TrialSites: &[]dto.TrialSites{
			{
				Name:           "Los Angeles Area 51",
				State:          "California",
				BuildingNumber: 89067,
			},
			{
				Name:           "San Diego Area 52",
				State:          "California",
				BuildingNumber: 37425,
			},
		},
	},
	{
		ID:                2,
		Name:              "Miracle Pill for Cancer",
		RiskLevel:         "high",
		NumberOfSuccesses: 37,
		NumberOfFailures:  63,
		SimilarMedicines: []string{
			"miracle cancer pill beta",
			"miracle cancer pill pre release",
		},
		TrialSites: &[]dto.TrialSites{
			{
				Name:           "Miami Area 51",
				State:          "Florida",
				BuildingNumber: 28374,
			},
			{
				Name:           "Jacksonville Area 52",
				State:          "Florida",
				BuildingNumber: 91820,
			},
		},
	},
	{
		ID:                3,
		Name:              "Miracle Pill for Common Cold",
		RiskLevel:         "medium",
		NumberOfSuccesses: 98,
		NumberOfFailures:  2,
		SimilarMedicines: []string{
			"miracle common cold pill beta",
			"miracle common cold pill pre release",
		},
		TrialSites: &[]dto.TrialSites{
			{
				Name:           "Houston Area 51",
				State:          "Texas",
				BuildingNumber: 59203,
			},
			{
				Name:           "Dallas Area 52",
				State:          "Texas",
				BuildingNumber: 19384,
			},
		},
	},
}
