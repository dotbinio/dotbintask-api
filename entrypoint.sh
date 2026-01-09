#!/bin/bash

# Default to UID/GID 1000 if not specified
PUID=${PUID:-1000}
PGID=${PGID:-1000}

echo "Starting with PUID=${PUID} and PGID=${PGID}"

# Modify the appuser UID and GID to match the requested values
groupmod -o -g "${PGID}" appuser
usermod -o -u "${PUID}" appuser

# Execute the API as the appuser
exec su appuser -c "./taskwarrior-api"

