package builderWithDirector

type laptop struct {
	electronicProduct ElectronicProduct
}

func NewLaptop() BuildProcess {
	return &laptop{}
}

func (l *laptop) SetStructure() BuildProcess {
	l.electronicProduct.Structure = "Laptop"
	return l
}

func (l *laptop) SetMonitor() BuildProcess {
	l.electronicProduct.Monitor = 1
	return l
}

func (l *laptop) SetCamera() BuildProcess {
	l.electronicProduct.Camera = 1
	return l
}

func (l *laptop) GetGadget() ElectronicProduct {
	return l.electronicProduct
}