package main

import (
	"fmt"
	"log"
	"os"

	"github.com/what-the-fake/src/lib/models"
	"github.com/what-the-fake/src/lib/repo"
)

func main() {

	// Get the "PORT" env variable
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	sites, err := getAllSites("%.com%")

	for _, site := range sites {
		fmt.Println(site.Domain)
	}

	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}

}

func getAllSites(name string) ([]*models.SiteData, error) {

	sites, err := repo.GetSites(name)

	return sites, err
}
