package builderWithDirector

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	manufacturingDirector := ManufacturingDirector{}

	laptop := NewLaptop()
	manufacturingDirector.SetBuilder(laptop)
	laptopPro := manufacturingDirector.Construct()
	structureExpect1 := "Laptop"
	structureResult1 := laptopPro.Structure
	if structureResult1 != structureExpect1 {
		t.Errorf("Expect output to %s, but %s\n", structureExpect1, structureResult1)
	}
	cameraExpect1 := 1
	cameraResult1 := laptopPro.Camera
	if cameraResult1 != cameraExpect1 {
		t.Errorf("Expect output to %v, but %v\n", cameraExpect1, cameraResult1)
	}
	monitorExpect1 := 1
	monitorResult1 := laptopPro.Monitor
	if monitorResult1 != monitorExpect1 {
		t.Errorf("Expect output to %v, but %v\n", monitorExpect1, monitorResult1)
	}

	smartphone := NewSmartphone()
	manufacturingDirector.SetBuilder(smartphone)
	smartphonePro := manufacturingDirector.Construct()

	structureExpect2 := "Smartphone"
	structureResult2 := smartphonePro.Structure
	if structureResult2 != structureExpect2 {
		t.Errorf("Expect output to %s, but %s\n", structureExpect2, structureResult2)
	}
	cameraExpect2 := 2
	cameraResult2 := smartphonePro.Camera
	if cameraResult2 != cameraExpect2 {
		t.Errorf("Expect output to %v, but %v\n", cameraExpect2, cameraResult2)
	}
	monitorExpect2 := 1
	monitorResult2 := smartphonePro.Monitor
	if monitorResult2 != monitorExpect2 {
		t.Errorf("Expect output to %v, but %v\n", monitorExpect2, monitorResult2)
	}
}