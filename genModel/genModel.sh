#!/usr/bin/env bash

# 使用方法： 第一个参数为数据库名 第二个参数为表名
# ./genModel.sh mall user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=127.0.0.1
port=3306
#dbname=looklook_$1
dbname=$1
username=root
passwd=123456


echo "开始创建库：$dbname 的表：$2"
# -cache=true 带缓存
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero

# 不带缓存
#goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" --style=goZero