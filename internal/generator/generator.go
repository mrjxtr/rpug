// Package generator
package generator

import (
	"log/slog"
	mathrand "math/rand"
	"os"

	"github.com/mrjxtr/rpug/internal/config"
)

// Generator is the interface for generating Pinoy data.
type Generator interface {
	Generate(string, int) (*PinoyResponse, error)
}

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
		Street struct {
			Number int    `json:"number"`
			Name   string `json:"name"`
		} `json:"street"`
		City    string `json:"city"`
		Region  string `json:"region"`
		Country string `json:"country"`
		Zipcode string `json:"zipcode"`
	} `json:"location"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Login  struct {
		UUID     string `json:"uuid"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"login"`
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
		cfg,
		"",
		&mathrand.Rand{},
	}
}

// Generate creates a PinoyResponse with n Pinoy records.
func (g *PinoyGenerator) Generate(
	seedParam string,
	resParam int,
) (*PinoyResponse, error) {
	if seedParam != "" {
		g.seed = seedParam
	}

	if seed, err := generateSeed(); err == nil {
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

// generatePinoys creates n Pinoy records. Placeholder for now.
func (g *PinoyGenerator) generatePinoys(n int) (*[]Pinoy, error) {
	// TODO: Implement
	slog.Info("Generating Pinoy records", "n", n)
	data, _ := readDataFromJSON()
	results := &[]Pinoy{*data}

	return results, nil
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

func readDataFromJSON() (*Pinoy, error) {
	fileName := "data/data.json"
	file, err := os.Open(fileName)
	if err != nil {
		slog.Warn("Error while opening file", "error", err)
	}
	defer file.Close()

	testData := &Pinoy{}
	testData.Name.Title = "Mr"
	testData.Name.First = "John"
	testData.Name.Last = "Doe"
	testData.DOB.Date = "1990-01-01"
	testData.DOB.Age = 20
	testData.Location.Street.Number = 123
	testData.Location.Street.Name = "Main St"
	testData.Location.City = "New York"
	testData.Location.Region = "NY"
	testData.Location.Country = "USA"
	testData.Location.Zipcode = "10001"
	testData.Gender = "male"
	testData.Phone = "+1234567890"
	testData.Email = "john.doe@example.com"
	testData.Login.UUID = "1234567890"
	testData.Login.Username = "john.doe"
	testData.Login.Password = "password"
	testData.Registered.Date = "2021-01-01"
	testData.Registered.Age = 20

	return testData, nil
}
