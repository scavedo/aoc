#!/usr/bin/env bash

DAY=$1
echo "Running day $DAY"

go run "./src/D$DAY/solution.go"
