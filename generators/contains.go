package generators

import (
	"fmt"
	"math/rand"

	"github.com/BlaviButcher/data-generator/generators/helper"
	"github.com/BlaviButcher/data-generator/io"
)

func GenerateContains(dataLength, perscriptionLength, medicationLength int, prescriptionUsers []User) error {

	const maxMG = 1000
	const maxRepeats = 5

	prescriptionID := Field[int]{
		name: "prescription_id",
	}

	medicationID := Field[int]{
		name: "medication_id",
	}

	patientUsername := Field[string]{
		name: "patient_username",
	}

	practitionerUsername := Field[string]{
		name: "practitioner_username",
	}

	dosageMG := Field[int]{
		name: "dosage_mg",
	}

	instruction := Field[string]{
		name: "instruction",
	}

	repeats := Field[int]{
		name: "repeats",
	}

	lines := make([]string, dataLength)

	for i := 0; i < dataLength; i++ {

		prescriptionIndex := rand.Intn(len(prescriptionUsers))

		prescriptionID.value = prescriptionIndex + 1
		medicationID.value = rand.Intn(medicationLength) + 1
		patientUsername.value = prescriptionUsers[prescriptionIndex].patient
		practitionerUsername.value = prescriptionUsers[prescriptionIndex].practitioner
		dosageMG.value = rand.Intn(maxMG)
		instruction.value = helper.GetBabble(10)
		repeats.value = rand.Intn(maxRepeats)

		line := fmt.Sprintf("INSERT INTO contains (%s, %s, %s, %s, %s, %s, %s) VALUES (%d, %d, '%s', '%s', %d, '%s', %d);\n",
			prescriptionID.name, medicationID.name, patientUsername.name, practitionerUsername.name, dosageMG.name, instruction.name, repeats.name,
			prescriptionID.value, medicationID.value, patientUsername.value, practitionerUsername.value, dosageMG.value, instruction.value, repeats.value)

		lines[i] = line
	}

	err := io.WriteFile("sql_scripts/contains.sql", lines)
	if err != nil {
		return fmt.Errorf("generating contains relationship: %s", err)
	}
	return nil

}
