package topsecret

// Service interface to represent a service layer for quasar-fire logic and make a dependency injection
type Service interface {
	TopSecret(info []Satellite) (Ship, error)
	TopSecretSplitInfo() ([]Records, error)
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

// TopSecretSplitInfo returns the information like list of satellites
func (s service) TopSecretSplitInfo() ([]Records, error) {
	return s.repository.GetInfo()
}
