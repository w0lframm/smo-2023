package pkg

type Config struct {
	Mode    string `yaml:"mode"`
	Time    int    `yaml:"time"`
	Sources int    `yaml:"source"`
	Devices int    `yaml:"device"`
	Bufsize int    `yaml:"bufsize"`
}
