#!/usr/bin/env bash
# 主要是为了测试while
echo -e "请输入您的数字，并且确保它等于12"
read title
while [[ $title -eq 12 ]]; do # 特别注意此处的$title 要跟这个[ 有空格，不然就无法识别是两个命令
  echo "+++++++++"
done
