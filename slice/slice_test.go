package slice

import (
	"reflect"
	"strings"
	"testing"
)

func TestContainsAll(t *testing.T) {
	// Testing when sliceA contains all elements of sliceB
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{2, 4}
	expected := true
	result := ContainsAll(sliceA, sliceB...)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Testing when sliceA does not contain all elements of sliceB
	sliceA = []int{1, 2, 3, 4, 5}
	sliceB = []int{6, 7}
	expected = false
	result = ContainsAll(sliceA, sliceB...)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Testing when sliceA is empty
	sliceA = []int{}
	sliceB = []int{1, 2, 3}
	expected = false
	result = ContainsAll(sliceA, sliceB...)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Testing when sliceB is empty
	sliceA = []int{1, 2, 3}
	sliceB = []int{}
	expected = true
	result = ContainsAll(sliceA, sliceB...)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Testing when both sliceA and sliceB are empty
	sliceA = []int{}
	sliceB = []int{}
	expected = true
	result = ContainsAll(sliceA, sliceB...)
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	func(t *testing.T) {
		sliceA := []string{
			"surname", "given_name", "type_id", "total_visit", "referral_code", "mobile", "updated_at",
		}
		sliceB := []string{}
		expected = true
		result = ContainsAll(sliceA, sliceB...)
		if result != expected {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	}(t)
}

func TestConcat(t *testing.T) {
	// Test case 1: Concatenating two integer arrays
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	expectedResult1 := []int{1, 2, 3, 4, 5, 6}
	result1 := Concat(arr1, arr2)
	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Errorf("Concatenation of two integer arrays failed. Expected: %v, but got: %v", expectedResult1, result1)
	}

	// Test case 2: Concatenating three string arrays
	arr3 := []string{"hello", "world"}
	arr4 := []string{"foo", "bar"}
	arr5 := []string{"baz"}
	expectedResult2 := []string{"hello", "world", "foo", "bar", "baz"}
	result2 := Concat(arr3, arr4, arr5)
	if !reflect.DeepEqual(result2, expectedResult2) {
		t.Errorf("Concatenation of three string arrays failed. Expected: %v, but got: %v", expectedResult2, result2)
	}

	// Test case 3: Concatenating empty arrays
	arr6 := []int{}
	arr7 := []int{}
	expectedResult3 := []int{}
	result3 := Concat(arr6, arr7)
	if !reflect.DeepEqual(result3, expectedResult3) {
		t.Errorf("Concatenation of empty arrays failed. Expected: %v, but got: %v", expectedResult3, result3)
	}
}

func TestConcatUnique(t *testing.T) {
	// Test case 1: Concatenating slices with no duplicates
	arrs1 := []int{1, 2, 3}
	arrs2 := []int{4, 5, 6}
	expected1 := []int{1, 2, 3, 4, 5, 6}
	result1 := ConcatUnique(arrs1, arrs2)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("ConcatUnique(%v, %v) = %v, want %v", arrs1, arrs2, result1, expected1)
	}

	// Test case 2: Concatenating slices with duplicates
	arrs3 := []int{1, 2, 3}
	arrs4 := []int{2, 3}
	arrs5 := []int{3, 4, 5}
	expected2 := []int{1, 2, 3, 4, 5}
	result2 := ConcatUnique(arrs3, arrs4, arrs5)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("ConcatUnique(%v, %v, %v) = %v, want %v", arrs3, arrs4, arrs5, result2, expected2)
	}

	// Test case 3: Concatenating empty slices
	arrs6 := []int{}
	arrs7 := []int{}
	expected3 := []int{}
	result3 := ConcatUnique(arrs6, arrs7)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("ConcatUnique(%v, %v) = %v, want %v", arrs6, arrs7, result3, expected3)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	// Test case 1: test with an empty slice
	arr1 := []int{}
	expected1 := []int{}
	result1 := RemoveDuplicates(arr1...)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Test case 2: test with a slice containing only unique elements
	arr2 := []int{1, 2, 3, 4, 5}
	expected2 := []int{1, 2, 3, 4, 5}
	result2 := RemoveDuplicates(arr2...)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}

	// Test case 3: test with a slice containing duplicate elements
	arr3 := []int{1, 2, 3, 2, 4, 3, 5}
	expected3 := []int{1, 2, 3, 4, 5}
	result3 := RemoveDuplicates(arr3...)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, result3)
	}

	// Test case 4: test with a slice containing duplicate elements of a different type
	arr4 := []string{"apple", "banana", "cherry", "banana", "date", "cherry", "elderberry"}
	expected4 := []string{"apple", "banana", "cherry", "date", "elderberry"}
	result4 := RemoveDuplicates(arr4...)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Expected %v, but got %v", expected4, result4)
	}
}

