#!/bin/bash

cd /Users/stefan/go/src/github.com/stefanhans/go-present/slides/CloudFunctions/helloworld

GCP_PROJECT="cloud-functions-talk-22365"
GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/.secret/cloud-functions-talk-22365-6ba5c5af57f4.json"

gcloud alpha functions deploy HelloWorld --region europe-west1 --runtime go111 --trigger-http
