package generators

import (
	"fmt"
	"math/rand"

	"github.com/BlaviButcher/data-generator/generators/helper"
	"github.com/BlaviButcher/data-generator/io"
)

type User struct {
	patient      string
	practitioner string
}

func GeneratePrescriptions(dataLength int, patients, practitioners []string) ([]User, error) {

	// This is passed back for use in contains - BAD PRACTICE I know

	patientUsername := Field[string]{
		name: "patient_username",
	}

	practitionerUsername := Field[string]{
		name: "practitioner_username",
	}

	notes := Field[string]{
		name: "notes",
	}

	lines := make([]string, dataLength)
	users := make([]User, dataLength)

	for i := 0; i < dataLength; i++ {
		patientUsername.value = patients[rand.Intn(len(patients))]
		practitionerUsername.value = practitioners[rand.Intn(len(practitioners))]
		notes.value = helper.GetBabble(10)

		line := fmt.Sprintf("INSERT INTO prescription (%s, %s, %s) VALUES ('%s', '%s', '%s');\n",
			patientUsername.name, practitionerUsername.name, notes.name,
			patientUsername.value, practitionerUsername.value, notes.value)

		lines[i] = line

		user := User{patient: patientUsername.value, practitioner: practitionerUsername.value}
		users[i] = user

	}

	err := io.WriteFile("sql_scripts/prescription.sql", lines)
	if err != nil {
		return nil, fmt.Errorf("generating prescriptons: %s", err)
	}
	return users, nil

}
