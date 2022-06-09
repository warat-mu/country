package models

import (
	"database/sql"

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

func GetCountry() []Data {

	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}
	defer db.Close()

	result, err := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata;")
	defer result.Close()

	var DataList []Data
	for result.Next() {
		var country Data
		err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)
		if err != nil {
			panic(err.Error())
		}
		DataList = append(DataList, country)
	}
	return DataList
}

func GetCountryDESC() []Data {

	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}
	defer db.Close()
	result, _ := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata ORDER BY Population DESC;")
	defer result.Close()

	var DataList []Data
	for result.Next() {
		var country Data
		err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)

		if err != nil {
			panic(err.Error())
		}
		DataList = append(DataList, country)
	}
	return DataList
}

func GetCountryASC() []Data {

	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}
	defer db.Close()
	result, _ := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata ORDER BY Population ASC;")
	defer result.Close()

	var DataList []Data
	for result.Next() {
		var country Data
		err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)

		if err != nil {
			panic(err.Error())
		}
		DataList = append(DataList, country)
	}
	return DataList
}

func GetCountryByReion(region string) []Data {

	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
	if err != nil {
		return nil
	}
	defer db.Close()

	result, _ := db.Query("SELECT Name,Region,Population,Flag_png,Currencies_Name FROM countrydata WHERE Region = ?;", region)
	defer result.Close()

	var DataList []Data
	for result.Next() {
		var country Data
		err := result.Scan(&country.Name, &country.Region, &country.Population, &country.Flag_png, &country.Currencies_Name)
		if err != nil {
			panic(err.Error())
		}
		DataList = append(DataList, country)
	}
	return DataList
}

// func Test() {
// 	db, err := sql.Open("mysql", "tester:secret@tcp(host.docker.internal:3306)/api")
// 	if err != nil {
// 		log.Println("connect fail")
// 	} else {
// 		log.Println("connect success")
// 	}
// 	defer db.Close()
// }
