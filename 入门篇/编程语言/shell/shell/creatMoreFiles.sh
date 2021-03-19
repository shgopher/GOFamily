#!/usr/bin/env bash
echo '创建很多月份其实挺容易。'
sleep 2s
echo '只需要先找到一个适合的目录，然后使用{}展开机制，就OK了'
sleep 2s
echo '第一步我们选取桌面作为路径'
sleep 5s
cd ~/Desktop
mkdir linshi
cd ./linshi
echo '然后我们使用mkdir 201{0..7}-{1..12}-{1..30}就可以了'
mkdir 201{0..7}-{1..12}-{1..30}
sleep 5s
echo "这样就OK了"
echo "很容易吧😀"
