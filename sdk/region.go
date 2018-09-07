package sdk

// Regioner region interface
type Regioner interface {
	String() string
}

// Region region
type Region string

func (reg Region) String() string {
	return string(reg)
}
