package builderWithDirector

import "fmt"

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() ElectronicProduct {
	m.builder.SetStructure().SetMonitor().SetCamera()
	return m.builder.GetGadget()
}

func (m *ManufacturingDirector) PrintProduct()  {
	gadget := m.builder.GetGadget()
	fmt.Printf("Structure: %s \n", gadget.Structure)
	fmt.Printf("Monitor: %d \n", gadget.Monitor)
	fmt.Printf("Camera: %d \n", gadget.Camera)
	fmt.Printf("===============\n")
}