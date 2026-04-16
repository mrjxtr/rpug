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

// Generator is the interface for generating Pinoy data.
type Generator interface {
	Generate(int, string) (*PinoyResponse, error)
}

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
	seed string
	rnd  *mathrand.Rand
}

// NewPinoyGenerator creates a new PinoyGenerator.
func NewPinoyGenerator(cfg *config.Config, d *data.Data) *PinoyGenerator {
	return &PinoyGenerator{
		cfg:  cfg,
		data: d,
		seed: "",
		rnd:  &mathrand.Rand{},
	}
}

// Generate creates a PinoyResponse with n Pinoy records.
// Use if seedParam as seed if available, if not, we generate
func (g *PinoyGenerator) Generate(
	resParam int,
	seedParam string,
) (*PinoyResponse, error) {
	if seedParam == "" {
		seed, err := generateSeed()
		if err != nil {
			return nil, err
		}
		g.seed = seed
		g.rnd = newRNGfromSeed(seed)
	} else {
		g.seed = seedParam
		g.rnd = newRNGfromSeed(seedParam)
	}

	results := g.generatePinoys(resParam)
	info, err := g.generateInfo(results)
	if err != nil {
		return &PinoyResponse{}, err
	}

	return &PinoyResponse{
		Results: results,
		Info:    info,
	}, nil
}

// generatePinoys creates n Pinoy records. Placeholder for now.
func (g *PinoyGenerator) generatePinoys(n int) *[]Pinoy {
	pinoys := make([]Pinoy, n)
	for i := range pinoys {
		var p Pinoy

		nameList := g.data.Names
		lastNameList := nameList.LastNames
		titleList := nameList.Titles
		locationList := g.data.Locations[g.rnd.IntN(len(g.data.Locations))]

		globeTM := g.data.MobileProviders.GlobeTM
		smartTntSun := g.data.MobileProviders.SmartTntSun
		dito := g.data.MobileProviders.Dito

		providerList := append(append(globeTM, smartTntSun...), dito...)

		// ? NOTE: Randomize gender based on seed
		// ? Then generate the title, first name, and last name based on gender and seed
		if g.rnd.IntN(2) == 0 {
			p.Gender = "male"
			p.Name.Title = titleList.Male[g.rnd.IntN(len(titleList.Male))]
			p.Name.First = nameList.MaleFirstNames[g.rnd.IntN(len(nameList.MaleFirstNames))]
		} else {
			p.Gender = "female"
			p.Name.Title = titleList.Female[g.rnd.IntN(len(titleList.Female))]
			p.Name.First = nameList.FemaleFirstNames[g.rnd.IntN(len(nameList.FemaleFirstNames))]
		}
		p.Name.Last = lastNameList[g.rnd.IntN(len(lastNameList))]

		// ? NOTE: Generate a random Age from a configurable referenceDate
		// ? Then we derive the DOB from the age based on seed
		referenceDate := time.Date(g.cfg.ReferenceDate, 1, 1, 0, 0, 0, 0, time.UTC)
		age := g.rnd.IntN(42) + 18 // 18-60 years old
		dob := referenceDate.AddDate(-age, -g.rnd.IntN(12), -g.rnd.IntN(28))

		p.DOB.Age = age
		p.DOB.Date = dob.Format(time.RFC3339)

		// TODO: Support more locations
		// ? NOTE: Grab a random city based on seed
		selectedCity := locationList.Cities[g.rnd.IntN(len(locationList.Cities))]

		p.Location.City = selectedCity.Name
		p.Location.Region = locationList.Region
		p.Location.Country = "Philippines"
		p.Location.Zipcode = selectedCity.Zipcode

		prefix := providerList[g.rnd.IntN(len(providerList))]
		suffix := fmt.Sprintf("%07d", g.rnd.IntN(10000000))
		p.Phone = prefix + suffix

		// ? NOTE: Create a generic email from first and last name
		// ? Remove whitespace since names can have multiple words (e.g., "Maria Clara", "Dela Cruz")
		firstName := strings.ReplaceAll(p.Name.First, " ", "")
		lastName := strings.ReplaceAll(p.Name.Last, " ", "")

		p.Email = strings.ToLower(
			fmt.Sprintf("%s.%s@gmail.com", firstName, lastName),
		)

		// ? NOTE: Generate random regestration age and date based on seed
		regAge := g.rnd.IntN(5) // within 5 years
		regDage := time.Now().AddDate(regAge, -g.rnd.IntN(12), -g.rnd.IntN(28))

		p.Registered.Age = regAge
		p.Registered.Date = regDage.Format(time.RFC3339)

		pinoys[i] = p
	}

	return &pinoys
}

// generateInfo fills the response metadata based on n. Placeholder for now.
func (g *PinoyGenerator) generateInfo(results *[]Pinoy) (Info, error) {
	return Info{
		Seed:    g.seed,
		Results: len(*results),
		// TODO: Implement pagination
		Version: g.cfg.Version,
	}, nil
}
