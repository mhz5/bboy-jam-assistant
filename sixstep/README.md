# API server

# Run Locally
dev_appserver.py app.yaml

# Local Datastore Viewer
When running locally, datastore uses a local file that persists between invocations of the local server. More info here: https://cloud.google.com/appengine/docs/standard/python/tools/using-local-server
Access local datastore viewer here:
http://localhost:8000/datastore

# Deploy to gCloud
gcloud app deploy
