# bboy-jam-assistant
Helps bboy jam organizers run better events

# Setup
Install Google Cloud command line tool:
https://cloud.google.com/sdk/docs/

If working on backend, make sure the following are installed:
- Golang: https://golang.org/doc/install#install
- Gcloud golang component: `gcloud components install app-engine-go`

# Env variables
add `path/to/bboy-jam-assistant` to your $PATH.

# gCloud setup
gcloud auth login
gcloud config set project bboy-jam-prod

# Tips and tricks
Verbose deploy: gcloud beta app deploy . --verbosity=debug
