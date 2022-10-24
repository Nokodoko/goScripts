package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/bitfield/script"
)

//(awsListUsers.sh) aws iam list-users | rg -i username | awk -F : '{print $2}' |sed 's/"//g' | sed 's/,.*//g' | dmenu

func main() {
	awsUsers, err := script.Args("aws iam list-users").Match("username")
	if err != nil {
		fmt.Println(err)
	}
	awk()
}

func awk(s string) (y string) {
	s := awsUsers()
	//i need to use the flags package (args[:] is a placeholder for just taking arguments given off the command-line)
	//USE THE SCRIPTS LIBRARY STOOPID -- will negate needing to use the flags package
	y := exec.Command("awk", args[:])
	var buf bytes.Buffer
	y.Stdout = &buf

	err := y.Run()
	if err != nil {
		fmt.Println(err)
	}
	return y
}

func sed(s string) (y string) {

}
