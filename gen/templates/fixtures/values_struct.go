package fixtures

type CityValues struct{}

func (_ *CityValues) Values() []string {
	return []string{
		"london",
		"oakland",
		"portland",
		"seattle",
		"San Francisco",
		`"`,
		"sekret",
	}
}
