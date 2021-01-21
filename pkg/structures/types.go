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
