package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Name       string
	Region     string
	Population int
	Flag_png   string
	// Flag_svg          string
	// Currencies_Code   string
	Currencies_Name string
	// Currencies_Symbol string
}

func GetCountry(db *sql.DB) []*Data {
	result, _ := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata;")
	defer result.Close()

  var dataList []*Data
  for result.Next() {
    country := &Data{}
    err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)
    if err != nil {
      log.Println(err.Error())
    }
    dataList = append(dataList, country)
  }
  return dataList
}

func GetCountryDESC(db *sql.DB) []*Data {
	result, _ := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata ORDER BY Population DESC;")
	defer result.Close()

	var dataList []*Data
	for result.Next() {
		country := &Data{}
		err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)

		if err != nil {
			panic(err.Error())
		}
		dataList = append(dataList, country)
	}
	return dataList
}

func GetCountryASC(db *sql.DB) []*Data {

	result, _ := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata ORDER BY Population ASC;")
	defer result.Close()

	var dataList []*Data
	for result.Next() {
		country := &Data{}
		err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)

		if err != nil {
			panic(err.Error())
		}
		dataList = append(dataList, country)
	}
	return dataList
}

func GetCountryByReion(region string, db *sql.DB) []*Data {

	result, _ := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata WHERE Region = ?;", region)
	defer result.Close()

	var dataList []*Data
	for result.Next() {
		country := &Data{}
		err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)
		if err != nil {
			panic(err.Error())
		}
		dataList = append(dataList, country)
	}
	return dataList
}
