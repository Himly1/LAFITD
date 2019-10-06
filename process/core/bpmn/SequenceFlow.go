package bpmn

import "lafitd/expression/core/condition"

type SequenceFlow struct {
	sourceActivity *Activity
	targetActivity *Activity
	condition condition.Condition
}