package dal

import "fmt"

type LangAble interface {
	GetSuffix() string
	GetType() string
	GetCmd(src, target string) string
}
type CLang struct{}
type Java struct{}

func (c CLang) GetSuffix() string {
	return "c"
}
func (c CLang) GetType() string {
	return "c"
}
func (c CLang) GetCmd(src, target string) string {
	return fmt.Sprintf("#!/bin/bash\ngcc -O2 -w -std=c99 %s -lm -o %s", src, target)
}
