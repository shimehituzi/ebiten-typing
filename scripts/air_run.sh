#!/bin/bash

if [ -z "$CMD" ]; then
  echo "Error: CMD environment variable is not set."
  exit 1
fi

if [ ! -f "air/$CMD" ]; then
  echo "Error: Failed to build cmd/$CMD."
  exit 1
fi

exec ./air/$CMD
