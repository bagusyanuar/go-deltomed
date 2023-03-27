package main

import (
	"github.com/bagusyanuar/go-deltomed/cmd/database/scheme"
	"github.com/bagusyanuar/go-deltomed/config"
)

func main() {
	configuration := config.New()
	database := config.NewDatabaseConnection(configuration)
	scheme.Migrate(database)
}
