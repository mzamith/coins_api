package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const mainUrl string = "http://download.finance.yahoo.com/d/quotes?f=sl1d1t1&s="

func getTax(conversion *Conversion) string {

	url := mainUrl + conversion.ConvertFrom + conversion.ConvertTo + "=X"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return ""
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	response := string(body)

	fmt.Println("Response From YAHOO: ", response)

	return response
}
