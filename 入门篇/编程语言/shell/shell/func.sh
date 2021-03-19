#!/usr/bin/env bash
#函数
function slep {
    echo "step two"
}

name () {
  local apple='12'
  echo apple
}

echo "step 1"
sleep 1s
slep
sleep 1s
echo "step 2"
name
# step 1
# step two
# step 2
