package sdk

// Label label
type Label struct {
	key   string
	value string
}

// Labeler label interface
type Labeler interface {
	Name() string
}
