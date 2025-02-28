package validate

import (
	"errors"
	"strconv"
	"strings"

	"github.com/eirka/eirka-libs/config"
)

// Validate will check string length
type Validate struct {
	Input string
	Max   int
	Min   int
}

// ValidateParam will parse parameters from requests to see if they are uint or too huge
func ValidateParam(param string) (id uint, err error) {

	// make sure its a uint
	pid, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return
	}

	id = uint(pid)

	// check maximum param size
	if id > config.Settings.Limits.ParamMaxSize {
		err = errors.New("parameter too large")
		return
	}

	return
}

// MaxLength checks string for length
func (v *Validate) MaxLength() bool {
	return len(v.Input) > v.Max && len(v.Input) != 0
}

// MinLength checks string for length
func (v *Validate) MinLength() bool {
	return len(v.Input) < v.Min || len(v.Input) == 0
}

// MinPartsLength will check if the entire string is less than the defined minimum
func (v *Validate) MinPartsLength() bool {

	// check if the entire string is less than min
	if len(v.Input) < v.Min || v.Input == "" {
		return true
	}

	// break up into fields
	parts := strings.Fields(v.Input)

	// Check if parts is empty
	if len(parts) == 0 {
		return true
	}

	// check if the first part is min chars for search
	if len(parts[0]) < v.Min {
		return true
	}

	// use the fields slice as the input to get rid of any weird whitespace
	v.Input = strings.Join(parts, " ")

	return false
}

// IsEmpty checks to see if string is empty
func (v *Validate) IsEmpty() bool {
	return len(strings.TrimSpace(v.Input)) == 0
}

// Clamp a value to a max and min
func Clamp(value, max, min uint) uint {

	if value > max {
		return max
	} else if value < min {
		return min
	}

	return value
}
