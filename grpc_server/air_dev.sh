#!/bin/bash

catch_signal() {
  echo "delete tmp_folder"
  rm -rf ./tmp
}

trap catch_signal SIGTERM SIGQUIT SIGINT

# Start air
air
