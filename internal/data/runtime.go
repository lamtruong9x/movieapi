package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	//Generate json value
	js := fmt.Sprintf("%d mins", r)

	//Put it in double quotes to satisfy JSON format
	js = strconv.Quote(js)
	return []byte(js), nil
}
