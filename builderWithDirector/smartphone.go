package builderWithDirector

type smartphone struct {
	electronicProduct ElectronicProduct
}

func NewSmartphone() BuildProcess {
	return &smartphone{}
}

func (s *smartphone) SetStructure() BuildProcess {
	s.electronicProduct.Structure = "Smartphone"
	return s
}

func (s *smartphone) SetMonitor() BuildProcess {
	s.electronicProduct.Monitor = 1
	return s
}

func (s *smartphone) SetCamera() BuildProcess {
	s.electronicProduct.Camera = 2
	return s
}

func (s *smartphone) GetGadget() ElectronicProduct {
	return s.electronicProduct
}