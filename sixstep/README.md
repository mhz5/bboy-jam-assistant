# API server

## Directory structure
/appengine: Contains .yaml config files and  main.go, which is entrypoint into the app. This directory is required because of an issue with Appengine doubly importing packages: https://stackoverflow.com/questions/26794225/google-go-appengine-imports-and-conflicts-when-serving-testing
/src: Contains entirety of source code.

## Run Locally
dev_appserver.py app.yaml

## Local Datastore Viewer
When running locally, datastore uses a local file that persists between invocations of the local server. More info here: https://cloud.google.com/appengine/docs/standard/python/tools/using-local-server
Access local datastore viewer here:
http://localhost:8000/datastore

## Deploy to gCloud
gcloud app deploy
