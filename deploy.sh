#!/bin/bash

YAML=./www/app.yaml
COMMIT=
yq write --inplace $YAML env_variables.COMMIT $(git rev-parse --short HEAD)
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
yq write --inplace $YAML env_variables.LASTMOD $LASTMOD

/usr/local/google_appengine/appcfg.py --oauth2 update .

#
# restore committed values
#
yq write --inplace $YAML env_variables.COMMIT dev
yq write --inplace $YAML env_variables.LASTMOD dev

