# pcrdfans_forward
啊这。

## Pcrdfans API 转发

此分支为 golang 版, 使用标准库 。

Python 版请看 [这个分支](https://github.com/pcrbot/pcrdfans_forward/tree/fastapi) 。

## 开始使用

Fork 本项目, 在 [main.go](https://github.com/pcrbot/pcrdfans_forward/blob/gin/main.go) 第二十行填入您在 <https://pcrdfans.com/bot> 申请的 API key, commit 后在您 fork 的项目的 action 中下载 与您使用的操作系统 / 平台对应的可执行文件, 运行。

然后将您服务器的公网 ip 告知于您想要分享 API 的那个人, 并让 他 / 她 把 http 请求的 url 由原来的 `https://api.pcrdfans.com/x/v1/search` 修改成 `http://{你服务器的公网 ip}:7777/x/v1/search` 。

若您不知道您使用的 bot 发起 http 请求部分的代码在哪, 您可尝试使用全文搜索 `https://api.pcrdfans.com/x/v1/search` 。

## 免责声明

本项目数据由 `pcrdfans.com` 提供, 但它不隶属于 `pcrdfans.com`, 也不受其支持。