package replaceText

import (
	"fmt"

	"replaceText/method"
	"replaceText/rules"
	"replaceText/util"
)

type ReplaceText struct {
	rule   rules.Rule
	method method.Method

	// target file path list
	TargetPaths []string
}

func NewReplace(TargetPaths []string, rulesPath string, regList ...string) (r *ReplaceText, err error) {
	r = &ReplaceText{
		TargetPaths: TargetPaths,
		rule:        rules.New("simple", rulesPath),
		method:      method.New("brutal"),
	}

	return
}

func (r *ReplaceText) Exec() (err error) {
	var v string
	var filePath string

	for _, v = range r.TargetPaths {

		var files []string
		files, err = util.GetFilePath(v)

		for _, filePath = range files {

			fmt.Println("handle file : " + filePath)

			if err = r.method.Replace(filePath, r.rule); nil != err {
				return
			}
		}

	}

	return
}

func (r *ReplaceText) matchSingleLineByReg(line string) (exist bool, newLine string, err error) {
	return
}
