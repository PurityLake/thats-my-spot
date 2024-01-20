package data

type Property struct {
	Name, PropType string
	Value          any
}

func NewProperty(name, propType string, value any) *Property {
	return &Property{
		Name:     name,
		PropType: propType,
		Value:    value,
	}
}
