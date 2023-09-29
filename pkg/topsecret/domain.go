package topsecret

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"strings"
)

// Satellite struct to represent a satellite in the space
type Satellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

// Ship struct to represent a ship in the space
type Ship struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

// Position struct to represent a position in the space
type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

// GetLocation calculates the location of a ship returning
// the x and y coordinates of location making triangulation
// of distances from 3 known point
func GetLocation(distances ...float32) (float32, float32) {

	// -- Satellites locations -- //

	// Kenobi
	x1, y1 := -500.0, -200.0
	// Skywalker
	x2, y2 := 100.0, -100.0
	// Sato
	x3, y3 := 500.0, 100.0

	// -- Distances from satellites -- //

	// Distances from satellites
	d1 := float64(distances[0])
	d2 := float64(distances[1])
	d3 := float64(distances[2])

	// -- Calculate the location -- //
	// To calculate the location of ship we need to solve the system of linear equations

	// Making matrix of coefficients A
	A := mat.NewDense(3, 2, []float64{2 * (x1 - x3), 2 * (y1 - y3), 2 * (x2 - x3), 2 * (y2 - y3), 2 * (x1 - x2), 2 * (y1 - y2)})
	// Making vector of constants b
	b := mat.NewVecDense(3, []float64{d1*d1 - d3*d3 - x1*x1 + x3*x3 - y1*y1 + y3*y3, d2*d2 - d3*d3 - x2*x2 + x3*x3 - y2*y2 + y3*y3, d1*d1 - d2*d2 - x1*x1 + x2*x2 - y1*y1 + y2*y2})

	// Solve the system of linear equations
	var xVec mat.VecDense
	if err := xVec.SolveVec(A, b); err != nil {
		fmt.Println("Error solving the system of linear equations")
		return 0, 0
	}

	// Return the coordinates
	return float32(xVec.AtVec(0)), float32(xVec.AtVec(1))
}

// GetMessage returns the message from the ship comparing the messages
// from the satellites and returning the message with the missing words
func GetMessage(messages ...[]string) string {
	var reverseMessage []string

	for index, message := range messages {
		counter := 0

		// Iterate over the slice from back to front to avoid the gaps in the message
		for i := len(message) - 1; i >= 0; i-- {
			if index == 0 {
				reverseMessage = append(reverseMessage, message[i])
				continue
			}

			// If the message is not empty and the message is different from the message in reverseMessage
			if message[i] != "" && message[i] != reverseMessage[counter] {
				reverseMessage[counter] = message[i]
			}

			counter++
		}
	}

	var decodedMessage string
	// Iterate over the slice from back to front to organize the message
	for i := range reverseMessage {
		index := len(reverseMessage) - 1 - i
		decodedMessage += reverseMessage[index] + " "
	}

	// Remove the spaces from the message
	return strings.Trim(decodedMessage, " ")
}
