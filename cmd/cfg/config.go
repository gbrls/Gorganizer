package cfg

import (
	"bufio"
	"io"
)

//Config holds data extracted from the configuration file.
type Config struct {
	//Exposing data coz why not
	Data map[string]string
}

//GetFolder is used to de code be clearer and more explicit.
func (c Config) GetFolder(s string) string {
	return c.Data[s]
}

//NewConfig parses the configuration file for the *main* function.
func NewConfig(f io.Reader) *Config {

	cfg := &Config{}
	cfg.Data = make(map[string]string)

	scn := bufio.NewScanner(f)
	for scn.Scan() {

		//TODO: parse more complex tokens like: (a,b) = c
		txt := scn.Text()
		cfg.Data[txt] = txt
	}

	return cfg
}
