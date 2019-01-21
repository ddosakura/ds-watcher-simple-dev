package main

import (
	"log"
)

func er(msg interface{}) {
	log.Panicln(msg)
}
