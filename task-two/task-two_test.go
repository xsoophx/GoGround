package task_two

import (
	"math"
	"reflect"
	"testing"
)

// 1
func TestCheckOverUnderZero(t *testing.T) {
	if result := checkOverUnderZero(1); result != "positive" {
		t.Errorf("expected 'positive', got %v", result)
	}
	if result := checkOverUnderZero(-1); result != "negative" {
		t.Errorf("expected 'negative', got %v", result)
	}
	if result := checkOverUnderZero(0); result != "zero" {
		t.Errorf("expected 'zero', got %v", result)
	}
}

// 2
func TestCircle(t *testing.T) {
	circumference, area := circle(5.0)
	expectedCircumference := 31.41592653589793
	expectedArea := 78.53981633974483

	if math.Abs(circumference-expectedCircumference) > 1e-9 {
		t.Errorf("expected circumference %v, got %v", expectedCircumference, circumference)
	}
	if math.Abs(area-expectedArea) > 1e-9 {
		t.Errorf("expected area %v, got %v", expectedArea, area)
	}
}

// 3
func TestSecondsToHoursMinutesSeconds(t *testing.T) {
	hours, minutes, seconds := secondsToHoursMinutesSeconds(3661)
	if hours != 1 || minutes != 1 || seconds != 1 {
		t.Errorf("expected (1, 1, 1), got (%v, %v, %v)", hours, minutes, seconds)
	}
}

// 4
func TestIsLeapYear(t *testing.T) {
	if !isLeapYear(2020) {
		t.Errorf("expected true, got false")
	}
	if isLeapYear(2021) {
		t.Errorf("expected false, got true")
	}
}

// 5
func TestToLowerCase(t *testing.T) {
	if result := toLowerCase("HALLO"); result != "hallo" {
		t.Errorf("expected 'hallo', got %v", result)
	}
}

// 6
func TestDigitSum(t *testing.T) {
	if result := digitSum(123); result != 6 {
		t.Errorf("expected 6, got %v", result)
	}
}

