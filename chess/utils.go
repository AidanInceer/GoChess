package chess

import (
	"fmt"
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

func IndexToRow(index int) (string, error) {
	rowMap := map[int]string{
		0: "1", 1: "2", 2: "3", 3: "4",
		4: "5", 5: "6", 6: "7", 7: "8",
	}

	if rank, exists := rowMap[index]; exists {
		return rank, nil
	}
	return "", fmt.Errorf("invalid index for row: %d", index)
}

// Convert index (0-7) back to file ("a"-"h")
func IndexToColumn(index int) (string, error) {
	columnMap := map[int]string{
		0: "a", 1: "b", 2: "c", 3: "d",
		4: "e", 5: "f", 6: "g", 7: "h",
	}

	if file, exists := columnMap[index]; exists {
		return file, nil
	}
	return "", fmt.Errorf("invalid index for column: %d", index)
}
