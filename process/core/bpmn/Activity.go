package bpmn


type Activity interface {
	//type: create, complete
	addListener(listenerType string, delegateExpresion string)
}

