package addDB

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Country struct {
	Name       string       `json:"name"`
	Subregion  string       `json:"subregion"`
	Region     string       `json:"region"`
	Population int          `json:"population"`
	Flags      Flags        `json:"flags"`
	Currencies []Currencies `json:"currencies"`
	Flag       string       `json:"flag"`
}

type Currencies struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Flags struct {
	Svg string `json:"svg"`
	Png string `json:"png"`
}

func DBcheck(db *sql.DB) {
	var check int
	result := db.QueryRow("SELECT EXISTS (SELECT 1 FROM api.countrydata);")
	errs := result.Scan(&check)
	if errs != nil {
		panic(errs.Error())
	}
	fmt.Println(check)
	if check == 0 {
		NewCountry(db)
	}

}

func NewCountry(db *sql.DB) {

	resp, err := http.Get("https://restcountries.com/v2/all")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var result []Country
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(err.Error())
	}

	for _, rec := range result {
		insert, err := db.Query("INSERT INTO api.countrydata (Name,Region,Population,Flag_png,Flag_svg,Currencies_Code,Currencies_Name,Currencies_Symbol) VALUES (?,?,?,?,?,?,?,?)", rec.Name, countryRegion(rec.Region, rec.Subregion), rec.Population, rec.Flags.Png, rec.Flags.Svg, getcurrenciesCode(rec.Currencies), getcurrenciesName(rec.Currencies), getcurrenciesSymbol(rec.Currencies))
		if err != nil {
			fmt.Println(err.Error())
		}
		insert.Close()
	}

}

func countryRegion(region string, subregion string) string {
	if region == "Antarctic Ocean" || region == "Polar" {
		return "Antarctic"
	}
	if region == "Americas" {
		if subregion == "South America" {
			return "South America"
		}
		return "North Antarctic"
	}

	return region
}

func getcurrenciesCode(Currencies []Currencies) string {
	if len(Currencies) == 0 {
		return ""
	}
	if len(Currencies) == 1 {
		return Currencies[0].Code
	}
	var str []string
	for _, rec := range Currencies {
		str = append(str, rec.Code)
	}
	return strings.Join(str, ",")
}

func getcurrenciesName(Currencies []Currencies) string {
	if len(Currencies) == 0 {
		return ""
	}
	if len(Currencies) == 1 {
		return Currencies[0].Name
	}
	var str []string
	for _, rec := range Currencies {
		str = append(str, rec.Name)
	}
	return strings.Join(str, ",")
}

func getcurrenciesSymbol(Currencies []Currencies) string {
	if len(Currencies) == 0 {
		return ""
	}
	if len(Currencies) == 1 {
		return Currencies[0].Symbol
	}
	var str []string
	for _, rec := range Currencies {
		str = append(str, rec.Symbol)
	}
	return strings.Join(str, ",")
}
