package bad

import (
	"fmt"
	"math"
)

type Grade int

const (
	Junior Grade = iota
	Middle
	Senior
	Expart
)

type Employee struct {
	name  string
	grade Grade
}

func NewEmployee(name string, grade Grade) *Employee {
	return &Employee{name, grade}
}

type BonusCalcrator struct {
	base float64
}

func NewBonusCalcrator(base float64) *BonusCalcrator {
	return &BonusCalcrator{base}
}

func (b *BonusCalcrator) getBonus(employee Employee) float64 {

	if employee.grade == Junior {
		return math.Floor(b.base * 1.1)
	} else if employee.grade == Middle {
		return math.Floor(b.base * 1.5)
	} else if employee.grade == Senior {
		return math.Floor(b.base * 2)
	} else {
		return math.Floor(b.base * 3)
	}

}

func Run() {
	employee1 := NewEmployee("takashi", Junior)
	employee2 := NewEmployee("daikichi", Middle)
	employee3 := NewEmployee("migel", Senior)

	bonusCalc := NewBonusCalcrator(100.0)

	fmt.Println(bonusCalc.getBonus(*employee1))
	fmt.Println(bonusCalc.getBonus(*employee2))
	fmt.Println(bonusCalc.getBonus(*employee3))
}
