package zerologgelfoutput

func getStringFromMap(m map[string]interface{}, key, def string) string {
	if val, ok := m[key].(string); ok {
		return val
	}

	return def
}
