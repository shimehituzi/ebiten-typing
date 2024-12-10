#!/bin/bash

if [ -z "$CMD" ]; then
  echo "Error: CMD environment variable is not set."
  exit 1
fi

if [ ! -d "cmd/$CMD" ]; then
  echo "Error: cmd/$CMD directory does not exist."
  exit 1
fi

go build -o "./air/$CMD" "./cmd/$CMD"
