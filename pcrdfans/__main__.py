"""
 - Author: DiheChen
 - Date: 2021-06-22 00:36:00
 - LastEditTime: 2021-09-15 00:02:48
 - LastEditors: DiheChen
 - Description: None
 - GitHub: https://github.com/DiheChen
"""
from time import time
from aiohttp import ClientSession
from fastapi import FastAPI

from config import Config

session = None
app = FastAPI()


@app.on_event("startup")
async def _():
    global session
    session = session or ClientSession()


@app.on_event("shutdown")
async def _():
    global session
    if session:
        await session.close()


headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
    "Authorization": Config.AUTH_KEY
}
payload = {"_sign": "a", "def": [], "nonce": "a",
           "page": 1, "sort": 1, "ts": 0, "region": 1}


@app.post("/x/v1/search")
async def _(item: dict):
    payload.update({"def": item["def"], "ts": int(
        time()), "region": item["region"]})
    try:
        async with session.post("https://api.pcrdfans.com/x/v1/search", headers=headers, json=payload) as resp:
            return await resp.json()
    except Exception as e:
        return {"error": e}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app=app, port=7777, host="0.0.0.0")
