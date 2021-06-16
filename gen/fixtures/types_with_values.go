package fixtures

type City string

const (
	// CityLondon is London
	CityLondon  City = "london"
	CityOakland City = "oakland" // CityOakland is Oakland
)

const CityPortland, CitySeattle City = "portland", "seattle"

const CitySanFrancisco City = "San Francisco"

const CityQuotes City = `"`

// Not Supported yet
const CitySanJose City = "San " + "Jose"
const CitySanJose2 City = CitySanJose

func (c City) Stuff() {}

type Number int

const (
	NumberOne   Number = 1
	NumberTwo   Number = 2
	NumberThree Number = 3
)
