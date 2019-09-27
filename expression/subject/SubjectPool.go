package subject

import (
	"github.com/sirupsen/logrus"
	"pm/erros/expression/subject"
)

type Pool interface {
	GetSubjectByCode(code string) (*subject.NotFoundError, Subject)
	GetSubjects() []Subject
	AddSubject(subject Subject) *subject.DuplicatedError
	AddSubjects(subjects []Subject) error
}

type DefaultPool struct {
	subjects map[string]Subject
}

func (DefaultPool *DefaultPool) GetSubjectByCode(code string) (*subject.NotFoundError, Subject)  {
	err := ensureSubjectExists(code, DefaultPool)
	if err != nil {
		return err, nil
	}
	
	return nil, DefaultPool.subjects[code]
}

func ensureSubjectExists(code string, pool *DefaultPool) *subject.NotFoundError {
	exists := isSubjectExists(code, pool)
	if !exists {
		return &subject.NotFoundError{Code: code}
	}
	
	return nil
}

func isSubjectExists(code string, pool *DefaultPool) bool {
	targetSubject := pool.subjects[code]
	if targetSubject == nil {
		return false
	}else {
		return true
	}
}

func (DefaultPool *DefaultPool) AddSubject(subject Subject) *subject.DuplicatedError {
	err := ensureNoDuplicatedSubject(subject.GetSubjectInfo().Code, DefaultPool)
	if err != nil {
		return err
	}

	DefaultPool.subjects[subject.GetSubjectInfo().Code] = subject
	return nil
}


func ensureNoDuplicatedSubject(code string, pool *DefaultPool) *subject.DuplicatedError {
	exists := isSubjectExists(code, pool)
	if exists {
		return &subject.DuplicatedError{Code:code}
	}

	return nil
}

func (DefaultPool *DefaultPool) AddSubjects(subjects []Subject) error {
	for _, e := range subjects {
		err := DefaultPool.AddSubject(e)
		if err != nil {
			logrus.Error("Add the subject failed. The code is ", e.GetSubjectInfo().Code)
		}
	}

	return nil
}


func (DefaultPool *DefaultPool) GetSubjects() []Subject  {
	var arr []Subject
	for _, v := range DefaultPool.subjects {
		arr = append(arr, v)
	}

	return arr
}