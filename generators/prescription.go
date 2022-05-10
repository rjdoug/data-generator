package generators


import (
	"fmt"
	"math/rand"
	"github.com/BlaviButcher/data-generator/io"
	"github.com/BlaviButcher/data-generator/generators/helper"
)

func GeneratePrescriptions(dataLength int, patients, practitioners []string) error {

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

	for i := 0; i < dataLength; i++ {
		patientUsername.value = patients[rand.Intn(len(patients))]
		practitionerUsername.value = practitioners[rand.Intn(len(practitioners))]
		notes.value = helper.GetBabble(10)

		line := fmt.Sprintf("INSERT INTO prescription (%s, %s, %s) VALUES ('%s', '%s', '%s')\n",
			patientUsername.name, practitionerUsername.name, notes.name,
			patientUsername.value, practitionerUsername.value, notes.value)

		lines[i] = line
	}

	err := io.WriteFile("sql_scripts/prescriptions.sql", lines)
	if err != nil {
		return fmt.Errorf("generating prescriptons: %s", err)
	}
	return nil

}
