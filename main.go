package main

import (
	"fmt"
	"log"

	gen "github.com/BlaviButcher/data-generator/generators"
)

const userCount int = 700
const practitionerCount int = 10
const prescriptionCount = 1000
const medicationCount = 10000

func main() {
	// set babble number of words to generate per call

	if err := gen.GenerateMedication(medicationCount); err != nil {
		log.Fatal(err)
	}

	usernames, err := gen.GenerateUser(userCount)
	if err != nil {
		log.Fatal(err)
	}

	practitioners := usernames[:practitionerCount]
	patients := usernames[practitionerCount:]

	if err := gen.GeneratePractioners(practitioners); err != nil {
		log.Fatal(err)
	}

	if err := gen.GeneratePatients(patients); err != nil {
		log.Fatal(err)
	}

	if err := gen.GenerateAppointments(1000, patients, practitioners); err != nil {
		log.Fatal(err)
	}

	if err := gen.GeneratePrescriptions(1000, patients, practitioners); err != nil {
		log.Fatal(err)
	}

	if err := gen.GenerateContains(1000, prescriptionCount, medicationCount); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")

}
