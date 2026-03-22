package worker

func MergeState(currentState map[string]interface{}, nodeID string, output map[string]interface{}) map[string]interface{} {
	if currentState == nil {
		currentState = make(map[string]interface{})
	}
	currentState[nodeID] = output
	return currentState
}
