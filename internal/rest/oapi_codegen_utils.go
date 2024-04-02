package rest

import (
	"fmt"
	"slices"
	"strconv"
)

// example: map[0:map[key:true] 1:map[key:false]] -> [map[key:true] map[key:false]]
func sliceMapToSlice(m map[string]interface{}) ([]interface{}, error) {
	var result []interface{}

	keys := make([]int, 0, len(m))
	for k := range m {
		key, err := strconv.Atoi(k)
		if err != nil {
			return nil, fmt.Errorf("array indexes must be integers: %w", err)
		}
		keys = append(keys, key)
	}
	for i := 0; i <= slices.Max(keys); i++ {
		val, ok := m[strconv.Itoa(i)]
		if !ok {
			result = append(result, nil)
			continue
		}
		result = append(result, val)
	}
	return result, nil
}
