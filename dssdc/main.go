//go:generate statik -src=./assets -f
//go:generate go fmt statik/statik.go

package main

import (
	"fmt"
	"log"

	"./cmd"
)

func main() {
	fmt.Println(logo)
	log.SetPrefix("[DDoSakura]: ")
	// log.SetFlags(2)
	cmd.Ver(version)
	cmd.Execute()
}
