#!/bin/sh
find /data/certs -type f -exec chmod u=rwx,og=r {} \;
