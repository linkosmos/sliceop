package sliceop

// Prefill - prefil array with values
func Prefill(size int, symbol string) (output []string) {
	output = make([]string, size)
	for i := 0; i < size; i++ {
		output[i] = symbol
	}
	return output
}

// Map - maps array of strings with func
func Map(f func(input string) string, input ...string) (output []string) {
	if len(input) == 0 {
		return input
	}
	for _, key := range input {
		output = append(output, f(key))
	}
	return output
}

// Includes - whether key includes array
func Includes(input []string, key string) bool {
	for _, k := range input {
		if k == key {
			return true
		}
	}
	return false
}

// NotIncludes - whether key is not included in array
func NotIncludes(input []string, key string) bool {
	return !Includes(input, key)
}

// Reject - reject specified keys
func Reject(input []string, toReject ...string) (output []string) {
	if len(toReject) == 0 {
		return input
	}
	for _, key := range input {
		if NotIncludes(toReject, key) {
			output = append(output, key)
		}
	}
	return output
}

// Select - select and return specified keys
func Select(input []string, toSelect ...string) (output []string) {
	if len(toSelect) == 0 {
		return input
	}
	for _, key := range input {
		if Includes(toSelect, key) {
			output = append(output, key)
		}
	}
	return output
}

// Unique - returns unique values for given input
func Unique(input ...string) []string {
	if len(input) == 0 {
		return input
	}
	var output []string
	for _, key := range input {
		if NotIncludes(output, key) {
			output = append(output, key)
		}
	}
	return output
}
