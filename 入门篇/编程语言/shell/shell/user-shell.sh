#!/usr/bin/env bash
# 用户和shell的交互

echo -e "开始"
read -sp "请输入您的数字在这个地方输入" int int2
if (("$int")) && (("$int2"));then
  echo "恭喜您您的数字不等于0"
elif (("$int2")) ; then
  echo "第二个选项大于20"
else
  echo "您的数字等于0或者是其它字符"
fi
