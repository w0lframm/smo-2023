package pkg

type Config struct {
	Mode    string `yaml:"mode"`
	Time    uint64 `yaml:"time"`
	Sources uint64 `yaml:"source"`
	Devices uint64 `yaml:"device"`
	Bufsize uint64 `yaml:"bufsize"`
}
