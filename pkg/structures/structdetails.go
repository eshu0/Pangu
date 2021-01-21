package pangudata

//StructDetails Go Lang Struct
type StructDetails struct {
	Name       string
	Properties []*Property
	Comment    string

	// The property that is an Identifier for the struct
	// this might need to be handled differently
	ID *Property

	//Functions for the structure
	Functions []*Function
}
