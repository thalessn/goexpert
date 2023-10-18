package main

import "api/config"

func main() {
	conf, _ := config.LoadConfig(".")
	println(conf.DBDriver)
}
