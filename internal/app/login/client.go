package login

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func RunClient() {
	url := "http://localhost:8000/login"

	//resp, err:= http.Get("http://localhost:8000/login")

	j := `{"key":"value"}`
	r := strings.NewReader(j)
	resp, err := http.Post(url, "application/json", r)

	log.Println(resp)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(b)

	var v map[string]interface{}

	json.Unmarshal(b, &v)

	log.Println(v)

}
