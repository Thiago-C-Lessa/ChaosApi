package chaos

type Config struct {
	ID         string
	Path       string  `json:"path"`
	Method     string  `json:"method"`
	ErrorRate  float64 `json:"error_rate"`
	MinDelayMs int     `json:"min_delay_ms"`
	MaxDelayMs int     `json:"max_delay_ms"`
	RateLimit  int
}
