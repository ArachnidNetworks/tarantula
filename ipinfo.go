package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type IP struct {
  Address   string `json:"ip"`
  Hostname  string `json:"hostname"`

  City      string `json:"city"`
  Region    string `json:"region"`
  Country   string `json:"country"`

  Loc       string `json:"loc"`
  Org       string `json:"org"`
}

func Lookup(address string) string {
	spaceClient := http.Client {
		Timeout: time.Second * 2, // 2 second timeout
	}

	req, err := http.NewRequest(http.MethodGet, "http://ipinfo.io/" + address, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "curl") // Set UA to curl so ipinfo will respond with JSON

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	addr := IP{}
	jsonErr := json.Unmarshal(body, &addr) // Parse response as JSON
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

  return fmt.Sprintf(">*%s* _(%s)_\n>%s\n>%s %s %s %s", addr.Address,
                                             addr.Hostname,
																						 addr.Org,
                                             addr.Loc,
                                             addr.City,
                                             addr.Region,
                                             addr.Country)

}
