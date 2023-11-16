#!/bin/bash
###
 # @Author: shgopher shgopher@gmail.com
 # @Date: 2023-11-16 14:31:20
 # @LastEditors: shgopher shgopher@gmail.com
 # @LastEditTime: 2023-11-16 14:40:54
 # @FilePath: /GOFamily/lint.sh
 # @Description: 
 # 
 # Copyright (c) 2023 by shgopher, All Rights Reserved. 
### 

find_md(){

  local printed_files=()

  for item in $(ls -R "$1")
  do
    if [[ "$item" == *.md ]]; then

      path="$1/$item"

      if [[ ! " ${printed_files[@]} " =~ " $path " ]]; then
        
        echo "$path"
        
        # 调用zhlint修复格式
        zhlint "$path" --fix
        
        printed_files+=("$path")
      fi

    elif [[ -d "$1/$item" ]]; then

      find_md "$1/$item"

    fi

  done

}

find_md "$(pwd)"