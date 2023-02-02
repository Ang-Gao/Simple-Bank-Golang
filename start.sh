#!/bin/sh
#use the start script to do something before launching the app
#if catch error -> quit directly
set -e

echo "start the app"
exec "$@"
