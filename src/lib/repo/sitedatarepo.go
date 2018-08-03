package repo

import (
	"github.com/what-the-fake/src/lib/models"
	"github.com/what-the-fake/src/lib/mysql"
)

func GetSites(name string) ([]*models.SiteData, error) {

	db := mysqlutil.GetDbConnection()

	rows, err := db.Query(`SELECT Id,
		Domain,
		Name,
		Organization,
		Address,
		Ifnull(Postalcode,'') as Postalcode,
		Ifnull(Country,'') as Country,
		Ifnull(Email,'') as Email,
		Ifnull(IpAddress,'') as IpAddress,
		Ifnull(Phone,'') as Phone,
		Ifnull(Phone2,'') as Phone2,
		Ifnull(Fax,'') as Fax,
		Ifnull(Outcome,'') as Outcome		
		FROM websitedata.website where domain like ?`, name)

	sites := make([]*models.SiteData, 0)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		site := new(models.SiteData)
		err := rows.Scan(
			&site.Id,
			&site.Domain,
			&site.Name,
			&site.Organization,
			&site.Address,
			&site.Postalcode,
			&site.Country,
			&site.Email,
			&site.IPAddress,
			&site.Phone,
			&site.Phone2,
			&site.Fax,
			&site.Outcome)

		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	defer db.Close()

	return sites, nil

}
