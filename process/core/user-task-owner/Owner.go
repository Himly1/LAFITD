package user_task_owner

import (
	user_task_owner "LAFITD/erros/process/core/user-task-owner"
	"LAFITD/expression/subject"
	"LAFITD/utils"
	"sync"
)

type Owner interface {
	AddUserOwner(subject subject.Subject, property subject.Property) *user_task_owner.BadUserOwnerError
	AddDepartmentOwner(departmentCode string) *user_task_owner.BadUserOwnerError
	GetOwnerDefinition() *OwnerDefinition
}

type OwnerDefinition struct {
	UserOwners []*SubjectPropertyOwner

	//department code
	DepartmentOwners []string
}

type SubjectPropertyOwner struct {
	Subject subject.Subject
	Owner subject.Property
}

type SimpleOwner struct {
	userOwners []*SubjectPropertyOwner
	departmentCode []string
	lock sync.RWMutex
}

func (SimpleOwner *SimpleOwner) AddUserOwner(subject subject.Subject, property subject.Property) *user_task_owner.BadUserOwnerError {
	SimpleOwner.lock.Lock()
	defer SimpleOwner.lock.Unlock()
	err := ensureThePropertyCanAsUserOwner(subject, property)
	if err != nil {
		return err
	}

	owner := &SubjectPropertyOwner{Subject:subject, Owner:property}
	SimpleOwner.userOwners = append(SimpleOwner.userOwners, owner)

	return nil
}

func ensureThePropertyCanAsUserOwner(aSubject subject.Subject, property subject.Property) *user_task_owner.BadUserOwnerError {
	flag := utils.IsThePropertyIsOneOfThePropertiesOfTheSubject(property, aSubject)
	if !flag {
		return &user_task_owner.BadUserOwnerError{Msg: "The property is`t one of the subject`s property"}
	}

	if property.ValueType != subject.NUMBER {
		return &user_task_owner.BadUserOwnerError{Msg: "The value type of the owner must be number type"}
	}

	return nil
}

func (SimpleOwner *SimpleOwner)AddDepartmentOwner(code string) *user_task_owner.BadUserOwnerError {
	SimpleOwner.lock.Lock()
	defer SimpleOwner.lock.Unlock()

	err := ensureNoDuplicatedDepartmentOwner(SimpleOwner.departmentCode, code)
	if err != nil {
		return err
	}

	SimpleOwner.departmentCode = append(SimpleOwner.departmentCode, code)
	return nil
}

func ensureNoDuplicatedDepartmentOwner(owners []string, code string) *user_task_owner.BadUserOwnerError {
	for _, e := range owners {
		if e == code {
			return &user_task_owner.BadUserOwnerError{Msg:"Duplicated department owner. The code is " + code}
		}
	}

	return nil
}

func (SimpleOwner *SimpleOwner)GetOwnerDefinition() *OwnerDefinition {
	SimpleOwner.lock.RLock()
	defer SimpleOwner.lock.RUnlock()
	return &OwnerDefinition{UserOwners:SimpleOwner.userOwners, DepartmentOwners:SimpleOwner.departmentCode}
}