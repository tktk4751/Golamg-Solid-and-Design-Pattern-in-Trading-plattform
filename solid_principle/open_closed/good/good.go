package good

import (
	"fmt"
	"math"
)

type IEmplotee interface {
	name() string
	getBonus(base float64) float64
}

type Junior struct {
	Name string
}

func NewJunior(name string) *Junior {
	return &Junior{name}
}

func (e *Junior) name() string {
	return e.Name
}

func (e *Junior) getBonus(base float64) float64 {
	return math.Floor(base * 1.1)
}

type Middle struct {
	Name string
}

func (e *Middle) name() string {
	return e.Name
}

func (e *Middle) getBonus(base float64) float64 {
	return math.Floor(base * 1.5)
}

type Senior struct {
	Name string
}

func NewSenior(name string) *Senior {
	return &Senior{name}
}

func (e *Senior) name() string {
	return e.Name
}

func (e *Senior) getBonus(base float64) float64 {
	return math.Floor(base * 2)
}

func Run() {

	var base float64 = 100.0

	var emp1 IEmplotee = NewJunior("takashi")
	var emp2 IEmplotee = &Middle{Name: "daikichi"}
	var emp3 IEmplotee = NewSenior("daikichi")

	fmt.Printf("Name:%s Bonus:%v\n", emp1.name(), emp1.getBonus(base))
	fmt.Printf("Name:%s Bonus:%v\n", emp2.name(), emp2.getBonus(base))
	fmt.Printf("Name:%s Bonus:%v\n", emp3.name(), emp3.getBonus(base))
	// fmt.Println(emp2.getBonus(base))
	// fmt.Println(emp3.getBonus(base))
}
