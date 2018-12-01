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
	   GCP_PROJECT="gke-serverless-211907"
	   GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/.secret/gke-serverless-211907-181ed186fa7f.json"

	   gcloud alpha functions deploy HelloWorld --region europe-west1 --runtime go111 --trigger-http
	   // END OMIT
	*/
}
