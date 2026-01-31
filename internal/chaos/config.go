package chaos

type Config struct {
	ID         string
	Path       string
	Method     string
	ErrorRate  float64
	MinDelayMs int
	MaxDelayMs int
	RateLimit  int
}
