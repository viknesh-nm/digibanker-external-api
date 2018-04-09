package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	soap "github.com/viknesh-nm/digibanker-external-api/api"
)

type DigiSecurity struct {
	*soap.Client
}

// Init inits the Digisecurity client api
func Init() *DigiSecurity {
	return &DigiSecurity{
		soap.NewClient(),
	}
}

// BankListHandler -
func (api *DigiSecurity) BankListHandler(w http.ResponseWriter, r *http.Request) {
	creds := soap.BankListRequest{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	a, _ := api.GetListOfBanks(creds)

	JSONW(w, a)

}

// PayBankHandler -
func (api *DigiSecurity) PayBankHandler(w http.ResponseWriter, r *http.Request) {
	creds := soap.PaymentRequest{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data, _ := api.PayBank(creds)

	// this has to be verifed for complete checking after the credentials
	JSONW(w, data)

}

// JSONW writes JSON response to the given writer
func JSONW(w http.ResponseWriter, data interface{}) {
	d, err := json.MarshalIndent(data, "", "  ")
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(d)
	if nil != err {
		return
	}
	_, err = w.Write([]byte("\n"))
	if nil != err {
		return
	}
}

func main() {
	handle := Init()
	r := mux.NewRouter()
	r.HandleFunc("/list", handle.BankListHandler).Methods("POST")
	r.HandleFunc("/pay_bank", handle.PayBankHandler).Methods("POST")

	http.ListenAndServe(":9080", r)

}
