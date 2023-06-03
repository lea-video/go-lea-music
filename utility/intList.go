package utility

import (
	"database/sql"
	"strconv"
	"strings"
)

func SplitList(list sql.NullString) ([]int, error) {
	if !list.Valid {
		return []int{}, nil
	}
	// Split the string by commas
	strArr := strings.Split(list.String, ",")

	// Convert string elements to integers
	var intArr []int
	for _, s := range strArr {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		intArr = append(intArr, num)
	}

	return intArr, nil
}
