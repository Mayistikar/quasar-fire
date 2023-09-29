package topsecret

type Service interface {
	TopSecret(info []Satellite) (Ship, error)
}

type service struct{}

func NewService() service {
	return service{}
}

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
