# Drone ServerChan(Server酱)

[![Build Status](https://cloud.drone.io/api/badges/yakumioto/drone-serverchan/status.svg)](https://cloud.drone.io/yakumioto/drone-serverchan)
![Go Report](https://goreportcard.com/badge/github.com/yakumioto/drone-serverchan)
![Docker Pulls](https://img.shields.io/docker/pulls/yakumioto/drone-serverchan.svg)
[![](https://images.microbadger.com/badges/image/yakumioto/drone-serverchan.svg)](https://microbadger.com/images/yakumioto/drone-serverchan)
![GitHub release](https://img.shields.io/github/release/yakumioto/drone-serverchan.svg)

drone ServerChan(Server酱) 微信消息通知插件

## 简介

基于 [ServerChan(Server酱)](https://sct.ftqq.com/) 封装的微信消息通知插件

## 升级说明

> 因为微信发布公告将在4月底下线模板消息，Server酱开发了以企业微信为主的多通道新版（ Turbo版 sct.ftqq.com ）。旧版将在4月后下线，请尽快完成配置的更新。

`v1` 版本将在 4 月底不可用，请根据 [企业微信应用消息配置说明](https://sct.ftqq.com/forward) 进行升级。

`v2` 版本在不改变当前结构的情况下进行了升级，某些字段可能产生歧义。

- text 对应 turbo 版本中的 title 。

## 栗子

明文 key 配置

```yml
---
kind: pipeline
name: default

steps:
  - name: send-wechat
    image: yakumioto/drone-serverchan
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
    image: yakumioto/drone-serverchan
    settings:
      key:
        from_secret: {your key}
      text: {消息标题}
      desp: {消息内容 支持MarkDown}
```
