#!/bin/sh
antlr4 \
  -Dlanguage=Go \
  -o parser \
  "$@" \
  ./Antex.g4 && echo "OK"
