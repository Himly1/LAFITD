package core


type FirstConditionUnsetError struct {

}

func (FirstConditionUnsetError)Error() string  {
	return "Please set the first condition before do the operate"
}