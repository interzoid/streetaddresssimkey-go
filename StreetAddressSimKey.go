// Provides a hashed similarity key from the input data used to match with other similar street address data. 
// Use the generated similarity key, rather than the actual data itself, to match and/or sort street
// address data by similarity. This avoids the problems of data inconsistency, misspellings, and address variations
// when matching within a single dataset, and can also help matching across datasets or for more advanced searching.

package StreetAddressSimKey

// visit www.interzoid.com to obtain the required API license key - free trial available

import (
	"encoding/json"
	"errors"
	"net/http"
	url2 "net/url"
)


// used to retrieve the data payload from the Interzoid API
type payload struct {
	SimKey string  // generated similarity key
	Code string  // success or fail
	Credits string // credits remaining for this API license key
}

// this function takes a company name and an API license key and returns the generated similarity key
func GetSimKey(license,address string) (string, string, string, error) {
  response, err := http.Get("https://api.interzoid.com/getaddressmatchadvanced?license="+url2.QueryEscape(license)+"&address="+url2.QueryEscape(address)+"&algorithm=narrow")
  
  if err != nil || response.StatusCode != 200 {
  	switch response.StatusCode {
		case 403 : response.Status = "Invalid API license key. Register at www.interzoid.com"
		case 400 : response.Status = "Invalid parameters"
	}
  	return "","Fail","0",errors.New(response.Status)
  }

  thePayload := payload{}
  json.NewDecoder(response.Body).Decode(&thePayload)

  return thePayload.SimKey,thePayload.Code,thePayload.Credits,nil
}

