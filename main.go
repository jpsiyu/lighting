package main

import "github.com/jpsiyu/lighting/cmd"

import "log"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	cmd.Execute()
}
