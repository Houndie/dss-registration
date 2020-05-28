#/bin/bash

/migrate.linux-amd64 -path=/migrations -database "postgres:///$POSTGRES_USER?host=/var/run/postgresql&user=$POSTGRES_USER" up
