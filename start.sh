#!/bin/sh

#if catch error -> quit directly
set -e

echo "start the app"
exec "$@"
