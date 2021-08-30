package method

import (
	"replaceText/method/brutal"
	"replaceText/rules"
)

type Method interface {
	Replace(string, rules.Rule) error
}

func New(methodName string) Method {
	switch methodName {

	case "brutal":
		return brutal.NewMethod()
	}

	return nil
}
