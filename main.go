package living

import "github.com/cyan21/greeting/v2"

func Dit() string {
	return "Le renard dit :'" + greeting.Fr("Pierre") + "'"
}

func Says() string {
	return "The bear says " + greeting.En("Rick") + "'"
}

func Iu() string {
	return "Inu wa " + greeting.Jp("Yasu") + "da to iimasu"
}
