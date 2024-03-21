#!/bin/bash
envsubst '${APP_PORT} ${PORTAINER_PORT}' < /etc/nginx/nginx.conf.template > /etc/nginx/conf.d/nginx.conf
exec "$@"