package config

const DefaultConfigTemplate = `
# 是否显示进度
print-progress: true

# 并发线程数
concurrent: 200

# 检查间隔(分钟)
check-interval: 30

# 超时时间(毫秒)(节点的最大延迟)
timeout: 5000

# 测速地址(注意 并发数*节点速度<最大网速 否则测速结果不准确)
speed-test-url: https://speed.cloudflare.com/__down?bytes=52428800

# 最低测速结果舍弃(KB/s)
min-speed: 1024

# 下载测试大小(MB)(目前无效，根据提供的下载链接控制，提供的下载链接多大，则下载多大)
download-size: 20

# 上传测试大小(MB)(目前无效，没有上传测速功能)
upload-size: 20

# mihomo api url(测试完成后自动更新mihomo订阅)
mihomo-api-url: https://api.mihomo.me/v3/

# mihomo api secret
mihomo-api-secret: ""

# 保存方法
# 目前支持的保存方法: r2, local, gist, webdav
save-method: r2

# webdav
webdav-url: "https://example.com/dav/"
webdav-username: "admin"
webdav-password: "admin"

# gist id
github-gist-id: ""

# github token
github-token: ""

# github api mirror
github-api-mirror: ""

# 将测速结果推送到Worker的地址
worker-url: https://example.worker.dev

# Worker令牌
worker-token: 1234567890

# 重试次数(获取订阅失败后重试次数)
sub-urls-retry: 3

# 订阅地址 支持 clash/mihomo/v2ray/base64 格式的订阅链接
sub-urls:
  - https://example.com/sub.txt
  - https://example.com/sub2.txt

`
