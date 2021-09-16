# pcrdfans_forward
啊这。

## Pcrdfans API 转发

此分支为 golang 版, 使用标准库 。

Python 版请看 [这个分支](https://github.com/pcrbot/pcrdfans_forward/tree/fastapi) 。

两分支略有差别:

性能上: ~~Golang 版性能约为 Python 版 5 - 6 倍。~~ 未计量

占用上: Golang 版约占 1.4mb, Python 版约占 24.4mb 。

## 开始使用

使用之前请确保您正确安装 golang 环境。

```bash
git clone https://github.com/pcrbot/pcrdfans_forward

cd pcrdfans_forward/pcrdfans
```

编辑 `main.go`, 在第二十行填入您的 Pcrdfans API Key, 保存。

若您希望编译成适用于您当前使用的系统 / 平台的二进制文件, 请使用以下命令:

```bash
go build -trimpath -ldflags="-s -w" -o "pcrdfans.exe" .
```

若您使用 Windows 电脑, 并且您希望编译成 Linux 可执行的二进制文件, 请使用 cmd ( 而不是 Powershell ) 执行以下命令:

```cmd
set GOOS=linux
set GOARCH=amd64
go build -trimpath -ldflags="-s -w" -o "pcrdfans-linux-amd64" .
```

若您使用其它系统 / 平台, 您可从 [这里](https://www.google.com/search?q=golang%E4%BA%A4%E5%8F%89%E7%BC%96%E8%AF%91) 获得帮助。

运行编译好的二进制文件 ( 请在能使用 Pcrdfans API 的机器上运行 ):

```bash
chmod +x pcrdfans-linux-amd64

./pcrdfans-linux-amd64
```

建议使用 tmux, screen 之类的窗口复用器使其在后台运行。

若您希望直接运行源代码, 您可执行以下命令 ( 必须安装 Golang 环境 ):

```bash
go run main.go
```

然后将您服务器的公网 ip 告知于您想要分享 API 的那个人, 并让 他 / 她 把 http 请求的 url 由原来的 `https://api.pcrdfans.com/x/v1/search` 修改成 `http://{你服务器的公网 ip}:7777/x/v1/search` 。

若您不知道您使用的 bot 发起 http 请求部分的代码在哪, 您可尝试使用全文搜索 `https://api.pcrdfans.com/x/v1/search` 。

## 免责声明

呜呜呜求光佬放过。