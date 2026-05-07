// Package generator
package generator

import (
	"fmt"
	mathrand "math/rand/v2"
	"strings"
	"time"

	"github.com/mrjxtr/rpug/internal/config"
	"github.com/mrjxtr/rpug/internal/data"
)

const (
	minAge               = 18
	maxAge               = 60
	phoneSuffixMax       = 10_000_000 // 7-digit suffix, matches "%07d" below
	maxRegistrationYears = 5
)

// TODO: Populate more Pinoy data

type Pinoy struct {
	Name struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	DOB struct {
		Date string `json:"date"`
		Age  int    `json:"age"`
	} `json:"dob"`
	Location struct {
		// Street struct {
		// 	Number int    `json:"number"`
		// 	Name   string `json:"name"`
		// } `json:"street"`
		City    string `json:"city"`
		Region  string `json:"region"`
		Country string `json:"country"`
		Zipcode string `json:"zipcode"`
	} `json:"location"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	// Login  struct {
	// 	UUID     string `json:"uuid"`
	// 	Username string `json:"username"`
	// 	Password string `json:"password"`
	// } `json:"login"`
	Registered struct {
		Date string `json:"date"`
		Age  int    `json:"age"`
	} `json:"registered"`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page,omitempty"`
	Version string `json:"version,omitempty"`
}

type PinoyResponse struct {
	Results *[]Pinoy `json:"results"`
	Info    Info     `json:"info"`
}

type PinoyGenerator struct {
	cfg  *config.Config
	data *data.Data
}

// NewPinoyGenerator creates a new PinoyGenerator.
func NewPinoyGenerator(cfg *config.Config, d *data.Data) *PinoyGenerator {
	return &PinoyGenerator{
		cfg:  cfg,
		data: d,
	}
}

// Generate creates a PinoyResponse with n Pinoy records.
// Use seedParam as seed if available, otherwise generate one.
// The RNG is local to this call, so Generate is safe for concurrent use.
func (g *PinoyGenerator) Generate(
	resParam int,
	seedParam string,
) (*PinoyResponse, error) {
	seed := seedParam
	if seed == "" {
		s, err := generateSeed()
		if err != nil {
			return nil, err
		}
		seed = s
	}
	rng := newRNGfromSeed(seed)

	results := g.generatePinoys(resParam, rng)
	info, err := g.generateInfo(results, seed)
	if err != nil {
		return &PinoyResponse{}, err
	}

	return &PinoyResponse{
		Results: results,
		Info:    info,
	}, nil
}

// generatePinoys creates n Pinoy records using the provided RNG.
func (g *PinoyGenerator) generatePinoys(n int, rng *mathrand.Rand) *[]Pinoy {
	pinoys := make([]Pinoy, n)

	globeTM := g.data.MobileProviders.GlobeTM
	smartTntSun := g.data.MobileProviders.SmartTntSun
	dito := g.data.MobileProviders.Dito

	providerList := append(append(globeTM, smartTntSun...), dito...)

	nameList := g.data.Names
	lastNameList := nameList.LastNames
	titleList := nameList.Titles

	locations := g.data.Locations

	now := time.Now()
	referenceDate := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)

	for i := range pinoys {
		var p Pinoy

		locationList := locations[rng.IntN(len(locations))]

		// ? NOTE: Randomize gender based on seed
		// ? Then generate the title, first name, and last name based on gender and seed
		if rng.IntN(2) == 0 {
			p.Gender = "male"
			p.Name.Title = titleList.Male[rng.IntN(len(titleList.Male))]
			p.Name.First = nameList.MaleFirstNames[rng.IntN(len(nameList.MaleFirstNames))]
		} else {
			p.Gender = "female"
			p.Name.Title = titleList.Female[rng.IntN(len(titleList.Female))]
			p.Name.First = nameList.FemaleFirstNames[rng.IntN(len(nameList.FemaleFirstNames))]
		}
		p.Name.Last = lastNameList[rng.IntN(len(lastNameList))]

		// ? NOTE: Generate a random Age from a configurable referenceDate
		// ? Then we derive the DOB from the age based on seed
		age := rng.IntN(maxAge-minAge) + minAge
		dob := referenceDate.AddDate(-age, -rng.IntN(12), -rng.IntN(28))

		p.DOB.Age = age
		p.DOB.Date = dob.Format(time.RFC3339)

		// TODO: Support more locations
		// ? NOTE: Grab a random city based on seed
		selectedCity := locationList.Cities[rng.IntN(len(locationList.Cities))]

		p.Location.City = selectedCity.Name
		p.Location.Region = locationList.Region
		p.Location.Country = "Philippines"
		p.Location.Zipcode = selectedCity.Zipcode

		prefix := providerList[rng.IntN(len(providerList))]
		suffix := fmt.Sprintf("%07d", rng.IntN(phoneSuffixMax))
		p.Phone = prefix + suffix

		// ? NOTE: Create a generic email from first and last name
		// ? Remove whitespace since names can have multiple words (e.g., "Maria Clara", "Dela Cruz")
		firstName := strings.ReplaceAll(p.Name.First, " ", "")
		lastName := strings.ReplaceAll(p.Name.Last, " ", "")

		p.Email = strings.ToLower(
			fmt.Sprintf("%s.%s@gmail.com", firstName, lastName),
		)

		// ? NOTE: Generate random regestration age and date based on seed
		regAge := rng.IntN(maxRegistrationYears)
		regDage := now.AddDate(regAge, -rng.IntN(12), -rng.IntN(28))

		p.Registered.Age = regAge
		p.Registered.Date = regDage.Format(time.RFC3339)

		pinoys[i] = p
	}

	return &pinoys
}

// generateInfo fills the response metadata based on n.
func (g *PinoyGenerator) generateInfo(results *[]Pinoy, seed string) (Info, error) {
	return Info{
		Seed:    seed,
		Results: len(*results),
		// TODO: Implement pagination
		Version: g.cfg.Version,
	}, nil
}
