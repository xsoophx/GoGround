package main

import (
	"reflect"
	"testing"
)

// Beispiel / Example
func TestPrintHallo(t *testing.T) {
	expected := "Hallo Welt!"
	if result := printHallo(); result != expected {
		t.Errorf("printHallo() = %v; want %v", result, expected)
	}
}

// 1.
func TestDouble(t *testing.T) {
	if result := double(5); result != 10 {
		t.Errorf("double(5) = %v; want 10", result)
	}
}

// 2.
func TestSumOf(t *testing.T) {
	if result := sumOf(5, 3); result != 8 {
		t.Errorf("sumOf(5, 3) = %v; want 8", result)
	}
}

// 3.
func TestLengthOfString(t *testing.T) {
	if result := lengthOfString("Hello"); result != 5 {
		t.Errorf("lengthOfString(\"Hello\") = %v; want 5", result)
	}
}

// 4.
func TestIsEven(t *testing.T) {
	if !isEven(4) {
		t.Errorf("isEven(4) = false; want true")
	}
	if isEven(3) {
		t.Errorf("isEven(3) = true; want false")
	}
}

// 5.
func TestPrintName(t *testing.T) {
	expected := "Hallo, Max!"
	if result := printName("Max"); result != expected {
		t.Errorf("printName(\"Max\") = %v; want %v", result, expected)
	}
}

// 6.
func TestAverage(t *testing.T) {
	list := []int{2, 4, 6}
	expected := 4.0
	if result := average(list); result != expected {
		t.Errorf("average(%v) = %v; want %v", list, result, expected)
	}
}

// 7.
func TestSortIntegers(t *testing.T) {
	list := []int{3, 1, 2}
	expected := []int{1, 2, 3}
	if result := sortIntegers(list); !reflect.DeepEqual(result, expected) {
		t.Errorf("sort(%v) = %v; want %v", list, result, expected)
	}
}

// 8.
func TestContains(t *testing.T) {
	list := []int{1, 2, 3}
	if !contains(list, 2) {
		t.Errorf("contains(%v, 2) = false; want true", list)
	}
	if contains(list, 4) {
		t.Errorf("contains(%v, 4) = true; want false", list)
	}
}

// 9.
func TestLongestWord(t *testing.T) {
	words := []string{"apple", "banana", "pear"}
	expected := "banana"
	if result := longestWord(words); result != expected {
		t.Errorf("longestWord(%v) = %v; want %v", words, result, expected)
	}
}

// 10.
func TestFillWithTenNumbers(t *testing.T) {
	expected := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if result := fillWithTenNumbers(); result != expected {
		t.Errorf("fillWithTenNumbers() = %v; want %v", result, expected)
	}
}

// 11.
func TestRandomNumbers(t *testing.T) {
	if result := fillWithRandomNumbers(10); len(result) != 10 {
		t.Errorf("randomNumbers() = %v; want 10", result)
	}
}

// 12.
func TestMinMax(t *testing.T) {
	list := []int{5, 1, 9, 3}
	minimum, maximum := minMax(list)
	if minimum != 1 || maximum != 9 {
		t.Errorf("minMax(%v) = (%v, %v); want (1, 9)", list, minimum, maximum)
	}
}

// 13.
func TestFactorial(t *testing.T) {
	if result := factorial(5); result != 120 {
		t.Errorf("factorial(5) = %v; want 120", result)
	}
}

// 14.
func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8}
	if result := fibonacci(7); !reflect.DeepEqual(result, expected) {
		t.Errorf("fibonacci(7) = %v; want %v", result, expected)
	}
}

// 15.
func TestSquares(t *testing.T) {
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	if result := squares(); !reflect.DeepEqual(result, expected) {
		t.Errorf("squares() = %v; want %v", result, expected)
	}
}

// 16.
func TestReverseString(t *testing.T) {
	if result := reverseString("hello"); result != "olleh" {
		t.Errorf("reverseString(\"hello\") = %v; want \"olleh\"", result)
	}
}

