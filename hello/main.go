package main

import (
	// Your imported package
	"hello/entities"
)

func main() {

	entities.ReadConfigFile("/Users/anilkumar.patel/go/src/hello/config.json")
	entities.LoadTemplates()

	//defer entities.Close()
	//entities.Connect()

	/* start: Dispatch crons */
	//go entities.MetricCachingCron()
	/* end: Dispatch crons */

	entities.Start()
}
