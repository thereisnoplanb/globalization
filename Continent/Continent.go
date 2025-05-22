package Continent

type Enum int

const (
	Africa       Enum = 1
	Asia         Enum = 2
	Europe       Enum = 3
	NorthAmerica Enum = 4
	SouthAmerica Enum = 5
	Australia    Enum = 6
	Antarctica   Enum = 7
)

var names map[Enum]string = map[Enum]string{
	Africa:       "Africa",
	Asia:         "Asia",
	Europe:       "Europe",
	NorthAmerica: "North America",
	SouthAmerica: "South America",
	Australia:    "Australia",
	Antarctica:   "Antartica",
}

func (enum Enum) String() string {
	if name, ok := names[enum]; ok {
		return name
	}
	return "Unknown"
}

func IsDefined(enum Enum) bool {
	_, ok := names[enum]
	return ok
}

func (enum Enum) IsDefined() bool {
	_, ok := names[enum]
	return ok
}

func GetValues() []Enum {
	return []Enum{
		Africa,
		Asia,
		Europe,
		NorthAmerica,
		SouthAmerica,
		Australia,
		Antarctica,
	}
}
