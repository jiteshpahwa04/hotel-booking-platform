package utils

func FormatRoleIdsForQuery(roleIds []int64) []interface{} {
	formatted := make([]interface{}, len(roleIds))
	for i, id := range roleIds {
		formatted[i] = id
	}
	return formatted
}