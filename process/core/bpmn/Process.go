package bpmn

type Process interface {
	ToJsonDefinition() string
	ToMarkDownDoc() string
	ToBpmnXml() string
}

