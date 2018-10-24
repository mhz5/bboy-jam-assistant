#!/bin/bash

VERBOSE_DEBUG="--go_debugging=true"
EXPOSE_TO_NETWORK="--host=0.0.0.0 --enable_host_checking=false"

dev_appserver.py ${VERBOSE_DEBUG} $EXPOSE_TO_NETWORK ../cmd/main/app_dev.yaml
