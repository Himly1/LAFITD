package camunda

import "pm/process/core/bpmn"

type BpmnProcess struct {
	name string
	key string
	activities []*bpmn.Activity
	sequenceFloes []*bpmn.SequenceFlow
}

func (BpmnProcess *BpmnProcess) ToJsonDefinition() string  {

}

func (BpmnProcess *BpmnProcess)ToMarkDownDoc() string  {

}

func (BpmnProcess *BpmnProcess)ToBpmnXml() string  {

}