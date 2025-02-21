# landns

## 快速开始

准备配置文件和 Docker Compose 声明文件

```sh
git clone https://github.com/117503445/landns.git
cd landns/docs/example
```

启动服务

```sh
docker compose up -d
```

验证

```sh
dig @127.0.0.1 archlinux.lan A -p 4053
# archlinux1.lan.         60      IN      A       192.168.60.100
dig @127.0.0.1 archlinux1.lan A -p 4053
# archlinux1.lan.         60      IN      A       192.168.61.100
```
