package domain

type Port struct {
	ID          string
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}
