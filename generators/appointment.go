package generators

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/BlaviButcher/data-generator/io"
)

// CREATE TABLE APPOINTMENT (
// 	patient_username VARCHAR(40) NOT NULL,
// 	practitioner_username VARCHAR(40)NOT NULL,
// start_time TIME NOT NULL,
// end_time Time NOT NULL,
// reason VARCHAR(200) NOT NULL,
// cost NUMBER (50,2) CHECK (cost >= 0),
// is_acc NUMBER (0,1) DEFAULT 0 NOT NULL CHECK (ACC in (0,1)),
// PRIMARY KEY (PatientUsername, PractitionerUsername, startTime)
// );
//

func GenerateAppointments(dataLength int, patients, practitioners []string) error {
	
	const appointmentTime = 30
	const minAppointmentCost = 20.0
	const maxAppointmentCost = 100.0

	patientUsername := Field[string]{
		name: "patient_username",
	}

	practitionerUsername := Field[string]{
		name: "practitioner_username",
	}

	startTime := Field[string]{
		name: "start_time",
	}

	endTime := Field[string]{
		name: "end_time",
	}

	reason := Field[string]{
		name: "reason",
	}

	cost := Field[float64]{
		name: "cost",
	}

	isACC := Field[int]{
		name: "is_acc",
	}

	lines := make([]string, dataLength)

	randomReasons, err := io.ReadTextFile("data/injuries.txt")
	if err != nil {
		return fmt.Errorf("generating appointments: %s", err)
	}

	for i := 0; i < dataLength; i++ {

		patientUsername.value = patients[rand.Intn(len(patients))]

		practitionerUsername.value = practitioners[i%len(practitioners)]

		// approximates what practioner we are working with
		// - used to reduce appointment clashes
		approximateCurrentPractioner := i / len(practitioners)

		startTime.value, endTime.value = GetStartEndTime(appointmentTime, approximateCurrentPractioner)

		reason.value = randomReasons[rand.Intn(len(randomReasons))]

		cost.value = minAppointmentCost + rand.Float64()*(maxAppointmentCost-minAppointmentCost)

		isACC.value = rand.Intn(2)

		line := fmt.Sprintf("INSERT INTO appointment(%s, %s, %s, %s, %s, %s, %s) VALUES('%s', '%s', TO_DATE('%s', 'YYYY/MM/DD HH24:MI:SS'), TO_DATE('%s', 'YYYY/MM/DD HH24:MI:SS'), '%s', '%.2f', %d);\n",
			patientUsername.name, practitionerUsername.name, startTime.name, endTime.name, reason.name, cost.name, isACC.name,
			patientUsername.value, practitionerUsername.value, startTime.value, endTime.value, reason.value, cost.value, isACC.value)

		lines[i] = line
	}

	err = io.WriteFile("sql_scripts/appointment.sql", lines)
	if err != nil {
		return fmt.Errorf("generating medication: %s", err)
	}
	return nil

}

// GetRandDateString returns a start and end dates for appointments
// This isn't random, but systematic enough to avoid clashes
func GetStartEndTime(appointmentWindow, practitionerApproximate int) (string, string) {
	sec := time.Date(1990, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	t := time.Unix(sec, 0)

	// buffer out time based on practititoner approximate - stops doctor from having time clashes
	t = t.Add(time.Duration(appointmentWindow*practitionerApproximate) * time.Minute)

	tStart := t.String()
	tEnd := t.Add(time.Duration(appointmentWindow) * time.Minute).String()

	// clip off time zone and return (probably should turn to array and manipulate, but this is fine for now)
	return tStart[:len(tStart)-11], tEnd[:len(tEnd)-11]
}
