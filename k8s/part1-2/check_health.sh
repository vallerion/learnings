#!/bin/sh

status_code=$(curl --write-out %{http_code} --silent --output /dev/null locahost:8081)

if [[ "$status_code" -ne 200 ]] ; then
  echo "Site status changed to $status_code"
else
  exit 0
fi