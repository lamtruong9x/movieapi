package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Define an error that our UnmarshalJSON() method can return if we're unable to parse
// or convert the JSON string successfully.
var ErrInvalidRuntimeFormat error = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	//Generate json value
	js := fmt.Sprintf("%d mins", r)

	//Put it in double quotes to satisfy JSON format
	js = strconv.Quote(js)
	return []byte(js), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	// Split the unquotedJSONValue to parts
	parts := strings.Split(unquotedJSONValue, " ")

	// Check if the json value follow Runtime type format
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	*r = Runtime(n)
	return nil
}
