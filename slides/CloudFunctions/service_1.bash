#!/bin/bash

GCP_PROJECT="cloud-functions-talk-22365"
GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/.secret/cloud-functions-talk-22365-6ba5c5af57f4.json"

curl -s -d "hi https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/HelloWorld hi says hello world" https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/register