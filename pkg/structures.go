package pangu

// Go lang Constant
type Constant struct {
	Name string 
	Gtype string
	Comment string
	Value string
}

// Go Lang Struct
type StructDetails struct {
	Name string
	Properties []*Property
	Comment string

	// These indicate what properties can be updated 
	// optimisation for my use cae might remove if not needed in general
	UpdateProperties []*Property

	// The property that is an Identifier for the struct
	// this might need to be handled differently
	Id *Property
}

// A Go lang struct property
type Property struct {
	Name string
	GType string
	Json string
	Comment string
	IsIdentifier bool

	// This links the property to a constant
	// so if a constant should set this etc
	Constant *Constant
}