// 17.
func TestIsPalindrome(t *testing.T) {
	if !isPalindrome("madam") {
		t.Errorf("isPalindrome(\"madam\") = false; want true")
	}
	if isPalindrome("hello") {
		t.Errorf("isPalindrome(\"hello\") = true; want false")
	}
}

// 18.
func TestCountVowels(t *testing.T) {
	if result := countVowels("hello"); result != 2 {
		t.Errorf("countVowels(\"hello\") = %v; want 2", result)
	}
}

// 19.
func TestPascalTriangle(t *testing.T) {
	expected := [][]int{{1}, {1, 1}, {1, 2, 1}}
	if result := pascalTriangle(3); !reflect.DeepEqual(result, expected) {
		t.Errorf("pascalTriangle(3) = %v; want %v", result, expected)
	}
}

// 20.
func TestSortAscendingDescending(t *testing.T) {
	list := []int{3, 1, 2}
	ascending := []int{1, 2, 3}
	descending := []int{3, 2, 1}
	if result1, result2 := sortAscendingDescending(list); !reflect.DeepEqual(result1, ascending) || !reflect.DeepEqual(result2, descending) {
		t.Errorf("sortAscendingDescending(%v) = (%v, %v); want (%v, %v)", list, result1, result2, ascending, descending)
	}
}

// 21.
func TestIsPrime(t *testing.T) {
	if !isPrime(7) {
		t.Errorf("isPrime(7) = false; want true")
	}
	if isPrime(4) {
		t.Errorf("isPrime(4) = true; want false")
	}
}

// 22.
func TestSum(t *testing.T) {
	list := []int{1, 2, 3}
	if result := sum(list); result != 6 {
		t.Errorf("sum(%v) = %v; want 6", list, result)
	}
}

// 23.
func TestDuplicateList(t *testing.T) {
	list := []int{1, 2, 3}
	expected := []int{1, 2, 3, 1, 2, 3}
	if result := duplicateList(list); !reflect.DeepEqual(result, expected) {
		t.Errorf("duplicateList(%v) = %v; want %v", list, result, expected)
	}
}

// 24.
func TestAreIdentical(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{1, 2, 3}
	if !areIdentical(list1, list2) {
		t.Errorf("areIdentical(%v, %v) = false; want true", list1, list2)
	}
	list3 := []int{1, 2, 3}
	list4 := []int{1, 2, 4}
	if areIdentical(list3, list4) {
		t.Errorf("areIdentical(%v, %v) = true; want false", list3, list4)
	}
}

// 25.
func TestReversedList(t *testing.T) {
	list := []int{1, 2, 3}
	expected := []int{3, 2, 1}
	if result := reversedList(list); !reflect.DeepEqual(result, expected) {
		t.Errorf("reversedList(%v) = %v; want %v", list, result, expected)
	}
}

// 26.
func TestFindSecondSmallestAndSecondLargest(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	if smallest, largest := findSecondSmallestAndSecondLargest(array); smallest != 2 || largest != 4 {
		t.Errorf("findSecondSmallestAndSecondLargest(%v) = (%v, %v); want (2, 4)", array, smallest, largest)
	}
}

// 27.
func TestCopyArrays(t *testing.T) {
	array1 := []int{1, 2}
	array2 := []int{3, 4}
	expected := []int{1, 2, 3, 4}
	if result := copyArrays(array1, array2); !reflect.DeepEqual(result, expected) {
		t.Errorf("copyArrays(%v, %v) = %v; want %v", array1, array2, result, expected)
	}
}

// 28.
func TestFindCommonElements(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	expected := []int{2, 3}
	if result := findCommonElements(a, b); !reflect.DeepEqual(result, expected) {
		t.Errorf("findCommonElements(%v, %v) = %v; want %v", a, b, result, expected)
	}
}

// 29.
func TestIsOdd(t *testing.T) {
	if !isOdd(3) {
		t.Errorf("isOdd(3) = false; want true")
	}
	if isOdd(4) {
		t.Errorf("isOdd(4) = true; want false")
	}
}
