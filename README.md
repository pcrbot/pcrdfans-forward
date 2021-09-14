# pcrdfans_forward
啊这。

## Pcrdfans API 转发

此分支为 Python 版, 使用的 Web 框架: [FastAPI](https://github.com/tiangolo/fastapi) 。

## 开始使用

提示: 请在能使用 Pcrdfans API 的服务器上运行本项目。

```bash
git clone -b fastapi https://github.com/pcrbot/pcrdfans_forward

pip3 install poetry

poetry install
```

在 `pcrdfans/config.py` 里面填入你的 API Key, 然后执行:

```bash
poetry run python pcrdfans
```

然后将您服务器的公网 ip 告知于您想要分享 API 的那个人, 并让 他 / 她 把 http 请求的 url 由原来的 `https://api.pcrdfans.com/x/v1/search` 修改成 `http://{你服务器的公网 ip}:7777/x/v1/search` 。

若您不知道您使用的 bot 发起 http 请求部分的代码在哪, 您可尝试使用全文搜索 `https://api.pcrdfans.com/x/v1/search` 。
