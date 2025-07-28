#!/bin/bash


go build ./main

CODE=$(./main)

if [[ "$CODE" -ne 0 ]]; then
  echo "I expected error!!!"
else
  echo "Every thing works as expected!!!"
fi