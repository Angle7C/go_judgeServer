package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := exec.Command("/judgeserver/compiler.sh").CombinedOutput()
	fmt.Printf("%s\n%s", string(output), err.Error())
}
