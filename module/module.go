package module

type Test struct {
	ID   int    `validate:"max=2" json:"id"`
	Name string `validate:"required" json:"name"`
}
