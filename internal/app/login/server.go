package login

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.RemoteAddr)

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)

	log.Println(b)

	if err != nil {
		log.Fatal(err)
		return
	}

	var v map[string]interface{}

	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(v)

	bodyRes, err := json.Marshal(v)

	log.Println(bodyRes)
	log.Println(err)

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(bodyRes)

	return
}

func Run() {

	log.Println("login.RUn...")
	http.HandleFunc("/login", LoginHandler)

	err := http.ListenAndServe("localhost:8002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
