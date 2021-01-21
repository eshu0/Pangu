package pangudata

type GType int

const (
	North GType = iota
	East
	South
	West
)

func (d GType) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
