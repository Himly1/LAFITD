package utils

import (
	"github.com/sirupsen/logrus"
	"LAFITD/expression/subject"
)

func IsThePropertyIsOneOfThePropertiesOfTheSubject(property subject.Property, subject subject.Subject) bool {
	if subject == nil {
		logrus.Error("The subject is null pointer")
		return false
	}

	for _, e := range subject.GetSubjectInfo().Properties {
		if e.Code == property.Code {
			return true
		}

		if e.IsSubject {
			flag := IsThePropertyIsOneOfThePropertiesOfTheSubject(property, e.Subject)
			if flag {
				return true
			}
		}
	}

	return false
}