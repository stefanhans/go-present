package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("helloworld_2.bash")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{""}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	/*
		   // START OMIT
	curl https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/HelloWorld
		   // END OMIT
	*/
}
