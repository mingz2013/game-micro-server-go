
# 服务架构说明

## 客户端层

### web-front
前端项目，管理页面前端，数据平台前端等，后续可拆分成不同项目

### command-client
命令行页面client


### robot
不同的游戏rob逻辑


### client
客户端，游戏


---

## 接入层

### static-web
静态http

### web-proxy REST API
系统的web API入口，用于管理，或数据入口，后续可拆分成不同项目


### connector
用于客户端连接入口

### login
登陆入口


---

## 中间服务层

### daemon
用于接收api连接，更新配置，创建新的进程等

### common
暂时用于其他功能，后续可拆分


### game-manager
管理不同游戏的进程


---

## 具体游戏逻辑层


### game-*
游戏逻辑，不同游戏的逻辑


---