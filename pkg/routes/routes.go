package routes

import (
	"log"
	"net/http"

	"github.com/xjart2014x/benchmarkitinterview/pkg/service"
)

var mc service.MedicalClient

func InitiateRoutes() {
	router := http.NewServeMux()

	router.HandleFunc("/getPatients", mc.GetPatients) 
	router.HandleFunc("/getPatients/{tin}", mc.GetPatient)
	router.HandleFunc("/getMedicalDiseases", mc.GetMedicalDiseases)
	router.HandleFunc("/getMedicalDiseases/{id}", mc.GetMedicalDisease)

	http.ListenAndServe(":3000", router)

	log.Println("Listening on port 3000")
}
