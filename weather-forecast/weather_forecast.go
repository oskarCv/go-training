// Package weather provides tools to convert.
package weather

// CurrentCondition represents a condition.
var CurrentCondition string

// CurrentLocation represents a location.
var CurrentLocation string

// Forecast returns an string value with the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
