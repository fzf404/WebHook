#!/bin/bash 
cd /www/website/www.fzf404.top # 进入目录
git pull # 拉取代码
npm i # 安装依赖
npm run build # 构建静态页面
