package ipapi

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

// Response from ipapi.com
//
// https://ipapi.co/api/#complete-location
type Response struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryCode        string  `json:"country_code"`
	CountryCodeISO3    string  `json:"country_code_iso3"`
	CountryName        string  `json:"country_name"`
	CountryCapital     string  `json:"country_capital"`
	CountryTLD         string  `json:"country_tld"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  float64 `json:"country_population"`
	ContinentCode      string  `json:"continent_code"`
	InEU               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	LatLong            string  `json:"latlong"`
	Timezone           string  `json:"timezone"`
	UTCOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	ASN                string  `json:"asn"`
	Org                string  `json:"org"`
	Hostname           string  `json:"hostname"`
}

// GetLocation from ipapi.com with given ip address.
func GetLocation(ip string) (response Response, err error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			IdleConnTimeout:       10 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 10 * time.Second,
		},
	}

	var req *http.Request
	if req, err = http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/json/", ip), nil); err == nil {
		var resp *http.Response
		resp, err = httpClient.Do(req)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err == nil {
			defer resp.Body.Close()
			if err = json.NewDecoder(resp.Body).Decode(&response); err == nil {
				return response, nil
			}
		}
	}

	return response, err
}
