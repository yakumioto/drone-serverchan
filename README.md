# Drone ServerChan(Server酱)

[![Build Status](https://drone.mioto.me/api/badges/yakumioto/drone-serverchan/status.svg)](https://drone.mioto.me/yakumioto/drone-serverchan)
![Go Report](https://goreportcard.com/badge/github.com/yakumioto/drone-serverchan)
![Docker Pulls](https://img.shields.io/docker/pulls/yakumioto/drone-serverchan.svg)
[![](https://images.microbadger.com/badges/image/yakumioto/drone-serverchan.svg)](https://microbadger.com/images/yakumioto/drone-serverchan)
![GitHub release](https://img.shields.io/github/release/yakumioto/drone-serverchan.svg)

drone ServerChan(Server酱) 微信消息通知插件

## 简介

基于 [ServerChan(Server酱)](http://sc.ftqq.com/3.version) 封装的微信消息通知插件

## 栗子

明文 key 配置

```yml
---
kind: pipeline
name: default

steps:
  - name: send-wechat
    image: yakumioto/serverchan
    settings:
      key: {your key}
      text: {消息标题}
      desp: {消息内容 支持MarkDown}
```

密文 key 配置

```yml
---
kind: pipeline
name: default

steps:
  - name: send-wechat
    image: yakumioto/serverchan
    settings:
      key:
        from_secret: {your key}
      text: {消息标题}
      desp: {消息内容 支持MarkDown}
```
