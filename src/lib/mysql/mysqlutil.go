package mysqlutil

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/what-the-fake/src/lib/models"
)

func GetSites() ([]*models.SiteData, error) {

	db := GetDbConnection()

	rows, err := db.Query("select * from websitedata.website where sitename like '%site%'")

	sites := make([]*models.SiteData, 0)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		site := new(models.SiteData)
		err := rows.Scan(&site.Id)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	defer db.Close()

	return sites, nil
}

func GetDbConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:suhumar123@tcp(127.0.0.1:3306)/websitedata")

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Ping failed")
		panic(pingErr.Error())
	}

	// if there is an error opening the connection, handle it
	if err != nil {

		fmt.Println("error")
		panic(err.Error())
	}

	if err == nil {
		fmt.Println("Making a connection with mysql")
	}

	return db

}
