package topsecret

import "strings"

type RequestTopSecret struct {
	Satellites []Satellite `json:"satellites"`
}

type RequestTopSecretSplit struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

// Records struct to represent a satellites records in the database
type Records struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `json:"name"`
	Distance float32 `json:"distance"`
	Message  string  `json:"message"`
}

// ToSatellite converts a Records database register struct to a Satellite struct
func (r *Records) ToSatellite() *Satellite {
	return &Satellite{
		Name:     r.Name,
		Distance: r.Distance,
		Message:  strings.Split(r.Message, "|"),
	}
}
