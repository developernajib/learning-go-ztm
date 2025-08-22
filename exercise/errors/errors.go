//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
	Second int
}

func ParseTime(input string) (Time, error) {
	parts := strings.Split(input, ":")

	if len(parts) != 3 {
		return Time{}, errors.New("invalid time format, expected HH:MM:SS")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, errors.New("invalid hour value")
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return Time{}, errors.New("invalid minute value")
	}

	second, err := strconv.Atoi(parts[2])
	if err != nil {
		return Time{}, errors.New("invalid second value")
	}

	return Time{
		Hour:   hour,
		Minute: minute,
		Second: second,
	}, nil
}
