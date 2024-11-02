package core

type gender string

const (
	MALE   gender = "Male"
	FEMALE gender = "Female"
)

// Data class that contains client information
type ClientInfo struct {
	Name          string
	Gender        gender
	Age           int8
	Height        float64
	Weight        float64
	Smoker        bool
	MedicalIssues bool
}

func NewClient(n string, g gender, a int8, h, w float64, s, m bool) ClientInfo {
	return ClientInfo{
		Name:          n,
		Gender:        g,
		Age:           a,
		Height:        h,
		Weight:        w,
		Smoker:        s,
		MedicalIssues: m,
	}
}
