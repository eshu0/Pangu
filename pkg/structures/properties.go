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

func (p *Property) PrintString() string {
	switch p.GType {
	case GTint64:
		return "\n\tstr += fmt.Sprintf(\" %d \",data." + p.Name + ")"
	case GTstring:
		return "\n\tstr += fmt.Sprintf(\" %s \",data." + p.Name + ")"
	case GTfloat64:
		return "\n\tstr += fmt.Sprintf(\" %f \",data." + p.Name + ")"
	default:
		return "\n\tstr += fmt.Sprintf(\" %v \",data." + p.Name + ")"
	}
}
