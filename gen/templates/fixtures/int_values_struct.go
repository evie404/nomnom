package fixtures

type NumberValues struct{}

func (_ *NumberValues) Values() []int {
	return []int{
		1,
		2,
		3,
	}
}
