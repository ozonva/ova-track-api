package utils

import "testing"
import "reflect"

type divideIntoBatchesTestData struct {
	slice     []uint8
	batchSize int
	expected  [][]uint8
}

func TestDivideIntoBatches(t *testing.T) {

	testData := [...]divideIntoBatchesTestData{
		{[]uint8{1, 1, 2, 2, 3, 3, 4}, 2, [][]uint8{{1, 1}, {2, 2}, {3, 3}, {4}}},
		{[]uint8{1, 1, 2, 2, 3, 3, 4, 4}, 2, [][]uint8{{1, 1}, {2, 2}, {3, 3}, {4, 4}}},
		{[]uint8{1, 2, 3, 4}, 1, [][]uint8{{1}, {2}, {3}, {4}}},
		// {[]uint8{1, 2, 3, 4}, 0, [][]uint8{}},
	}

	for i := range testData {
		result := DivideIntoBatches(testData[i].slice, testData[i].batchSize)
		if !reflect.DeepEqual(result, testData[i].expected) {
			t.Errorf("DevideIntoBatchesData result %v expetcted %v", result, testData[i].expected)
		}
	}
}

func TestInverseMap(t *testing.T) {

	{
		result := InverseMap(map[string]int{"one": 1, "two": 2, "three": 3})
		expected := map[int]string{1: "one", 2: "two", 3: "three"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test inverse map result %v expected %v", result, expected)
		}
	}
	{
		defer func() { recover() }()
		InverseMap(map[string]int{"one": 1, "two": 1, "three": 3})
		t.Errorf("Double values in map")
	}
}

func TestFilterMap(t *testing.T) {
	{
		testData := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
		toRemove := []string{"two", "three"}
		expected := map[string]int{"one": 1, "four": 4}
		result := FilterMap(testData, toRemove)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test filter map result %v expected %v", result, expected)
		}
	}
	{
		testData := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
		toRemove := []string{}
		expected := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
		result := FilterMap(testData, toRemove)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test filter map result %v expected %v", result, expected)
		}
	}
	{
		testData := map[string]int{}
		toRemove := []string{"two", "three"}
		expected := map[string]int{}
		result := FilterMap(testData, toRemove)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Test filter map result %v expected %v", result, expected)
		}
	}
}
