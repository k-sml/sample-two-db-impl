package main

import "github.com/sethvargo/go-envconfig"

func main() {
	var envConfig env.Config
	if err := envconfig.Process("", &envConfig); err != nil {
		panic(err)
	}
}
