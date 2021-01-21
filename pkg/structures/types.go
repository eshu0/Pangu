package pangudata

type GType int

const (
	gtunknown GType = iota
	gtint64
	gtstring
	gtfloat64
)

func (d GType) String() string {
	return [...]string{"unknown", "int64", "string", "float64"}[d]
}

func (d GType) PropertyString() string {
	switch p.GType {
	case gtint64:
		String.Data += "\n\tstr += fmt.Sprintf(\" %d \",data." + p.Name + ")"
	case gtstring:
		String.Data += "\n\tstr += fmt.Sprintf(\" %s \",data." + p.Name + ")"
		break
	case gtfloat64:
		String.Data += "\n\tstr += fmt.Sprintf(\" %f \",data." + p.Name + ")"
		break
	default:
		String.Data += "\n\tstr += fmt.Sprintf(\" %v \",data." + p.Name + ")"
		break
	}
}
