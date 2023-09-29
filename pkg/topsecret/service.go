package topsecret

import "fmt"

// Service interface to represent a service layer for quasar-fire logic and make a dependency injection
type Service interface {
	TopSecret(info []Satellite) (Ship, error)
	TopSecretSplit(satellite *Satellite) error
	TopSecretSplitGet() (*Ship, error)
	TopSecretSplitInfo() ([]Satellite, error)
	TopSecretSplitDelete() error
}

// service struct to represent a service layer for quasar-fire logic
type service struct {
	repository Repository
}

// NewService creates a new service layer for quasar-fire logic
func NewService(repository Repository) service {
	return service{repository}
}

// TopSecret calculates the location of a ship and the message from a list of satellites
func (s service) TopSecret(info []Satellite) (Ship, error) {
	shipLocationX, shipLocationY := GetLocation(info[0].Distance, info[1].Distance, info[2].Distance)
	shipMessage := GetMessage(info[0].Message, info[1].Message, info[2].Message)

	ship := Ship{
		Position: Position{
			X: shipLocationX,
			Y: shipLocationY,
		},
		Message: shipMessage,
	}
	return ship, nil
}

// TopSecretSplit saves the information of a satellite
func (s service) TopSecretSplit(satellite *Satellite) error {
	return s.repository.Save(satellite)
}

// TopSecretSplitGet returns the location of a ship and the message
func (s service) TopSecretSplitGet() (*Ship, error) {
	satellites, err := s.repository.Find()
	if err != nil {
		return nil, err
	}

	if len(satellites) != 3 {
		return nil, fmt.Errorf("no hay suficiente información para calcular la posición o el mensaje")
	}

	shipLocationX, shipLocationY := GetLocation(satellites[0].Distance, satellites[1].Distance, satellites[2].Distance)
	shipMessage := GetMessage(satellites[0].Message, satellites[1].Message, satellites[2].Message)

	ship := Ship{
		Position: Position{
			X: shipLocationX,
			Y: shipLocationY,
		},
		Message: shipMessage,
	}
	return &ship, nil
}

// TopSecretSplitInfo returns the information like list of satellites
func (s service) TopSecretSplitInfo() ([]Satellite, error) {
	return s.repository.GetInfo()
}

func (s service) TopSecretSplitDelete() error {
	return s.repository.Delete()
}
