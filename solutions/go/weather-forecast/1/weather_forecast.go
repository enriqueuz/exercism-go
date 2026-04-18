// Package weather provides weather forecasting funtions.
package weather

var (
    // CurrentCondition represents the current weather condition of the country.
	CurrentCondition string
    // CurrentLocation represents the current location being queried.
	CurrentLocation  string
)

// Forecast receives city and condition arguments and returns a string with current location and current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
