#!/bin/bash
# measure.sh - run from inside each model's project folder
echo "== tests ==";  go test -race -count=1 -cover ./... 2>&1 | tee tests.txt
echo "== bench =="; go test -bench . -benchmem -count=5 ./... 2>&1 | tee bench.txt
echo "== size ==";  find . -name '*.go' | xargs wc -l | tail -1
echo "== test count =="; grep -rhc "func Test" --include='*_test.go' . | paste -sd+ - | awk '{s+=$1} END {print s+0}'
