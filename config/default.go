package config

const DefaultConfigTemplate = `
# 是否显示进度
print-progress: false

# 并发线程数
concurrent: 200

# 检查间隔(分钟)
check-interval: 30

# 超时时间(毫秒)(节点的最大延迟)
timeout: 5000

# 质量等级，等级越高，稳定性越高，等级过高可能导致筛选出来的节点数量过少
quality-level: 1

# 测速地址(注意 并发数*节点速度<最大网速 否则测速结果不准确)
# 尽量不要使用Speedtest，Cloudflare提供的下载链接，因为很多节点屏蔽测速网站
# 建议使用自己上传到Cloudflare R2的文件
speed-test-url: https://github.com/AaronFeng753/Waifu2x-Extension-GUI/releases/download/v2.21.12/Waifu2x-Extension-GUI-v2.21.12-Portable.7z

# 最低测速结果舍弃(KB/s)
min-speed: 1024

# 下载测试时间(s)(与下载链接大小相关，默认最大测试10s)
download-timeout: 10


# mihomo api url(测试完成后自动更新mihomo订阅)
mihomo-api-url: ""

# mihomo api secret
mihomo-api-secret: ""

# 保存方法
# 目前支持的保存方法: r2, local, gist, webdav
save-method: local

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
