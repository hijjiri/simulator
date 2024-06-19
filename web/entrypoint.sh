#!/bin/sh

# Substitute the PORT environment variable in the Nginx configuration template
envsubst '$PORT' < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf

# Start Nginx
nginx -g 'daemon off;'
