package bpmn

type Process interface {
	AddActivity(Activity Activity, parentActivity Activity, incomingActivity Activity) error
	AddSequenceFlow(flow SequenceFlow) error
	GetProcessDefinition() ProcessDefinition
}

type ProcessDefinition struct {
	Name string
	Key string
	Activities []Activity
	SequenceFlows []*SequenceFlow
}