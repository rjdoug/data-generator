package generators


import (
	"fmt"
	"math/rand"
	"github.com/BlaviButcher/data-generator/io"
	"github.com/BlaviButcher/data-generator/generators/helper"
)

func GenerateContains(dataLength, perscriptionLength, medicationLength int) error {

	const maxMG = 1000
	const maxRepeats = 5

	perscriptionID := Field[int]{
		name: "perscription_id",
	}

	medicationID := Field[int]{
		name: "medication_id",
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
		perscriptionID.value = rand.Intn(perscriptionLength)
		medicationID.value = rand.Intn(medicationLength)
		dosageMG.value = rand.Intn(maxMG)
		instruction.value = helper.GetBabble(10)
		repeats.value = rand.Intn(maxRepeats)

		line := fmt.Sprintf("INSERT INTO contains (%s, %s, %s, %s, %s) VALUES (%d, %d, %d, '%s', %d)\n",
			perscriptionID.name, medicationID.name, dosageMG.name, instruction.name, repeats.name,
			perscriptionID.value, medicationID.value, dosageMG.value, instruction.value, repeats.value)

		lines[i] = line
	}

	err := io.WriteFile("sql_scripts/contains.sql", lines)
	if err != nil {
		return fmt.Errorf("generating contains relationship: %s", err)
	}
	return nil

}
