#!/bin/bash

set -e

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <year> <day>"
  exit 1
fi

YEAR=$1
DAY=$2
BASE_DIR="./base"
TARGET_DIR="./${YEAR}/${DAY}"

if [ ! -d "$BASE_DIR" ]; then
  echo "Error: Base directory '$BASE_DIR' does not exist."
  exit 1
fi

mkdir -p "$TARGET_DIR"

cp -r "$BASE_DIR/"* "$TARGET_DIR/"

echo "Created folder structure at $TARGET_DIR"
