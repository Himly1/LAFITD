package bpmn

import "LAFITD/expression/core/condition"

type SequenceFlow struct {
	sourceActivity *Activity
	targetActivity *Activity
	condition condition.Condition
}