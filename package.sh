#/bin/bash

# package.js
#
# Author::    Chirantan Mitra
# Copyright:: Copyright (c) 2017. All rights reserved
# License::   MIT

set -euo pipefail
IFS=$'\n\t'

cd out
zip -9 -o ../dashy.zip dashy public/*
