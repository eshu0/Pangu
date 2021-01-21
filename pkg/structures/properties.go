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
	case gtint64:
		return "\n\tstr += fmt.Sprintf(\" %d \",data." + p.Name + ")"
	case gtstring:
		return "\n\tstr += fmt.Sprintf(\" %s \",data." + p.Name + ")"
	case gtfloat64:
		return "\n\tstr += fmt.Sprintf(\" %f \",data." + p.Name + ")"
	default:
		return "\n\tstr += fmt.Sprintf(\" %v \",data." + p.Name + ")"
	}
}

func (prop *Property) SetTypeFromCType(ctype CType) {
	switch ctype {
	case Integer:
		prop.GType = gtint64
		prop.UpdateValue = "11"
		break
	case Text:
		prop.GType = gtstring
		prop.UpdateValue = "\"Updated\""
		break
	case VarChar:
		prop.GType = gtstring
		prop.UpdateValue = "\"Updated\""
		break
	case Numeric:
		prop.GType = gtfloat64
		prop.UpdateValue = "1.11"
		break
	}
}
