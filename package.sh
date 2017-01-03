#/bin/bash

set -euo pipefail
IFS=$'\n\t'

cd out
zip -9 -o ../dashy.zip dashy public/*
