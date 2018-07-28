package repo

import (
	"github.com/what-the-fake/src/lib/models"
	"github.com/what-the-fake/src/lib/mysql"
)

func GetSites(name string) ([]*models.SiteData, error) {

	db := mysqlutil.GetDbConnection()
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
