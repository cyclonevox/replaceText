package rules

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"replaceText/rules/simpleRule"
	"replaceText/util"
)

type Rule interface {
	// Get for get the value form input a key
	Get(key string) (value string)
	// Set support save the a pier of kv
	Set(key string, value string)
	// RangeRule support range the  kv about rules store and exec func
	RangeRule(f func(key, value string) bool)
	// SetRegList set the regexp and compile
	SetRegList(regList []string) error
	// GetRegList get the regexp list
	GetRegList() []*regexp.Regexp
	// GetRuleType get the type of rule
	GetRuleType() string
}

func New(ruleType string, path string) Rule {
	switch ruleType {
	case "simple":
		r := simpleRule.NewSimpleRule()

		if err := initRulesFromFiles(path, r); err != nil {
			fmt.Println(err)
		}

		return r

	default:
		return nil
	}
}

func initRulesFromFiles(path string, r Rule) (err error) {
	var files []string
	files, err = util.GetFilePath(path)

	for _, v := range files {
		fmt.Println("find match rule file : ", v)

		if err = rulesFromFile(v, r); nil != err {
			return
		}
	}

	return
}

func rulesFromFile(filePath string, r Rule) (err error) {

	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	var (
		keyLine   string
		valueLine string
		br        = bufio.NewReader(fi)
	)

	for i := 1; ; {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		switch i % 3 {

		case 1:
			keyLine = string(a)

		case 2:
			valueLine = string(a)

		case 0:
			r.Set(keyLine, valueLine)
		}

		i++
	}

	r.RangeRule(func(key, value string) bool {
		fmt.Println(key, value)
		return true
	})

	return
}
