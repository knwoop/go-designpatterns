package builderWithDirector

type BuildProcess interface {
	SetStructure() BuildProcess
	SetCamera() BuildProcess
	SetMonitor() BuildProcess
	GetGadget() ElectronicProduct
}