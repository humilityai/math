// Copyright 2020 Hummility AI Incorporated, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package math

import (
	"math"
	"regexp"
	"strconv"
	"unicode"
)

// IsNaN checks if a float64 value is NaN and returns 0 if it is
// else it returns the orginal float64 value.
func IsNaN(value float64) float64 {
	if math.IsNaN(value) {
		return 0
	}

	return value
}

// IsNaNWithDefault will return the provided default value
// if the float64 value is NaN.
func IsNaNWithDefault(value float64, def float64) float64 {
	if math.IsNaN(value) {
		return def
	}

	return value
}

// IsInf checks if a float64 value is Inf+ or Inf- and returns 0 if it is
// else it returns the original float64 value.
func IsInf(value float64) float64 {
	if math.IsInf(value, 1) {
		return math.MaxFloat64
	}

	if math.IsInf(value, -1) {
		return float64(math.MinInt64)
	}

	return value
}

// IsInfWithDefault will return the default value if the provided
// value is +Inf or -Inf.
func IsInfWithDefault(value float64, def float64) float64 {
	if math.IsInf(value, 1) || math.IsInf(value, -1) {
		return def
	}

	return value
}

// IsNum composes IsNaN and IsInf, returns 0 if not a float64 number
// else it returns original float64 value.
func IsNum(value float64) float64 {
	return IsNaN(IsInf(value))
}

// IsNumWithDefault is IsNum but with the default value returned
// if not a standard float64 value.
func IsNumWithDefault(value float64, def float64) float64 {
	return IsNaNWithDefault(IsInfWithDefault(value, def), def)
}

// IntegerLength returns the number of digits used
// to represent the integer in base-10.
func IntegerLength(i int64) (count int) {
	for i != 0 {
		i /= 10
		count++
	}

	return
}

// ExtractNumbersFromString ...
func ExtractNumbersFromString(s string) []float64 {
	re := regexp.MustCompile(`\d[\d]*\.?\d*`)
	submatchall := re.FindAllString(s, -1)
	var finalMatch []float64
	for _, element := range submatchall {
		runeslice := []rune(element)
		number := true
		for _, r := range runeslice {
			if !unicode.IsDigit(r) && r != '.' {
				number = false
			}
		}
		if !number {
			continue
		}
		num, err := strconv.ParseFloat(element, 64)
		if err != nil {
			continue
		}
		finalMatch = append(finalMatch, num)
	}

	return finalMatch
}
