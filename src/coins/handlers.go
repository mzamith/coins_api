package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getRate(w http.ResponseWriter, r *http.Request) {

	var conversion Conversion

	parameters := r.URL.Query()

	conversion.ConvertFrom = parameters["from"][0]
	conversion.ConvertTo = parameters["to"][0]

	t := getTax(&conversion)

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(t))

}

func getTotalConversion(w http.ResponseWriter, r *http.Request) {

	//1. create a conversion Slice and a con Slice
	var coins Currencies

	//2. Get the value to convert to, from the query parameter in the URL
	to := r.URL.Query()["to"][0]

	fmt.Println("Gotta convert stuff to ", to)
	fmt.Println("")

	//3. Read the body of the request, which will be a JSON array of Currency values
	// and convert it to our Currency object
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &coins); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//4. For each Currency Value, we have to connect to Yahoo's server and get our conversion rate
	//Then, get the rate and calculate the values for each coin
	var total float64
	for i := 0; i < len(coins); i++ {

		conversion := Conversion{coins[i].coin, to}

		t := getTax(&conversion)
		rate := strings.Split(t, ",")[1]

		c, err := strconv.ParseFloat(rate, 64)
		if err != nil {
			panic(err)
		}

		fmt.Println("Rate for "+coins[i].coin+" is ", c)
		total = total + c*coins[i].value
	}

	fmt.Println("---------------------")
	fmt.Println("Total: ", total)
	fmt.Println("---------------------")

	//5, Set headers for response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//6. Create and send response
	response := Currency{to, total}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
