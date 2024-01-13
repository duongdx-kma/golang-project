#!/bin/sh

if [ ! -f ".env" ]; then
    echo Development environment
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
    echo "APP_ENV=$APP_ENV" >> .env

    echo "AWS_REGION=$AWS_REGION" >> .env
    echo "SECRET_MANAGER_KEY=$SECRET_MANAGER_KEY" >> .env

    cat .env
else
    echo "already have environment file"
fi

# start server
./artifacts
