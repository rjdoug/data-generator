package helper

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sethvargo/go-password/password"
	"github.com/tjarratt/babble"
)

var babbler = babble.NewBabbler()

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
