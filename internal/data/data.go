// Package data handles data structures
package data

// Data mirrors data.json structure
type Data struct {
	Names struct {
		Titles struct {
			Male   []string `json:"male"`
			Female []string `json:"female"`
		} `json:"titles"`
		MaleFirstNames   []string `json:"male_first_names"`
		FemaleFirstNames []string `json:"female_first_names"`
		LastNames        []string `json:"last_names"`
	} `json:"names"`
	Locations []struct {
		Region string `json:"region"`
		Cities []struct {
			Name    string `json:"name"`
			Zipcode string `json:"zipcode"`
		} `json:"cities"`
	} `json:"locations"`
}
