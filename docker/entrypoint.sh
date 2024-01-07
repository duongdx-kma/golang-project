#!/bin/sh

if [ "$APP_ENV" = "dev" ]; then
    echo Development environment
    rm .env
    echo "Creating env file"
    echo "DB_DRIVER=$DB_DRIVER" >> .env
    echo "DB_HOST=$DB_HOST" >> .env
    echo "DB_USER=$DB_USER" >> .env
    echo "DB_PASSWORD=$DB_PASSWORD" >> .env
    echo "DB_DATABASE=$DB_DATABASE" >> .env
    echo "DB_PORT=$DB_PORT" >> .env
    echo "CLIENT_ORIGIN=$CLIENT_ORIGIN" >> .env
    echo "JWT_SECRET=$JWT_SECRET" >> .env
    echo "APP_PORT=$APP_PORT" >> .env
else
    echo "Production environment"
fi

# start server
./artifacts
