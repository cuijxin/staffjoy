package main

type sendingConfig struct {
	WhitelistOnly bool
	Concurrency   int
	// Country code to national numbers
	Numbers map[int32]string
}

var sendingConfigs = map[string]sendingConfig{
	"development": {
		WhitelistOnly: true,
		Numbers:       map[int32]string{1: "13840490189"},
		Concurrency:   1,
	},
	"staging": {
		WhitelistOnly: true,
		Numbers:       map[int32]string{1: "13840490189"},
		Concurrency:   1,
	},
	"production": {
		WhitelistOnly: false,
		Numbers: map[int32]string{
			1:  "13840490189",
			44: "13840490189",
		},
		Concurrency: 1,
	},
}