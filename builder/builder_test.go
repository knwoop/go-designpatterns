package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	carBuilder := NewCarBuilder()
	car := carBuilder.TopSpeed(50).Paint(BLUE).Build()

	carSpeedExpect := "Driving at speed: 50"
	driveResult    := car.Drive()
	if driveResult != carSpeedExpect {
		t.Errorf("Expect output to %s, but %s\n", carSpeedExpect, driveResult)
	}

	elfKingExpect := "Stopping a blue car"
	stopResult    := car.Stop()
	if stopResult != elfKingExpect {
		t.Errorf("Expect output to %s, but %s\n", elfKingExpect, stopResult)
	}
}