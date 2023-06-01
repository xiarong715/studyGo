#!/bin/bash

PWD=`pwd`
MODULE_FILE=$PWD/go.mod
DIR_NAME=$(basename ${PWD}) # 获取最后一级目录名
MODULE_NAME=${DIR_NAME,,}   # 大写全部转小写

# 如果没有go.mod文件
# 获取文件夹名，变小写当做模块名
if [ ! -f $MODULE_FILE ]; then
    go mod init $MODULE_NAME
fi

go work use .

# 没有build/则创建
BUILD_NAME=build
if [ ! -d ${BUILD_NAME} ]; then
    mkdir ${BUILD_NAME}
fi

go build -o ${BUILD_NAME}/          # 输出的目录可自动创建，上面创建目录的过程可省略。

exec ${PWD}/${BUILD_NAME}/${MODULE_NAME}