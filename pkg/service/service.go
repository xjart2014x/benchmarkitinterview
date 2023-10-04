package service

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/xjart2014x/benchmarkitinterview/pkg/dto"
	"github.com/xjart2014x/benchmarkitinterview/pkg/mockdata"
)

var (
	patientInfo  = mockdata.Patients
	trialInfo    = mockdata.TrialInfo
	patientsInfo []dto.PatientInfo
	medicalTrial []dto.MedicineTrial
)

type MedicalClient interface {
	GetPatients(w http.ResponseWriter, r *http.Request)
	GetPatient(w http.ResponseWriter, r *http.Request)
	GetMedicalDiseases(w http.ResponseWriter, r *http.Request)
	GetMedicalDisease(w http.ResponseWriter, r *http.Request)
}

type Client struct {
	HttpClient http.Client
}

func (m *Client) GetPatients(w http.ResponseWriter, r *http.Request) {
	patientInf := []dto.PatientInfo{}
	medicalTri := []dto.MedicineTrial{}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for _, pi := range patientInfo {
			if pi.ActiveInsurance {
				patientInf = append(patientInf, pi)
			}
		}
		wg.Done()
	}()

	go func() {
		for _, mt := range medicalTrial {
			totalTrials := mt.NumberOfFailures + mt.NumberOfSuccesses
			if float64(mt.NumberOfSuccesses/totalTrials) >= 0.9 {
				medicalTri = append(medicalTri, mt)
			}
		}
		wg.Done()
	}()
	wg.Wait()

	output, err := json.MarshalIndent(medicalTrial, "", " ")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Write(output)
	w.WriteHeader(http.StatusOK)
}

func (m *Client) GetPatient(w http.ResponseWriter, r *http.Request){
	output, err := json.MarshalIndent(medicalTrial, "", " ")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(output)
	w.WriteHeader(http.StatusOK)
}

func (m *Client) GetMedicalDisease(w http.ResponseWriter, r *http.Request) {
	medicalTri := []dto.MedicineTrial{}

	for _, mt := range medicalTrial {
		totalTrials := mt.NumberOfFailures + mt.NumberOfSuccesses
		if float64(mt.NumberOfSuccesses/totalTrials) >= 0.9 {
			medicalTri = append(medicalTri, mt)
		}
	}

	output, err := json.MarshalIndent(medicalTri, "", " ")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(output)
	w.WriteHeader(http.StatusOK)
}

func (m *Client) GetMedicalDiseases(w http.ResponseWriter, r *http.Request) {
	filteredMT := []dto.MedicineTrial{}

	for _, mt := range medicalTrial {
		totalTrials := mt.NumberOfFailures + mt.NumberOfSuccesses
		if float64(mt.NumberOfSuccesses/totalTrials) >= 0.9 {
			filteredMT = append(filteredMT, mt)
		}
	}

	output, err := json.MarshalIndent(filteredMT, "", " ")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(output)
	w.WriteHeader(http.StatusOK)
}
