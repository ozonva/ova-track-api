package utils

import "fmt"

func DivideIntoBatches(slice []uint8, batchSize int) [][]uint8 {

	if batchSize == 0 {
		panic("Batch size can't be 0")
	}
	curSlice := make([]uint8, 0)
	result := make([][]uint8, 0)
	for i, v := range slice {
		curSlice = append(curSlice, v)
		if (i+1)%batchSize == 0 {
			tmp := make([]uint8, batchSize)
			copy(tmp, curSlice)
			result = append(result, tmp)
			curSlice = curSlice[:0]
		}
	}
	if len(curSlice) != 0 {
		result = append(result, curSlice)
	}
	return result
}

func InverseMap(data map[string]int) map[int]string {

	result := make(map[int]string)
	for key, value := range data {
		_, contains := result[value]
		if contains {
			errDescription := ""
			errDescription = fmt.Sprintf(errDescription, "Key %d contains twice", value)
			panic(errDescription)
		}
		result[value] = key
	}
	return result
}

func FilterMap(data map[string]int, exclude []string) map[string]int {
	for _, val := range exclude {
		_, contains := data[val]
		if contains {
			delete(data, val)
		}
	}
	return data
}
