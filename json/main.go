package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User : user struct
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// Quote : quote struct
type Quote struct {
	Img    string `json:"img"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func main() {

	// Decoding json to struct, when getting data from thrid party
	http.HandleFunc("/get_quote", func(w http.ResponseWriter, r *http.Request) {
		var quote Quote
		resp, err := http.Get("http://www.quotesapi.ml/random")
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(data, &quote)

		json.NewEncoder(w).Encode(quote)

	})

	// Decodig data from
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user.Firstname, user.Lastname, user.Age)
	})

	http.ListenAndServe(":4000", nil)
}
