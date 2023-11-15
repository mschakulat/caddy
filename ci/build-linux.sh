#!/bin/bash

set -e

make release

echo "Packaging Binaries"

tar -zcvf "release/$1.tar.gz" release/caddy release/caddy-shim