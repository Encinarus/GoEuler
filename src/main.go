/**
 * Created with IntelliJ IDEA.
 * User: alek
 * Date: 6/21/13
 * Time: 8:59 AM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func sumMultiples(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func sumEvenFibs(maxFib int) int {
	first := 1
	second := 2
	next := first + second
	sum := second

	if first > maxFib {
		return 0
	} else if second > maxFib {
		return 1
	}

	for next < maxFib {
		if next%2 == 0 {
			sum += next
		}
		first = second
		second = next
		next = first + second
	}

	return sum
}

func primesTo(maxPrime int) []int {
	// Find primes up to value
	// potentials := make(map[int]bool)
	potentials := make([]bool, maxPrime+1)

	for i := 2; i <= maxPrime; i++ {
		potentials[i] = true
	}

	for i := 2; i <= maxPrime; i++ {
		// If it's not prime, skip
		if !potentials[i] {
			continue
		}

		for j := i * i; j <= maxPrime; j += i {
			potentials[j] = false
		}
	}

	primes := make([]int, maxPrime)
	primeIndex := 0
	for i := 2; i <= maxPrime; i++ {
		if potentials[i] {
			primes[primeIndex] = i
			primeIndex++
		}
	}
	return primes[0:primeIndex]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func factor(primes []int, value int) int {
	//original := value
	largest := 1

	for _, prime := range primes {
		if prime == value {
			return max(largest, value)
		}

		for value%prime == 0 {
			largest = max(largest, prime)
			value /= prime
		}
	}

	return largest
}

func largestPrimeFactor(value int) int {
	maxPrime := int(math.Sqrt(float64(value)) + 1)
	primes := primesTo(maxPrime)
	return factor(primes, value)
}

func isPalindrome(str string) bool {
	lowered := strings.ToLower(str)
	length := len(lowered)
	for i := 0; i <= length/2; i++ {
		if lowered[i] != lowered[length-i-1] {
			return false
		}
	}

	return true
}

func largestPalindrome(digits int) string {
	min := int(math.Pow(10, float64(digits-1)))
	max := int(math.Pow(10, float64(digits)))

	largest := 0
	for i := min; i < max; i++ {
		for j := i; j < max; j++ {
			product := i * j
			if product > largest && isPalindrome(fmt.Sprintf("%d", product)) {
				largest = i * j
			}
		}

	}

	return fmt.Sprintf("%d", largest)
}

func smallestDivisible(maxDivisor int) int {
	primeFactors := make(map[int]int)

	primes := primesTo(maxDivisor)
	for i := 1; i <= maxDivisor; i++ {
		value := i
		for _, prime := range primes {
			count := 0
			for value%prime == 0 {
				count++
				value /= prime
			}
			primeFactors[prime] = max(count, primeFactors[prime])
		}
	}

	product := 1
	for prime, count := range primeFactors {
		for i := 0; i < count; i++ {
			product *= prime
		}
	}

	return product
}

func sumSquareDifference(naturalNumberCount int) int {
	sumOfSquares := 0
	sumOfNumbers := 0

	for i := 1; i <= naturalNumberCount; i++ {
		sumOfNumbers += i
		sumOfSquares += (i * i)
	}

	return (sumOfNumbers * sumOfNumbers) - sumOfSquares
}

func maxConsecutiveProduct(numberString string) int {
	maxProduct := 1
	digits := list.New()

	for _, r := range strings.Split(numberString, "") {
		currentNum, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println("Err converting", r, "to a number, err:", err)
		}

		if digits.Len() == 5 {
			digits.Remove(digits.Front())
		}

		digits.PushBack(currentNum)

		if digits.Len() == 5 {
			product := 1
			for iter := digits.Front(); iter != nil; iter = iter.Next() {
				product *= iter.Value.(int)
			}
			maxProduct = max(product, maxProduct)
		}
	}

	return maxProduct
}

func pythagoreanTripletSumProduct() int {
	for a := 1; a < 1000; a++ {
		aSqr := a * a
		for b := a; b < 1000; b++ {
			bSqr := b * b

			c := 1000 - (b + a)
			cSqr := c * c

			if (aSqr + bSqr) == cSqr {
				return a * b * c
			}
		}
	}

	return 123
}

func sumPrimes(threshold int) int {
	primes := primesTo(threshold)
	sum := 0
	for _, prime := range primes {
		sum += prime
	}

	return sum
}

func maxCellValue(runLength, row, col int, grid [][]int) int {
  maxProduct := 0

  products := [8]int { 1, 1, 1, 1, 1, 1, 1, 1 }
  for i := 0; i < runLength; i++ {
    products[0] *= grid[row - i][col]      // Up
    products[1] *= grid[row + i][col]      // Down
    products[2] *= grid[row][col - i]      // Left
    products[3] *= grid[row][col + i]      // Right
    products[4] *= grid[row - i][col - i]  // Up Left
    products[5] *= grid[row - i][col + i]  // Up Right
    products[6] *= grid[row + i][col - i]  // Down Left
    products[7] *= grid[row + i][col + i]  // Down Right
  }

  for _, product := range products {
    maxProduct = max(maxProduct, product)
  }
  return maxProduct
}

func gridProduct(runLength int, grid [][]int) int {
	// Assume square grid with at least one row
	height := len(grid)
  bufferedHeight := height + (2*runLength)
	width := len(grid[0])
  bufferedWidth := width + (2*runLength)

  // Allocate grid in one shot, more efficient for static sized grid
  // per effective go
  bufferedGrid := make([][]int, bufferedHeight)
  bufferedRows := make([]int, bufferedWidth * bufferedHeight)

	for row := 0; row < bufferedHeight; row++ {
    bufferedGrid[row], bufferedRows = bufferedRows[:bufferedWidth], bufferedRows[bufferedWidth:]
	}

  for row := 0; row < height; row++ {
    for col := 0; col < width; col++ {
      bufferedGrid[row + runLength][col + runLength] = grid[row][col]
    }
  }

  // Now that we're set up, find maxValue for every cell
  maxProduct := 0
  for row := runLength; row < height + runLength; row++ {
    for col := runLength; col < width + runLength; col++ {
      maxProduct = max(maxProduct, maxCellValue(runLength, row, col, bufferedGrid))
    }
  }

	return maxProduct
}

func countFactors(triangleNumber int) int {
  sqrt := int(math.Sqrt(float64(triangleNumber)))

  factorCount := 0
  for i := 1; i <= sqrt; i++ {
    if triangleNumber % i == 0 {
      factorCount++
      if i*i != triangleNumber {
        factorCount++
      }
    }
  }

  return factorCount
}

func smallestTriangleWithMinimumFactorCount(minFactorCount int) int {
  triangle := 0
  for i := 1; true; i++ {
    triangle += i

    if countFactors(triangle) > minFactorCount {
      break
    }
  }

  return triangle
}

