package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/sethvargo/go-password/password"
	"github.com/tjarratt/babble"
)

var babbler = babble.NewBabbler()

const userCount int = 700
const practitionerCount int = 10
const prescriptionCount = 1000
const medicationCount = 10000

func writeFile(file string, lines []string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("")
	}
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line)

		if err != nil {
			return fmt.Errorf("writing to file: %v", err)
		}
	}
	fmt.Printf("%s written successfully\n", file)

	return nil
}

func ReadTextFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading file: %v", err)
	}

	return lines, nil
}

func GetBabble(wordCount int) string {
	babbler.Count = wordCount
	babbler.Separator = " "
	return strings.Replace(babbler.Babble(), "'", "", -1)
}

func GetPassword() (string, error) {

	// chracters to be replaced
	r := strings.NewReplacer(
		"\"", "",
		"'", "",
		"`", "",
	)

	// generate password 16 characters long with 4 digits and symbols allow lower and upper and repeat characters
	pwd, err := password.Generate(16, 4, 4, false, true)
	if err != nil {
		return "", fmt.Errorf("generating password: %v", err)
	}
	return r.Replace(pwd), err
}

// GetRandDateString returns a random date between the start and end dates
func GetRandDateString(startYear, endYear int) string {
	min := time.Date(startYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(endYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	t := time.Unix(sec, 0).String()

	// clip off time zone and return (probably should turn to array and manipulate, but this is fine for now)
	return t[:len(t)-11]
}

func main() {
	// set babble number of words to generate per call

	if err := GenerateMedication(medicationCount); err != nil {
		log.Fatal(err)
	}

	usernames, err := GenerateUser(userCount)
	if err != nil {
		log.Fatal(err)
	}

	practitioners := usernames[:practitionerCount]
	patients := usernames[practitionerCount:]

	if err := GeneratePractioners(practitioners); err != nil {
		log.Fatal(err)
	}

	if err := GeneratePatients(patients); err != nil {
		log.Fatal(err)
	}

	if err := GenerateAppointments(1000, patients, practitioners); err != nil {
		log.Fatal(err)
	}

	if err := GeneratePrescriptions(1000, patients, practitioners); err != nil {
		log.Fatal(err)
	}

	if err := GenerateContains(1000, prescriptionCount, medicationCount); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")

}
