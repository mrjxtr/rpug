// Package data handles data structures
package data

// Data mirrors data.json structure
type Data struct {
	Names           Names           `json:"names"`
	Locations       []Location      `json:"locations"`
	MobileProviders MobileProviders `json:"mobile_providers"`
}

type Names struct {
	Titles           Titles   `json:"titles"`
	MaleFirstNames   []string `json:"male_first_names"`
	FemaleFirstNames []string `json:"female_first_names"`
	LastNames        []string `json:"last_names"`
}

type Titles struct {
	Male   []string `json:"male"`
	Female []string `json:"female"`
}

type Location struct {
	Region string `json:"region"`
	Cities []City `json:"cities"`
}

type City struct {
	Name    string `json:"name"`
	Zipcode string `json:"zipcode"`
}

type MobileProviders struct {
	GlobeTM     []string `json:"globe_tm"`
	SmartTntSun []string `json:"smart_tnt_sun"`
	Dito        []string `json:"dito"`
}
