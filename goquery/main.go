package main

import (
	"github.com/denimY/go_test/goquery/scrapper"
	"github.com/labstack/echo"
	"os"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}
func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := scrapper.CleanString(c.FormValue("term"))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)

}

func main() {
	//scrapper.Scrape("term")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

}
