package humanStruct

import "fmt"

type human struct {
	age    int
	name   string
	status string
}

type IHuman interface {
	SetAge(num int) error
	GetAge() int

	SetName(n string)
	GetName() string

	GetStatus() string
}

func NewHumanStruct(ageI int, nameI string) (IHuman, error) {
	hum := human{
		name: nameI,
	}
	err := hum.SetAge(ageI)
	hum.determineStatus(hum.age)
	return &hum, err
}

func (h *human) SetAge(num int) error {
	if num > 0 && num < 150 {
		h.age = num
	} else {
		return fmt.Errorf("invalid age")
	}
	return nil
}

func (h human) GetAge() int {
	return h.age
}

func (h *human) SetName(n string) {
	h.name = n
}

func (h human) GetName() string {
	return h.name
}

func (h human) GetStatus() string {
	return h.status
}

func (h *human) determineStatus(age int) {
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

func CalculateAverageAge(humanSlice []IHuman) float32 {
	var result int = 0
	for _, v := range humanSlice {
		result += v.GetAge()
	}
	return float32(result) / float32(len(humanSlice))
}

func TryAdd(targetSlice *[]IHuman, hum IHuman) bool {
	for _, v := range *targetSlice {
		if hum.GetAge() == v.GetAge() && hum.GetName() == v.GetName() {
			return false
		}
	}
	*targetSlice = append(*targetSlice, hum)
	return true
}
