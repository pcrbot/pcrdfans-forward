"""
 - Author: DiheChen
 - Date: 2021-06-22 00:36:00
 - LastEditTime: 2021-06-22 00:36:13
 - LastEditors: DiheChen
 - Description: None
 - GitHub: https://github.com/Chendihe4975
"""
from time import time
from aiohttp import ClientSession
from fastapi import FastAPI

from config import Config

app = FastAPI()


@app.post('/x/v1/search')
async def _(item: dict):
    header = {
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36',
        'authorization': Config.AUTH_KEY
    }
    payload = {"_sign": "a", "def": item["def"], "nonce": "a",
               "page": 1, "sort": 1, "ts": int(time()), "region": item["region"]}
    try:
        async with ClientSession() as session:
            async with session.post('https://api.pcrdfans.com/x/v1/search', headers=header, json=payload) as resp:
                return await resp.json()
    except Exception as e:
        return {"error": e}


if __name__ == '__main__':
    import uvicorn
    uvicorn.run(app=app, port=7777, host='0.0.0.0')