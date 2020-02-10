#!/bin/bash
#
# deploy to AppEngine
#

set -o errexit
set -o pipefail
set -o nounset

YAML=./app.yaml
COMMIT=
yq write --inplace $YAML env_variables.COMMIT $(git rev-parse --short HEAD)
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
yq write --inplace $YAML env_variables.LASTMOD $LASTMOD

#/usr/local/google_appengine/appcfg.py --oauth2 update .
gcloud app deploy --project=regexplanet-go

#
# restore committed values
#
yq write --inplace $YAML env_variables.COMMIT dev
yq write --inplace $YAML env_variables.LASTMOD dev

