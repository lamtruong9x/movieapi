package validator

import "regexp"

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]string
}

// Return a new Validator
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Validate if the validator 's map doesn't contain any entry
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message to the map (so long as no entry already exists for
// the given key).
func (v *Validator) AddError(key, message string) {
	if _, exist := v.Errors[key]; !exist {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// Generic function which returns true if a specific value is in a list.
func PermittedValues[T comparable](value T, permmitedValues ...T) bool {
	for i := range permmitedValues {
		if value == permmitedValues[i] {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}

// Generic function which returns true if all values in a slice are unique.
// func Unique[T comparable](values []T) bool {
// 	uniqueValues := make(map[T]bool)
// 	for i := range values {
// 		uniqueValues[values[i]] = true
// 	}
// 	return len(uniqueValues) == len(values)
// }

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}
