package main

import (
	"fmt"
	"log"

	gen "github.com/BlaviButcher/data-generator/generators"
	"github.com/BlaviButcher/data-generator/io"
)

const userCount int = 20
const practitionerCount int = 10
const prescriptionCount = 10
const medicationCount = 10
const appointmentCount = 10
const containsCount = prescriptionCount

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

	if err := gen.GenerateAppointments(appointmentCount, patients, practitioners); err != nil {
		log.Fatal(err)
	}

	prescriptionUsers, err := gen.GeneratePrescriptions(prescriptionCount, patients, practitioners)
	if err != nil {
		log.Fatal(err)
	}

	if err := gen.GenerateContains(containsCount, prescriptionCount, medicationCount, prescriptionUsers); err != nil {
		log.Fatal(err)
	}

	fmt.Println("COMBINING FILES...")
	if err := combineAndMinify(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")

}

func combineAndMinify() error {
	var combinedFile []string

	users, err := io.ReadFile("sql_scripts/users.sql")
	if err != nil {
		return fmt.Errorf("joining files %v", err)
	}
	users = append(users, "\n")

	patients, err := io.ReadFile("sql_scripts/patient.sql")
	if err != nil {
		return fmt.Errorf("joining files %v", err)
	}
	patients = append(patients, "\n")

	practitioners, err := io.ReadFile("sql_scripts/practitioner.sql")
	if err != nil {
		return fmt.Errorf("joining files %v", err)
	}
	practitioners = append(practitioners, "\n")

	appointments, err := io.ReadFile("sql_scripts/appointment.sql")
	if err != nil {
		return fmt.Errorf("joining files %v", err)
	}
	appointments = append(appointments, "\n")

	medications, err := io.ReadFile("sql_scripts/medication.sql")
	if err != nil {
		return fmt.Errorf("joining files %v", err)
	}
	medications = append(medications, "\n")

	prescriptions, err := io.ReadFile("sql_scripts/prescription.sql")
	if err != nil {
		return fmt.Errorf("joining files %v", err)
	}
	prescriptions = append(prescriptions, "\n")

	contains, err := io.ReadFile("sql_scripts/contains.sql")
	if err != nil {
		return fmt.Errorf("joining files %v", err)
	}
	contains = append(contains, "\n")

	combinedFile = append(combinedFile, users...)
	combinedFile = append(combinedFile, patients...)
	combinedFile = append(combinedFile, practitioners...)
	combinedFile = append(combinedFile, appointments...)
	combinedFile = append(combinedFile, medications...)
	combinedFile = append(combinedFile, prescriptions...)
	combinedFile = append(combinedFile, contains...)

	if err := io.WriteFile("sql_scripts/combined.sql", combinedFile); err != nil {
		return fmt.Errorf("writing to combined file %v", err)
	}

	return nil

}
