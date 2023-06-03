package sqlite

import "strings"

func buildINStatement(ids []int) (string, []interface{}) {
	placeholders := strings.Repeat("?, ", len(ids)-1) + "?"

	typecastIDs := make([]interface{}, 0, len(ids))
	for _, m := range ids {
		typecastIDs = append(typecastIDs, m)
	}

	return placeholders, typecastIDs
}