// 7
func TestDivisibleByThreeOrFive(t *testing.T) {
	expected := []int{3, 5, 6, 9, 10}
	result := divisibleByThreeOrFive(10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// 8
func TestPrimes(t *testing.T) {
	expected := []int{2, 3, 5, 7}
	result := primes(10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// 9
func TestIsPrime(t *testing.T) {
	if !isPrime(2) {
		t.Errorf("expected true, got false")
	}
	if isPrime(4) {
		t.Errorf("expected false, got true")
	}
}

// 10
func TestDigitSumRecursive(t *testing.T) {
	if result := digitSumRecursive(123); result != 6 {
		t.Errorf("expected 6, got %v", result)
	}
}

// 11
func TestWordCount(t *testing.T) {
	if result := wordCount("Das ist ein tolles Kotlin Programm."); result != 6 {
		t.Errorf("expected 6, got %v", result)
	}
}

// 12
func TestCountUpperAndLowerCase(t *testing.T) {
	upper, lower := countUpperAndLowerCase("Das ist ein tolles Kotlin Programm.")
	if upper != 3 || lower != 26 {
		t.Errorf("expected (3, 26), got (%v, %v)", upper, lower)
	}
}

// 13
func TestCountCharactersWithoutSpaces(t *testing.T) {
	if result := countCharactersWithoutSpaces("Das ist ein tolles Kotlin Programm."); result != 30 {
		t.Errorf("expected 30, got %v", result)
	}
}

// 14
func TestCountDigits(t *testing.T) {
	if result := countDigits("4238 Bilder können aus 3 Enten R3ST4P1's entnommen werden."); result != 8 {
		t.Errorf("expected 8, got %v", result)
	}
}

// 15
func TestCountSpecialCharacters(t *testing.T) {
	if result := countSpecialCharacters("Das ist <Zahl> * <Sprache> Programm."); result != 11 {
		t.Errorf("expected 11, got %v", result)
	}
}

// 16
func TestSumOfPrevious(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	result := sumOfPrevious(11)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// 17
func TestReverseDigits(t *testing.T) {
	if result := reverseDigits(123); result != 321 {
		t.Errorf("expected 321, got %v", result)
	}
}

// 18
func TestReverseString(t *testing.T) {
	if result := reverseString("Das ist ein tolles Go Programm."); result != ".mmargorP oG sellot nie tsi saD" {
		t.Errorf("expected '.mmargorP niltoK selot nie tsi saD', got %v", result)
	}
}

// 19
func TestIsUnique(t *testing.T) {
	if result := isUnique("Das ist ein tolles Kotlin Programm."); result != false {
		t.Errorf("expected false, got %v", result)
	}
	if result := isUnique("Kotlin, C# & R."); result != true {
		t.Errorf("expected true, got %v", result)
	}
}

// 20
func TestSquares(t *testing.T) {
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	result := squares(10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// 21
func TestSumOfEvenIndices(t *testing.T) {
	if result := sumOfEvenIndices([]int{1, 2, 3, 4, 5, 6}); result != 9 {
		t.Errorf("expected 9, got %v", result)
	}
}

// 22
func TestMinAndMax(t *testing.T) {
	minimum, maximum := minAndMax([]int{1, 2, 3, 4, 5, 6})
	if minimum != 1 || maximum != 6 {
		t.Errorf("expected (1, 6), got (%v, %v)", minimum, maximum)
	}
}

// 23
func TestGreaterThanAverage(t *testing.T) {
	expected := []int{4, 5, 6}
	result := greaterThanAverage([]int{1, 2, 3, 4, 5, 6})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// 24
func TestSum(t *testing.T) {
	if result := sum([]int{1, 2, 3, 4, 5, 6}); result != 21 {
		t.Errorf("expected 21, got %v", result)
	}
}

// 25
func TestCountGreaterThan(t *testing.T) {
	if result := countGreaterThan([]int{1, 2, 3, 4, 5, 6}, 3); result != 3 {
		t.Errorf("expected 3, got %v", result)
	}
}

// 26
func TestSumOfSquares(t *testing.T) {
	if result := sumOfSquares([]int{1, 2, 3, 4, 5, 6}); result != 91 {
		t.Errorf("expected 91, got %v", result)
	}
}

// 27
func TestSortOddAndEven(t *testing.T) {
	expectedOdd := []int{1, 3, 5, 7, 9}
	expectedEven := []int{2, 4, 6, 8, 10}
	odd, even := sortOddAndEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if !reflect.DeepEqual(odd, expectedOdd) || !reflect.DeepEqual(even, expectedEven) {
		t.Errorf("expected %v, %v, got %v, %v", expectedOdd, expectedEven, odd, even)
	}
}

// 28
func TestGGT(t *testing.T) {
	if result := ggt(12, 18); result != 6 {
		t.Errorf("expected 6, got %v", result)
	}
}

// 29
func TestKgv(t *testing.T) {
	if result := kgv(12, 18); result != 36 {
		t.Errorf("expected 36, got %v", result)
	}
}

// 30
func TestIsPerfectNumber(t *testing.T) {
	if result := isPerfectNumber(6); result != true {
		t.Errorf("expected true, got %v", result)
	}
	if result := isPerfectNumber(7); result != false {
		t.Errorf("expected false, got %v", result)
	}
}

// 31
func TestPerfectNumbers(t *testing.T) {
	expected := []int{6, 28, 496, 8128}
	result := perfectNumbers(4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// 32
func TestAverageOfOddNumbers(t *testing.T) {
	if result := averageOfOddNumbers([]int{1, 2, 3, 4, 5, 6}); result != 3 {
		t.Errorf("expected 3, got %v", result)
	}
}

// 33
func TestBinaryToDecimal(t *testing.T) {
	result, err := binaryToDecimal("1101")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != 13 {
		t.Errorf("expected 13, got %v", result)
	}
}

// 34
func TestDecimalToBinary(t *testing.T) {
	if result := decimalToBinary(13); result != "1101" {
		t.Errorf("expected '1101', got %v", result)
	}
}
