#!/bin/sh

# Substitute the PORT environment variable in the Envoy configuration template
envsubst '${PORT}' < /etc/envoy/envoy.yaml.template > /etc/envoy/envoy.yaml

# Start Envoy
/usr/local/bin/envoy -c /etc/envoy/envoy.yaml
