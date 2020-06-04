package pangu

type Constant struct {
	Name string 
	Gtype string
	Comment string
	Value string
}

type StructDetails struct {
	Name string
	Properties []*Property
	UpdateProperties []*Property
	Comment string
	Id *Property
}

type Property struct {
	Name string
	GType string
	Json string
	Comment string
	IsIdentifier bool
	Constant *Constant
}