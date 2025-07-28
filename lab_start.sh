#!/bin/bash


rm ./*.txt

go build

echo "I will run 1"

./nawaphon.com 1

echo "I will run 2 with 4 processes"

./nawaphon.com 2 &
PID_1=$!

./nawaphon.com 2 &
PID_2=$!

./nawaphon.com 2 &
PID_3=$!

./nawaphon.com 2 &
PID_4=$!

wait "$PID_1" "$PID_2" "$PID_3" "$PID_4"

echo "every thing done"