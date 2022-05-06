package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/tjarratt/babble"
)

type field[T any] struct {
	name  string
	value T
}

func generateMedication(dataLength int) error {

	isSubsidised := field[int]{
		name: "isSubsidised",
	}

	scientificName := field[string]{
		name: "scientific_name",
	}

	brand := field[string]{
		name: "brand",
	}

	cost := field[float64]{
		name: "cost",
	}

	lines := make([]string, dataLength)

	babbler := babble.NewBabbler()
	babbler.Count = 1

	min := 0.0
	max := 100.0

	for i := 0; i < dataLength; i++ {
		isSubsidised.value = rand.Intn(2)
		scientificName.value = strings.Replace(babbler.Babble(), "'", "", -1)
		brand.value = strings.Replace(babbler.Babble(), "'", "", -1)
		cost.value = min + rand.Float64()*(max-min)

		line := fmt.Sprintf("INSERT INTO medication (%s, %s, %s, %s) VALUES (%d, '%s', '%s', %.2f)\n",
			isSubsidised.name, scientificName.name, brand.name, cost.name,
			isSubsidised.value, scientificName.value, brand.value, cost.value)

		lines[i] = line
	}

	err := writeFile("medication.sql", lines)
	if err != nil {
		return fmt.Errorf("Error generating medication: %s", err)
	}
	return nil
}

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

func main() {
	generateMedication(10)
}
