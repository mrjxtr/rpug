// Package generator
package generator

import (
	"encoding/json"
	"fmt"
	"log/slog"
	mathrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mrjxtr/rpug/internal/config"
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
	seed string
	rnd  *mathrand.Rand
}

// NewPinoyGenerator creates a new PinoyGenerator.
func NewPinoyGenerator(cfg *config.Config) *PinoyGenerator {
	return &PinoyGenerator{
		cfg:  cfg,
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
		if err == nil {
			g.seed = seed
			g.rnd = newRNGfromSeed(seed)
		}
	} else {
		g.seed = seedParam
		g.rnd = newRNGfromSeed(seedParam)
	}

	results, err := g.generatePinoys(resParam)
	if err != nil {
		return &PinoyResponse{}, err
	}

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
// TODO: Implement
func (g *PinoyGenerator) generatePinoys(n int) (*[]Pinoy, error) {
	slog.Info("Generating Pinoy records", "n", n)
	data, err := readDataFromJSON()
	if err != nil {
		return nil, err
	}

	pinoys := make([]Pinoy, n)
	for i := range pinoys {
		var p Pinoy

		nameList := data.Names
		lastNameList := nameList.LastNames
		titleList := nameList.Titles
		locationList := data.Locations[g.rnd.Intn(len(data.Locations))]

		// ? NOTE: Randomize gender based on seed
		// ? Then generate the title, first name, and last name based on gender and seed
		if g.rnd.Intn(2) == 0 {
			p.Gender = "male"
			p.Name.Title = titleList.Male[g.rnd.Intn(len(titleList.Male))]
			p.Name.First = nameList.MaleFirstNames[g.rnd.Intn(len(nameList.MaleFirstNames))]
		} else {
			p.Gender = "female"
			p.Name.Title = titleList.Female[g.rnd.Intn(len(titleList.Female))]
			p.Name.First = nameList.FemaleFirstNames[g.rnd.Intn(len(nameList.FemaleFirstNames))]
		}
		p.Name.Last = lastNameList[g.rnd.Intn(len(lastNameList))]

		// ? NOTE: Generate a random Age from a configurable referenceDate
		// ? Then we derive the DOB from the age based on seed
		referenceDate := time.Date(g.cfg.ReferenceDate, 1, 1, 0, 0, 0, 0, time.UTC)
		age := g.rnd.Intn(42) + 18 // 18-60 years old
		dob := referenceDate.AddDate(-age, -g.rnd.Intn(12), -g.rnd.Intn(28))

		p.DOB.Age = age
		p.DOB.Date = dob.Format(time.RFC3339)

		// TODO: Support more locations
		// ? NOTE: Grab a random city based on seed
		selectedCity := locationList.Cities[g.rnd.Intn(len(locationList.Cities))]

		p.Location.City = selectedCity.Name
		p.Location.Region = locationList.Region
		p.Location.Country = "Philippines"
		p.Location.Zipcode = selectedCity.Zipcode

		// TODO: Support more providers
		p.Phone = "0909" + strconv.Itoa(g.rnd.Intn(9999999))

		// ? NOTE: Create a generic email from first and last name
		firstName := p.Name.First
		lastName := p.Name.Last

		p.Email = strings.ToLower(
			fmt.Sprintf("%s.%s@gmail.com", firstName, lastName),
		)

		// ? NOTE: Generate random regestration age and date based on seed
		regAge := g.rnd.Intn(5) // within 5 years
		regDage := time.Now().AddDate(regAge, -g.rnd.Intn(12), -g.rnd.Intn(28))

		p.Registered.Age = regAge
		p.Registered.Date = regDage.Format(time.RFC3339)

		pinoys[i] = p
	}

	return &pinoys, nil
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

// Data mirrors data.json.
// Fields are exported and tagged for json.
type data struct {
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

// readDataFromJSON returns the decode data from json
func readDataFromJSON() (*data, error) {
	fileName := "data/data.json"
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var d data
	if err := json.NewDecoder(file).Decode(&d); err != nil {
		return nil, err
	}

	return &d, nil
}
