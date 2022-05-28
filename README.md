# go-wows-bot

go-wows-bot 由Cunoe基于golang开发的一个用于丘丘群的战舰世界机器人助手

更易于部署的机器人

## 食用方法

1. 下载构建好的go-wows-bot release包 双击打开后即可
2. 下载[go-cqhttp](https://github.com/Mrs4s/go-cqhttp/releases)并选择websocket反向代理 根据直营填写QQ账号密码，并将以下内容写入config.yml文件结尾并且重新打开go-cqhttp
```yaml
- ws-reverse:
      universal: ws://127.0.0.1:9000/ws/
      reconnect-interval: 5000
      middlewares:
        <<: *default
```
3. 您已完成部署

## 感谢

[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)

[战舰世界API平台](https://wows.linxun.link/)