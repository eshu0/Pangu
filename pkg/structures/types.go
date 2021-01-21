package pangudata

type GType int

const (
	GTunknown GType = iota
	GTint64
	GTstring
	GTfloat64
)

func (d GType) String() string {
	return [...]string{"unknown", "int64", "string", "float64"}[d]
}
