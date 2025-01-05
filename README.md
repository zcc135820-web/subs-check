# 可用性检测

## 预览

![preview](./doc/images/preview.png)

## 功能

- 检测可用性
- 检测平台解锁情况
- 合并,去重,重命名
- 分类保存结果

## 特点

- 支持多平台
- 支持多线程

## TODO

- [ ] 适配多种订阅格式
- [ ] 支持更多的保存方式

## 使用方法

### Docker

```bash
docker run -itd \
    --name mihomo-check \
    -v /path/to/config.yaml:/app/config/config.yaml \
    -e TZ=Asia/Shanghai \
    --restart=always \
    bestrui/mihomo-check:latest
```

### 源码直接运行

```bash
go run main.go -f /path/to/config.yaml
```

### 二进制文件运行

直接运行即可,会在当前目录生成配置文件

