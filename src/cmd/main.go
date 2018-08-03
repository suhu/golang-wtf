package main

import (
	"fmt"

	"github.com/what-the-fake/src/lib/models"
	"github.com/what-the-fake/src/lib/repo"
)

func main() {

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
