#/bin/bash

# build.js
#
# Author::    Chirantan Mitra
# Copyright:: Copyright (c) 2017. All rights reserved
# License::   MIT

set -euo pipefail
IFS=$'\n\t'

make clean
make test
make compile
