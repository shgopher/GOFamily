#!/bin/zsh
###
 # @Author: shgopher shgopher@gmail.com
 # @Date: 2023-11-16 14:31:20
 # @LastEditors: shgopher shgopher@gmail.com
 # @LastEditTime: 2023-11-22 14:09:17
 # @FilePath: /GOFamily/lint.sh
 # @Description: 
 # 
 # Copyright (c) 2023 by shgopher, All Rights Reserved. 
### 


find_md(){

  local printed_files=()

  # 检查 zhlint 命令是否存在  
  if ! [ -x "$(command -v zhlint)" ]; then

    # 安装 zhlint
    npm install -g zhlint > /dev/null 
  fi

  for item in $(ls -R "$1"); do

    if [[ "$item" == *.md ]]; then
      
      path="$1/$item"

      if [[ ! " ${printed_files[@]} " =~ " $path " ]]; then
      
        echo "$path"

        # 调用zhlint格式化    
        zhlint "$path" --fix > /dev/null
      
        printed_files+=("$path")
      fi

    elif [[ -d "$1/$item" ]]; then

      find_md "$1/$item" 
    fi

  done  
}

# 调用 
find_md "$(pwd)"