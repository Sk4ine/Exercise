package humanStruct

import "fmt"

type human struct {
	age    int
	name   string
	status string
}

func (h *human) SetAge(num int) error {
	if num > 0 && num < 150 {
		h.age = num
	} else {
		return fmt.Errorf("invalid age")
	}
	return nil
}

func NewHumanStruct(ageI int, nameI string) (human, error) {
	hum := human{
		name: nameI,
	}
	err := hum.SetAge(ageI)
	hum.setStatus(hum.age)
	return hum, err
}

func (h *human) setStatus(age int) {
	if age < 12 {
		h.status = "Kid"
	} else if age >= 12 && age < 18 {
		h.status = "Teenager"
	} else if age >= 18 && age < 50 {
		h.status = "Adult"
	} else if age >= 50 && age < 80 {
		h.status = "Pensioner"
	} else {
		h.status = "Long-liver"
	}
}
