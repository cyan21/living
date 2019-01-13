package living

import "github.com/cyan21/greeting"

func Dit() string {
	return "Le renard dit :'" + greeting.Fr() + "'"
}

func Says() string {
	return "The bear says " + greeting.En() + "'"
}

func Iu() string {
	return "Inu wa " + greeting.Jp() + "da to iimasu"
}
