# pcrdfans_forward
啊这。

## 这是什么?

一个随便瞎写的项目。

## 这能做什么 ?

小明部署了一个 pcrbot, 但他发现 pcrdfans 的 APIkey 已经关闭申请渠道了, 碰巧小明的 hxd 申请到了 APIkey, 聪明的小明将这个项目打包发给了他的 hxd, 最后两个人都用上了 pcrdfans 的 API。

## 这怎么用 ?

### 如果你是小明的 hxd :

在使用了 pcrdfans api 的服务器上 :

```bash
pip3 install poetry

git clone https://github.com/pcrbot/pcrdfans_forward

cd pcrdfans_forward

poetry install
```

编辑 `config.py`, 填入自己的 APIkey。然后运行项目:

```bash
poetry run python api.py
```

### 如果你是小明 :

把这个项目打包发给你的 hxd , 修改你 bot 的 `arena.py`, 将 Pcrdfans 的 API 地址由原来的 `https://api.pcrdfans.com/x/v1/search` 修改为 `http://{你 hxd 的服务器 ip}:7777/x/v1/search`, 然后谢谢你的 hxd。

## 关于开源

使用 MIT 开源协议, 作者只保留版权, 不做其它限制。

## 老婆！老婆老婆老婆！老婆老婆老婆老婆老婆！

反正出了事我不管.jpg