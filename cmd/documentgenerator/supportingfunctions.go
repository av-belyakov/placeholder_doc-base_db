package documentgenerator

// searchEventSource выполняет поиск источника события
func searchEventSource(fieldBranch string, value any) (string, bool) {
	if fieldBranch != "source" {
		return "", false
	}

	if v, ok := value.(string); ok {
		return v, true
	}

	return "", false
}
