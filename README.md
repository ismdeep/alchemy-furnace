# alchemy-furnace

炼丹炉

一个用于执行定期任务的服务程序。


## Docker

Config: `/service/config.yaml`

```yaml
bind: 127.0.0.1:8000
jwt: 7imaha6eYGqSR7f6eZ5JjzvbnMtk5xHP
auth:
  username: admin
  password: admin
```

Volume: `/service/data`