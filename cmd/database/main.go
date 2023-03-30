package main

import (
	"flag"
	"log"

	"github.com/bagusyanuar/go-deltomed/cmd/database/scheme"
	"github.com/bagusyanuar/go-deltomed/config"
)

func main() {
	configuration := config.New()
	database := config.NewDatabaseConnection(configuration)
	seed := flag.String("m", "", "unsupport command")
	flag.Parse()
	command := *seed
	switch command {
	case "seed":
		scheme.Seed(database)
		return
	case "migrate":
		scheme.Migrate(database)
		return
	default:
		log.Println("unknown command")
		return
	}
	// if command == "seed" {
	// 	scheme.Seed(database)
	// 	return
	// } else {
	// 	return
	// }
	// scheme.Migrate(database)
}
