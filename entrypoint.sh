#!/bin/sh

set -e

if [ "$1" = 'webui' ]; then

    for file in `grep -r geodex.org static | cut -f1 -d: | sort -u`; 
    do 
      sed -i "s/geodex\.org/${GEODEX_HOST}/g" $file; 
    done

    /webui -host ${GEODEX_HOST}
else
    exec "$@"
fi