func TestUnshift(t *testing.T) {
	// Test case 1: Adding element to an empty slice
	{
		expected := []int{1}
		result := Unshift(1)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test case 1 failed: Expected %v, but got %v", expected, result)
		}
	}

	// Test case 2: Adding element to a non-empty slice
	{
		expected := []int{4, 1, 2, 3}
		result := Unshift(4, 1, 2, 3)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test case 2 failed: Expected %v, but got %v", expected, result)
		}
	}

	// Test case 3: Adding element to a slice of strings
	{
		expected := []string{"world", "hello"}
		result := Unshift("world", "hello")
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test case 3 failed: Expected %v, but got %v", expected, result)
		}
	}
}

func TestIUnshift(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		value int
		want  []int
	}{
		{
			name:  "Insert at the beginning of an empty slice",
			items: []int{},
			value: 5,
			want:  []int{5},
		},
		{
			name:  "Insert at the beginning of a non-empty slice",
			items: []int{1, 2, 3},
			value: 5,
			want:  []int{5, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IUnshift(tt.items, tt.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IUnshift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		items []interface{}
		want  []interface{}
	}{
		{
			name:  "Test case 1",
			items: []interface{}{1, 2, 3, 4, 5},
			want:  []interface{}{5, 4, 3, 2, 1},
		},
		{
			name:  "Test case 2",
			items: []interface{}{"a", "b", "c", "d", "e"},
			want:  []interface{}{"e", "d", "c", "b", "a"},
		},
		{
			name:  "Test case 3",
			items: []interface{}{true, false, true},
			want:  []interface{}{true, false, true},
		},
		// Add more test cases here...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.items...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	// Test case 1: Finding index of element that satisfies the predicate
	index1 := FindIndex(func(x int) bool { return x > 5 }, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	if index1 != 5 {
		t.Errorf("Expected index 5, got %d", index1)
	}

	// Test case 2: Finding index of element that satisfies the predicate in a slice of strings
	index2 := FindIndex(func(s string) bool { return len(s) > 4 }, "apple", "banana", "orange", "grape")
	if index2 != 0 {
		t.Errorf("Expected index 0, got %d", index2)
	}

	// Test case 3: No element satisfies the predicate
	index3 := FindIndex(func(x int) bool { return x < 0 }, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	if index3 != -1 {
		t.Errorf("Expected index -1, got %d", index3)
	}

	// Test case 4: No elements provided
	index4 := FindIndex(func(x int) bool { return x > 0 })
	if index4 != -1 {
		t.Errorf("Expected index -1, got %d", index4)
	}
}

func TestIFindIndex(t *testing.T) {
	// Test case 1: empty slice
	items1 := []int{}
	predicate1 := func(num int) bool {
		return num > 0
	}
	expected1 := -1
	if index := IFindIndex(items1, predicate1); index != expected1 {
		t.Errorf("Expected index %d, but got %d", expected1, index)
	}

	// Test case 2: slice with elements that satisfy the predicate
	items2 := []string{"apple", "banana", "cherry"}
	predicate2 := func(str string) bool {
		return len(str) > 5
	}
	expected2 := 1
	if index := IFindIndex(items2, predicate2); index != expected2 {
		t.Errorf("Expected index %d, but got %d", expected2, index)
	}

	// Test case 3: slice with elements that do not satisfy the predicate
	items3 := []int{1, 2, 3, 4, 5}
	predicate3 := func(num int) bool {
		return num > 10
	}
	expected3 := -1
	if index := IFindIndex(items3, predicate3); index != expected3 {
		t.Errorf("Expected index %d, but got %d", expected3, index)
	}
}

func TestFind(t *testing.T) {
	// Test case 1: Testing with an empty slice
	// Expecting that the function returns false as no item is found
	if _, found := Find(func(i int) bool { return i > 0 }, []int{}...); found {
		t.Errorf("Expected false, got true")
	}

	// Test case 2: Testing with a slice containing items that satisfy the predicate function
	// Expecting that the function returns the first item that satisfies the predicate function and true
	if _, found := Find(func(i int) bool { return i > 5 }, []int{1, 3, 7, 9}...); !found {
		t.Errorf("Expected true, got false")
	}

	// Test case 3: Testing with a slice containing items that do not satisfy the predicate function
	// Expecting that the function returns false as no item satisfies the predicate function
	if _, found := Find(func(i int) bool { return i < 0 }, []int{1, 3, 7, 9}...); found {
		t.Errorf("Expected false, got true")
	}

	// Test case 4: Testing with a slice of strings
	// Expecting that the function returns the first string that satisfies the predicate function and true
	if _, found := Find(func(s string) bool { return len(s) > 3 }, []string{"apple", "banana", "orange"}...); !found {
		t.Errorf("Expected true, got false")
	}

	// Test case 5: Testing with a slice of custom struct objects
	// Expecting that the function returns the first object that satisfies the predicate function and true
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	if _, found := Find(func(p Person) bool { return p.Age > 30 }, people...); !found {
		t.Errorf("Expected true, got false")
	}
}

func TestFindOrDefault(t *testing.T) {
	// Testing when no items are provided
	t.Run("No items provided", func(t *testing.T) {
		predicate := func(item int) bool {
			return item > 0
		}
		result := FindOrDefault(predicate)
		expected := 0

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	// Testing when no item satisfies the predicate
	t.Run("No item satisfies the predicate", func(t *testing.T) {
		predicate := func(item int) bool {
			return item > 5
		}
		result := FindOrDefault(predicate, 1, 2, 3, 4)
		expected := 0

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	// Testing when the first item satisfies the predicate
	t.Run("First item satisfies the predicate", func(t *testing.T) {
		predicate := func(item int) bool {
			return item > 5
		}
		result := FindOrDefault(predicate, 6, 7, 8, 9)
		expected := 6

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	// Testing when the last item satisfies the predicate
	t.Run("Last item satisfies the predicate", func(t *testing.T) {
		predicate := func(item int) bool {
			return item > 5
		}
		result := FindOrDefault(predicate, 1, 2, 3, 4, 6)
		expected := 6

		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}

func TestIFind(t *testing.T) {
	// Test case 1: Finding an element that satisfies the condition
	items1 := []int{1, 2, 3, 4, 5}
	predicate1 := func(n int) bool { return n%2 == 0 } // Even number
	expectedValue1 := 2
	expectedFound1 := true
	value1, found1 := IFind(items1, predicate1)
	if value1 != expectedValue1 || found1 != expectedFound1 {
		t.Errorf("Expected (%v, %v), but got (%v, %v)", expectedValue1, expectedFound1, value1, found1)
	}

	// Test case 2: Finding an element that satisfies the condition
	items2 := []string{"apple", "banana", "cherry"}
	predicate2 := func(s string) bool { return len(s) > 5 } // Length greater than 5
	expectedValue2 := "banana"
	expectedFound2 := true
	value2, found2 := IFind(items2, predicate2)
	if value2 != expectedValue2 || found2 != expectedFound2 {
		t.Errorf("Expected (%v, %v), but got (%v, %v)", expectedValue2, expectedFound2, value2, found2)
	}

	// Test case 3: Finding an element that does not satisfy the condition
	items3 := []int{1, 3, 5, 7, 9}
	predicate3 := func(n int) bool { return n%2 == 0 } // Even number
	expectedFound3 := false
	_, found3 := IFind(items3, predicate3)
	if found3 != expectedFound3 {
		t.Errorf("Expected %v, but got %v", expectedFound3, found3)
	}
}

func TestIFindOrDefault(t *testing.T) {
	// Test case 1: Finding an existing element in the slice.
	items := []int{1, 2, 3, 4, 5}
	predicate := func(x int) bool { return x%2 == 0 }
	expected := 2
	result := IFindOrDefault(items, predicate)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: Finding a non-existing element in the slice.
	items = []int{1, 3, 5, 7, 9}
	predicate = func(x int) bool { return x%2 == 0 }
	expected = 0 // Default value for int
	result = IFindOrDefault(items, predicate)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 3: Finding a string element in the slice.
	itemsStr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	predicateStr := func(s string) bool { return len(s) > 5 }
	expectedStr := "banana"
	resultStr := IFindOrDefault(itemsStr, predicateStr)
	if resultStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, resultStr)
	}
}

func TestFilter(t *testing.T) {
	// Test case 1: Filtering an empty slice should return an empty slice
	result := Filter(func(i int) bool { return i%2 == 0 })
	if len(result) != 0 {
		t.Error("Expected an empty slice, but got:", result)
	}

	// Test case 2: Filtering a slice of integers for even numbers
	result = Filter(func(i int) bool { return i%2 == 0 }, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	expectedResult := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Expected:", expectedResult, "but got:", result)
	}

	func(t *testing.T) {
		// Test case 3: Filtering a slice of strings for words starting with 'A'
		result := Filter(func(s string) bool { return strings.HasPrefix(s, "A") }, "Apple", "Banana", "Avocado")
		expectedResultString := []string{"Apple", "Avocado"}
		if !reflect.DeepEqual(result, expectedResultString) {
			t.Error("Expected:", expectedResultString, "but got:", result)
		}
	}(t)
}

func TestIFilter(t *testing.T) {
	// Test cases
	tests := []struct {
		name      string
		items     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "Empty slice",
			items:     []int{},
			predicate: func(n int) bool { return n > 0 },
			expected:  []int{},
		},
		{
			name:      "All items pass the predicate",
			items:     []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool { return n > 0 },
			expected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "No items pass the predicate",
			items:     []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool { return n > 10 },
			expected:  []int{},
		},
		{
			name:      "Some items pass the predicate",
			items:     []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool { return n%2 == 0 },
			expected:  []int{2, 4},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IFilter(test.items, test.predicate)
			if len(result) != len(test.expected) {
				t.Errorf("Expected %d items, but got %d", len(test.expected), len(result))
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("Expected %d, but got %d", test.expected[i], result[i])
				}
			}
		})
	}
}

func TestFilterAndSeparate(t *testing.T) {
	// Test case 1: Filtering even and odd numbers
	func(t *testing.T) {
		predicate := func(n int) bool {
			return n%2 == 0
		}

		items := []int{1, 2, 3, 4, 5}
		expectedResult := []int{2, 4}
		expectedSeparated := []int{1, 3, 5}

		result, separated := FilterAndSeparate(predicate, items...)
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected result to be %v, but got %v", expectedResult, result)
		}
		if !reflect.DeepEqual(separated, expectedSeparated) {
			t.Errorf("Expected separated to be %v, but got %v", expectedSeparated, separated)
		}
	}(t)

	// Test case 2: Filtering strings starting with "a"
	func(t *testing.T) {
		predicate := func(s string) bool {
			return s[0] == 'a'
		}

		items := []string{"apple", "banana", "avocado", "orange"}
		expectedResult := []string{"apple", "avocado"}
		expectedSeparated := []string{"banana", "orange"}

		result, separated := FilterAndSeparate(predicate, items...)
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected result to be %v, but got %v", expectedResult, result)
		}
		if !reflect.DeepEqual(separated, expectedSeparated) {
			t.Errorf("Expected separated to be %v, but got %v", expectedSeparated, separated)
		}

	}(t)

	// Test case 3: Filtering struct elements based on a custom predicate
	func(t *testing.T) {
		type Student struct {
			Name   string
			Age    int
			Active bool
		}
		predicate := func(s Student) bool {
			return s.Active
		}

		items := []Student{
			{Name: "John", Age: 20, Active: true},
			{Name: "Jane", Age: 22, Active: false},
			{Name: "Adam", Age: 25, Active: true},
		}
		expectedResult := []Student{
			{Name: "John", Age: 20, Active: true},
			{Name: "Adam", Age: 25, Active: true},
		}
		expectedSeparated := []Student{
			{Name: "Jane", Age: 22, Active: false},
		}

		result, separated := FilterAndSeparate(predicate, items...)
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected result to be %v, but got %v", expectedResult, result)
		}
		if !reflect.DeepEqual(separated, expectedSeparated) {
			t.Errorf("Expected separated to be %v, but got %v", expectedSeparated, separated)
		}
	}(t)
}

func TestEvery(t *testing.T) {
	// Test cases
	tests := []struct {
		name      string
		predicate func(int) bool
		items     []int
		want      bool
	}{
		{
			name:      "All items satisfy the predicate",
			predicate: func(n int) bool { return n > 0 },
			items:     []int{1, 2, 3},
			want:      true,
		},
		{
			name:      "Some items do not satisfy the predicate",
			predicate: func(n int) bool { return n > 0 },
			items:     []int{-1, 2, 3},
			want:      false,
		},
		{
			name:      "Empty items",
			predicate: func(n int) bool { return n > 0 },
			items:     []int{},
			want:      true,
		},
		{
			name:      "Empty predicate",
			predicate: nil,
			items:     []int{1, 2, 3},
			want:      true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Every(tt.predicate, tt.items...)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIEvery(t *testing.T) {
	// Test cases
	tests := []struct {
		name      string
		items     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			name:  "Empty slice",
			items: []int{},
			predicate: func(item int) bool {
				// Predicate always returns true
				return true
			},
			expected: true,
		},
		{
			name:  "All elements satisfy the predicate",
			items: []int{1, 2, 3, 4, 5},
			predicate: func(item int) bool {
				// Predicate checks if element is less than 6
				return item < 6
			},
			expected: true,
		},
		{
			name:  "Some elements do not satisfy the predicate",
			items: []int{1, 2, 3, 4, 5},
			predicate: func(item int) bool {
				// Predicate checks if element is less than 3
				return item < 3
			},
			expected: false,
		},
	}

	// Test each case
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IEvery(test.items, test.predicate)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	// Test case: reducing integers
	initialValue := 0
	callback := func(acc, curr int) int {
		return acc + curr
	}
	items := []int{1, 2, 3, 4, 5}
	expectedResult := 15
	if result := Reduce(initialValue, callback, items...); result != expectedResult {
		t.Errorf("Reduce() = %v, want %v", result, expectedResult)
	}

	// Test case: reducing strings
	initialValueStr := ""
	callbackStr := func(acc, curr string) string {
		return acc + curr
	}
	itemsStr := []string{"Hello", " ", "World", "!"}
	expectedResultStr := "Hello World!"
	if resultStr := Reduce(initialValueStr, callbackStr, itemsStr...); resultStr != expectedResultStr {
		t.Errorf("Reduce() = %v, want %v", resultStr, expectedResultStr)
	}

	// Test case: reducing floats
	initialValueFloat := 0.0
	callbackFloat := func(acc, curr float64) float64 {
		return acc + curr
	}
	itemsFloat := []float64{1.1, 2.2, 3.3}
	expectedResultFloat := 6.6
	if resultFloat := Reduce(initialValueFloat, callbackFloat, itemsFloat...); resultFloat != expectedResultFloat {
		t.Errorf("Reduce() = %v, want %v", resultFloat, expectedResultFloat)
	}
}

func TestDevide(t *testing.T) {
	// Test case 1: to is equal to the length of items
	to1 := 3
	items1 := []int{1, 2, 3, 4, 5, 6}
	expected1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	result1 := Devide(to1, items1...)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %+v, but got %+v", expected1, result1)
	}

	// Test case 2: to is greater than the length of items
	to2 := 4
	items2 := []string{"a", "b", "c"}
	expected2 := [][]string{{"a", "b", "c"}}
	result2 := Devide(to2, items2...)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %+v, but got %+v", expected2, result2)
	}

	// Test case 3: to is less than the length of items
	to3 := 2
	items3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expected3 := [][]float64{{1.1, 2.2}, {3.3, 4.4}, {5.5}}
	result3 := Devide(to3, items3...)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed: expected %+v, but got %+v", expected3, result3)
	}
}
