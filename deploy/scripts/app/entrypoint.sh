#!/bin/sh
# This is the third and final script of the entryPoint call sequence

echo "api entrypoint"
set -e

# Call command issued to the docker service
echo "exec: $@"
exec "$@"