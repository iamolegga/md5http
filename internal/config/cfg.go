package config

import (
	"flag"
)

var parallel uint

func init() {
	flag.UintVar(&parallel, "parallel", 10, "max parallel requests")
}

type Config struct {
	Parallel uint
	Inputs []string
}

func New() (cfg Config) {
	flag.Parse()
	cfg.Parallel = parallel
	cfg.Inputs = flag.Args()
	return cfg
}
