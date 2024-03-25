#!/bin/sh
envsubst '${APP_PORT} ${PORTAINER_PORT} ${MINIO_PORT} ${MINIO_CONSOLE_PORT}' < /etc/nginx/nginx.conf.template > /etc/nginx/conf.d/nginx.conf
exec "$@"