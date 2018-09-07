package sdk

// HostType 主机类型
type HostType struct {
	Name string
	CPU  int
	MEM  float64
}

// NewHostType new host type
func NewHostType(name string, cpu int, mem float64) HostTyper {
	return &HostType{
		Name: name,
		CPU:  cpu,
		MEM:  mem,
	}
}

// String name
func (t *HostType) String() string {
	return t.Name
}

// NumCPU implement HostTyper interface
func (t *HostType) NumCPU() int {
	return t.CPU
}

// NumMEM implement HostTyper interface
func (t *HostType) NumMEM() float64 {
	return t.MEM
}

// HostTyper 主机类型接口
type HostTyper interface {
	String() string
	NumCPU() int
	NumMEM() float64
}
