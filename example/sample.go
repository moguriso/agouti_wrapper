package main

import (
	"log"

	aw "github.com/moguriso/agouti_wrapper"
)

func main() {
	driver, err := aw.GetWebDriver()
	defer driver.Stop()
	if err != nil {
		log.Fatal(err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}
	page.Navigate("https://www.yahoo.co.jp")
	page.Screenshot("./sc.png")
}
