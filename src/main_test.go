/**
 * Created with IntelliJ IDEA.
 * User: alek
 * Date: 6/26/13
 * Time: 9:09 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEquals(t *testing.T, expect, actual interface{}) {
	if expect != actual {
		message := fmt.Sprintf("Expected %v, found %v", expect, actual)
		if fmt.Sprintf("%v", expect) == fmt.Sprintf("%v", actual) {
			message = fmt.Sprintf("Expected %v [%v], found %v [%v]", expect,
				reflect.TypeOf(expect), actual, reflect.TypeOf(actual))
		}
		t.Fatal(message)
	}
}

func TestSumMultiples(t *testing.T) {
	assertEquals(t, 23, sumMultiples(10))
	assertEquals(t, 233168, sumMultiples(1000))
}

func TestSumEqualFibs(t *testing.T) {
	assertEquals(t, 10, sumEvenFibs(10))
	assertEquals(t, 44, sumEvenFibs(35))
	assertEquals(t, 4613732, sumEvenFibs(4000000))
}

func TestPrimeFactors(t *testing.T) {
	assertEquals(t, 29, largestPrimeFactor(13195))
	assertEquals(t, 6857, largestPrimeFactor(600851475143))
}

func TestLargestPalindrome(t *testing.T) {
	assertEquals(t, true, isPalindrome("racecar"))
	assertEquals(t, true, isPalindrome("RaceCar"))
	assertEquals(t, false, isPalindrome("Calvin"))

	assertEquals(t, "9009", largestPalindrome(2))
	assertEquals(t, "906609", largestPalindrome(3))
}

func TestSmallestDivisible(t *testing.T) {
	assertEquals(t, 2520, smallestDivisible(10))
	assertEquals(t, 27720, smallestDivisible(11))
	assertEquals(t, 27720, smallestDivisible(12))
	assertEquals(t, 360360, smallestDivisible(13))
	assertEquals(t, 232792560, smallestDivisible(20))
}

func TestSumSquareDifference(t *testing.T) {
	assertEquals(t, 4, sumSquareDifference(2))
	assertEquals(t, 22, sumSquareDifference(3))
	assertEquals(t, 2640, sumSquareDifference(10))
	assertEquals(t, 25164150, sumSquareDifference(100))
}

func Test10001thPrime(t *testing.T) {
	primes := primesTo(200000)

	if len(primes) < 10001 {
		t.Fatalf("Not enough primes, only found %v", len(primes))
	}

	assertEquals(t, 104743, primes[10000])
}

func TestMaxConsecutiveProduct(t *testing.T) {

	assertEquals(t, 120, maxConsecutiveProduct("12345"))
	assertEquals(t, 240, maxConsecutiveProduct("1212111223451121212"))

	giant := "73167176531330624919225119674426574742355349194934969835203" +
		"12774506326239578318016984801869478851843858615607891129494" +
		"95459501737958331952853208805511125406987471585238630507156" +
		"93290963295227443043557668966489504452445231617318564030987" +
		"11121722383113622298934233803081353362766142828064444866452" +
		"38749303589072962904915604407723907138105158593079608667017" +
		"24271218839987979087922749219016997208880937766572733300105" +
		"33678812202354218097512545405947522435258490771167055601360" +
		"48395864467063244157221553975369781797784617406495514929086" +
		"25693219784686224828397224137565705605749026140797296865241" +
		"45351004748216637048440319989000889524345065854122758866688" +
		"11642717147992444292823086346567481391912316282458617866458" +
		"35912456652947654568284891288314260769004224219022671055626" +
		"32111110937054421750694165896040807198403850962455444362981" +
		"23098787992724428490918884580156166097919133875499200524063" +
		"68991256071760605886116467109405077541002256983155200055935" +
		"72972571636269561882670428252483600823257530420752963450"
	assertEquals(t, 40824, maxConsecutiveProduct(giant))
}

func TestPythagoreanTriplet(t *testing.T) {
	assertEquals(t, 31875000, pythagoreanTripletSumProduct())
}

func TestPrimeSums(t *testing.T) {
	assertEquals(t, 17, sumPrimes(10))
	assertEquals(t, 142913828922, sumPrimes(2000000))
}


func TestGridProduct(t *testing.T) {
	smallGrid := [][]int {
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	assertEquals(t, 13*14*15*16, gridProduct(4, smallGrid))

	grid := [][]int {
		{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
		{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0},
		{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65},
		{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91},
		{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
		{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
		{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
		{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
		{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
		{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95},
		{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92},
		{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57},
		{86, 56, 0, 48, 35, 71, 89, 7, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
		{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40},
		{4, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
		{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
		{4, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
		{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16},
		{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54},
		{1, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48},
	}

	assertEquals(t, "", grid)
}


