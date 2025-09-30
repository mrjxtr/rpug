// Package generator
package generator

import "github.com/mrjxtr/rpug/internal/config"

// Generator is the interface for generating Pinoy data.
type Generator interface {
	Generate(n int) (*PinoyResponse, error)
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
	Results []*Pinoy `json:"results"`
	Info    Info     `json:"info"`
}

type PinoyGenerator struct {
	cfg *config.Config
}

// NewPinoyGenerator creates a new PinoyGenerator.
func NewPinoyGenerator(cfg *config.Config) *PinoyGenerator {
	return &PinoyGenerator{}
}

// Generate creates a PinoyResponse with n Pinoy records.
func (g *PinoyGenerator) Generate(n int) (*PinoyResponse, error) {
	results, err := g.generatePinoys(n)
	if err != nil {
		return &PinoyResponse{}, err
	}
	info, err := g.generateInfo(n)
	if err != nil {
		return &PinoyResponse{}, err
	}

	return &PinoyResponse{
		Results: results,
		Info:    info,
	}, nil
}

// generatePinoys creates n Pinoy records. Placeholder for now.
func (g *PinoyGenerator) generatePinoys(n int) ([]*Pinoy, error) {
	// TODO: Implement
	return make([]*Pinoy, 0, n), nil
}

// generateInfo fills the response metadata based on n. Placeholder for now.
func (g *PinoyGenerator) generateInfo(n int) (Info, error) {
	// TODO: Implement
	return Info{
		Results: n,
	}, nil
}
