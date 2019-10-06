package runner

import runners2 "lafitd/runner/runners"

type Runner interface {
	Run() error
}

func Run() error {
	runners := getRunners()

	for _, e := range runners {
		err := e.Run()
		if  err != nil {
			return err
		}
	}

	return nil
}

func getRunners() []Runner {
	var runners []Runner
	subjectRegisterRunner := &runners2.SubjectRegisterRunner{}
	runners =  append(runners,subjectRegisterRunner)
	return runners
}
