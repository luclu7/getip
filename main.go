package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPAPI struct {
	IPVersion string `json:"ip_version"`
	IP        string `json:"ip"`
	Hostname  string `json:"hostname"`
	Isp       string `json:"as_org"`
}

func main() {
	resp, err := http.Get("https://anjara.eu/ip/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var ip IPAPI
	json.Unmarshal(body, &ip)

	fmt.Printf("%s: %s (%s)\n", ip.IPVersion, ip.IP, ip.Hostname)
	fmt.Printf("From %s\n", ip.Isp)
}
