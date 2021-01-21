package pangudata

//Property A Go lang struct property
type Property struct {
	Name    string
	GType   GType
	Json    string
	Comment string
	// Properties are written out like this
	// <Name> <GType> <Json> <Comment>
	// Examples:
	// FirstName string 'json:fname' // the usersname
	// Age int

	IsIdentifier bool

	// This links the property to a constant
	// so if a constant should set this etc
	Constant *Constant

	UpdateValue string
}
