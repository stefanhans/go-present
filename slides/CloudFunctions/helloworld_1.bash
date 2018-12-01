#!/bin/bash

GCP_PROJECT="gke-serverless-211907"
GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/.secret/gke-serverless-211907-181ed186fa7f.json"

gcloud alpha functions deploy HelloWorld --region europe-west1 --memory 128MB --runtime go111 --trigger-http
