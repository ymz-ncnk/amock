package amock

// CheckCalls checks if all registered methods calls of the given mocks were
// executed. If so, returns empty result map. Otherwise, returns a map in which
// key is an index of the mock in mocks array param, and value is a
// MethodCallsInfo.
func CheckCalls(mocks []*Mock) (result map[int][]MethodCallsInfo) {
	result = make(map[int][]MethodCallsInfo)
	for i := 0; i < len(mocks); i++ {
		info := mocks[i].CheckCalls()
		if len(info) > 0 {
			result[i] = info
		}
	}
	return
}
