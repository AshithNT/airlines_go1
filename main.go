package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type YourJson []struct {
	Airport struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}
	Time struct {
		Label      string `json:"label"`
		Month      int    `json:"month"`
		Month_Name string `json:"month_name"`
		Year       int    `json:"year"`
	}
	Statistics struct {
		Carriers struct {
			Names string
			Total int
		}
		Flights struct {
			Cancelled int
			Delayed   int
			Diverted  int
			On_Time   int
			Total     int
		}
		Minutes_Delayed struct {
			Carrier                  int
			Late_Aircraft            int
			National_Aviation_System string
			Security                 int
			Total                    int
			Weather                  int
		}
	}
}

// Social struct which contains a
// list of links

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var json_data YourJson

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &json_data)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	// for i := 0; i < len(users.Users); i++ {
	// 	fmt.Println("User Type: " + users.Users[i].Type)
	// 	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
	// 	fmt.Println("User Name: " + users.Users[i].Name)
	// 	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }

	json.NewEncoder(w).Encode(json_data)
}
func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		jsonFile, err := os.Open("users.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully Opened users.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)

		// we initialize our Users array
		var json_data YourJson
		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteValue, &json_data)

		// we iterate through every user within our users array and
		// print out the user Type, their name, and their facebook url
		// as just an example
		code := r.URL.Query().Get("code")
		fmt.Println(code)
		length := len(json_data)
		for i := 0; i < length; i++ {
			if json_data[i].Airport.Code == string(code) {

				fmt.Println(json_data[i])
				json.NewEncoder(w).Encode(json_data[i])

			}
		}
	})
	http.ListenAndServe(":3000", nil)

}
