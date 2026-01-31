package chaos

type Store interface {
	Find(path, method string) (*Config, bool)
	Create(cfg *Config) string
	List() []*Config
	Delete(id string) bool
}
