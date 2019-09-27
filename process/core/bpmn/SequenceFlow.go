package bpmn


type SequenceFlow struct {
	sourceActivity *Activity
	targetActivity *Activity
	condition string
}