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

import "math/big"

// PrimesLessThan will return the list
// of all primes less than the provided integer argument.
func PrimesLessThan(n int) []int {
	if n < 3 {
		return []int{}
	}

	var primes []int
	for i := 2; i < n; i++ {
		prime := true

		for _, p := range primes {
			if i%p == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, i)
		}
	}

	return primes
}

// FirstNPrimes wil return the first n-primes.
func FirstNPrimes(n int) []int {
	var primes []int
	i := 2
	for {
		prime := true

		for _, p := range primes {
			if i%p == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, i)
			if len(primes) == n {
				break
			}
		}
		i++
	}

	return primes
}

// TriangularSum returns the sum of all positive integers
// less than and equal to the supplied argument.
func TriangularSum(n int) int {
	var total int
	for i := 0; i <= n; i++ {
		total += i
	}

	return total
}

// Factorial will calculate the factorial
// function for the given float64 argument.
// Only integer increments will be used, but
// the final result is returned in big.Float format.
func Factorial(n float64) *big.Float {
	if n <= 0 {
		return big.NewFloat(float64(1))
	}

	total := big.NewFloat(1)

	for i := 1; float64(i) <= n; i++ {
		bigI := big.NewFloat(float64(i))
		total.Mul(total, bigI)
	}

	return total
}

// OddsProduct will return the product of all
// odd numbers less than `n` (the supplied float64 argument).
func OddsProduct(n float64) *big.Float {
	result := big.NewFloat(float64(1))
	for i := 1; float64(i) < n; i += 2 {
		bigI := big.NewFloat(float64(i))
		result.Mul(result, bigI)
	}

	return result
}

// Percentage will return a number between (or equal to) 0 and 1
func Percentage(numerator, denominator float64) (percentage float64) {
	ratio := numerator / denominator
	if ratio >= 1 {
		return float64(1)
	} else if ratio <= 0 {
		return float64(0)
	}

	percentage = ratio

	return
}
