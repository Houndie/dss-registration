#!/bin/bash

if [[ -z "${DSS_POSTGRESURL}" ]]; then
	if [[ -n "$DATABASE_URL" ]]; then
		export DSS_POSTGRESURL=$DATABASE_URL
	else
		echo "Unable to start up, DSS_POSTGRES_URL unset"
		exit 1
	fi
fi

if [[ -z "${DSS_PORT}" ]]; then
	if [[ -n "${PORT}" ]]; then
		export DSS_PORT=$PORT
	fi
fi

until nc -z $(echo $DSS_POSTGRESURL | sed "s%^.*@\(.*\):\(.*\)\/.*$%\1 \2%"); do
	echo "Waiting for database ..."; \
	sleep 1; \
done; \

../server
