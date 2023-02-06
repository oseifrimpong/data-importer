#!/bin/sh
# Charles Cyril Nettey <cyril@keyspecs.com>
# This is the third and final script of the entryPoint call sequence

echo "teasapp-api entrypoint"
set -e

# Call command issued to the docker service
echo "teasapp-api exec: $@"
exec "$@"