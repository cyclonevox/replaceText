package main

import (
	"flag"
	"fmt"

	"replaceText/replace"
)

func main() {
	ruleFilePath := flag.String("r", "./rules", "rule path")
	flag.Parse()
	targetList := flag.Args()

	fmt.Println(*ruleFilePath)
	fmt.Println(targetList)

	var r *replaceText.ReplaceText
	var err error

	if r, err = replaceText.NewReplace(targetList, *ruleFilePath); nil != err {
		fmt.Println(err)
	}

	err = r.Exec()
	if err != nil {
		fmt.Println(err)
	}

}
