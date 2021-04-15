package goutils

import (
	"encoding/json"
	"fmt"
	"math"
)

type KeepZero float64

func (f KeepZero) MarshalJSON() ([]byte, error) {
	if float64(f) == math.Trunc(float64(f)) {
		return []byte(fmt.Sprintf("%.2f", f)), nil
	}
	return json.Marshal(float64(f))
}
