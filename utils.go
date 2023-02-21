package amock

import "github.com/ymz-ncnk/amock/core"

// CheckCalls checks if all registered method calls were made for each mock.
// If yes, it returns an empty result map. Otherwise, it returns a map where
// key is the index in the mocks param array and value is the
// MethodCallsInfo array.
func CheckCalls(mocks []*core.Mock) (result map[int][]core.MethodCallsInfo) {
	result = make(map[int][]core.MethodCallsInfo)
	for i := 0; i < len(mocks); i++ {
		info := mocks[i].CheckCalls()
		if len(info) > 0 {
			result[i] = info
		}
	}
	return
}
