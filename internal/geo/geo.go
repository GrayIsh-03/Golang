package geo

import (
	"errors"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

// get-method return latitude for type Coordinates
func (c *Coordinates) Latitude() float64 {
	return c.latitude
}

// get-method return longitude for type Coordinates
func (c *Coordinates) Longitude() float64 {
	return c.longitude
}

// set-method for check correct value to latitude
func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid value is latitude")
	}
	c.latitude = latitude
	return nil
}

// set-method for check correct value to longitude
func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid value is longitude")
	}
	c.longitude = longitude
	return nil
}
