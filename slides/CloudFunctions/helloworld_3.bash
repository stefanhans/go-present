#!/bin/bash

GCP_PROJECT="cloud-functions-talk-22365"
GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/.secret/cloud-functions-talk-22365-6ba5c5af57f4.json"

gcloud alpha functions call --region=europe-west1 HelloWorld
