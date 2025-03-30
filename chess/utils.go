package chess

import (
	"strings"
)

func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

func RowToIndex(rank string) (int, error) {

	type RowMap struct {
		Map map[string]int
	}

	var rowMap = RowMap{
		Map: map[string]int{
			"1": 0, "2": 1, "3": 2, "4": 3,
			"5": 4, "6": 5, "7": 6, "8": 7,
		},
	}
	return rowMap.Map[rank], nil

}

// Convert file ("a"-"h") to index (0-7)
func ColumnToIndex(file string) (int, error) {
	type ColumnMap struct {
		Map map[string]int
	}

	var columnMap = ColumnMap{
		Map: map[string]int{
			"a": 0, "b": 1, "c": 2, "d": 3,
			"e": 4, "f": 5, "g": 6, "h": 7,
		},
	}

	return columnMap.Map[file], nil
}
