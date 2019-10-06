package camunda

import "lafitd/process/core/bpmn"

type BpmnProcess struct {
	name string
	key string
	activities []*bpmn.Activity
	sequenceFloes []*bpmn.SequenceFlow
}

func (BpmnProcess BpmnProcess)AddActivity(Activity bpmn.Activity, parentActivity bpmn.Activity, incomingActivity bpmn.Activity) error {

}

func (BpmnProcess BpmnProcess)AddSequenceFlow(flow bpmn.SequenceFlow) error  {

}