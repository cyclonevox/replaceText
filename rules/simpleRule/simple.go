package simpleRule

import "regexp"

// SimpleRule is a simple implement, not support concurrency
type SimpleRule struct {
	// store in map for rules
	rules map[string]string
	// regExp list generate by compiling
	regExp []*regexp.Regexp
}

func NewSimpleRule() *SimpleRule {
	return &SimpleRule{
		rules:  make(map[string]string),
		regExp: make([]*regexp.Regexp, 0, 0),
	}
}

// Get for get the value form input a key
func (sr *SimpleRule) Get(key string) (value string) {
	return sr.rules[key]
}

// Set support save the a pier of kv
func (sr *SimpleRule) Set(key string, value string) {
	sr.rules[key] = value
}

// RangeRule support range the rules store and exec func
func (sr *SimpleRule) RangeRule(f func(key, value string) bool) {
	for k, v := range sr.rules {
		// f(k, v) false,then break
		if !f(k, v) {
			break
		}
	}
}

// SetRegList set the regexp and compile
func (sr *SimpleRule) SetRegList(regList []string) (err error) {
	for _, v := range regList {
		var re *regexp.Regexp
		if re, err = regexp.Compile(v); nil != err {
			return
		}
		sr.regExp = append(sr.regExp, re)
	}
	return
}

// GetRegList get the regexp list
func (sr *SimpleRule) GetRegList() []*regexp.Regexp {
	return sr.regExp
}

func (sr *SimpleRule) GetRuleType() string {
	return "simple map"
}
