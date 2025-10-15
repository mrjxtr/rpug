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
	Generate(int) (*PinoyResponse, error)
	GenerateWithSeed(string) (*PinoyResponse, error)
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
		// Zipcode string `json:"zipcode"`
	} `json:"location"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	// Login  struct {
	// 	UUID     string `json:"uuid"`
	// 	Username string `json:"username"`
	// 	Password string `json:"password"`
	// } `json:"login"`
	// Registered struct {
	// 	Date string `json:"date"`
	// 	Age  int    `json:"age"`
	// } `json:"registered"`
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
		cfg,
		"",
		&mathrand.Rand{},
	}
}

// Generate creates a PinoyResponse with n Pinoy records.
func (g *PinoyGenerator) Generate(
	resParam int,
) (*PinoyResponse, error) {
	seed, err := generateSeed()
	if err == nil {
		g.seed = seed
		g.rnd = newRNGfromSeed(seed)
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

// GenerateWithSeed creates a PinoyResponse with seed.
func (g *PinoyGenerator) GenerateWithSeed(seed string) (*PinoyResponse, error) {
	return nil, nil
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

		lastNameList := data.Names.LastNames
		locationList := data.Locations[g.rnd.Intn(len(data.Locations))]

		// pick gender first, then choose matching title and first name
		if g.rnd.Intn(2) == 0 {
			p.Gender = "male"
			p.Name.Title = data.Names.Titles.Male[g.rnd.Intn(len(data.Names.Titles.Male))]
			p.Name.First = data.Names.MaleFirstNames[g.rnd.Intn(len(data.Names.MaleFirstNames))]
		} else {
			p.Gender = "female"
			p.Name.Title = data.Names.Titles.Female[g.rnd.Intn(len(data.Names.Titles.Female))]
			p.Name.First = data.Names.FemaleFirstNames[g.rnd.Intn(len(data.Names.FemaleFirstNames))]
		}
		p.Name.Last = lastNameList[g.rnd.Intn(len(lastNameList))]

		referenceDate := time.Date(g.cfg.ReferenceDate, 1, 1, 0, 0, 0, 0, time.UTC)
		age := g.rnd.Intn(42) + 18
		dob := referenceDate.AddDate(-age, -g.rnd.Intn(12), -g.rnd.Intn(28))

		p.DOB.Age = age
		// format is "1989-05-30T23:07:31.851Z"
		p.DOB.Date = dob.Format(time.RFC3339)

		p.Location.Region = locationList.Region
		p.Location.City = locationList.Cities[g.rnd.Intn(len(locationList.Cities))]
		p.Location.Country = "Philippines"

		// TODO: Support more providers
		p.Phone = "0909" + strconv.Itoa(g.rnd.Intn(9999999))

		p.Email = strings.ToLower(
			fmt.Sprintf("%s.%s@gmail.com", p.Name.First, p.Name.Last),
		)

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

// Data mirrors data.json. Fields are exported and tagged for json.
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
		Region string   `json:"region"`
		Cities []string `json:"cities"`
	} `json:"locations"`
}

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
