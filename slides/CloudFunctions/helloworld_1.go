package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("helloworld_1.bash")
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
	gcloud alpha functions deploy HelloWorld --region europe-west1 --runtime go111 --trigger-http
	// END OMIT
	*/
}
