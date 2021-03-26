#!/usr/bin/env python3
import requests
import json

response = requests.get("http://mbmd:8080/api/last/SUNSPEC1.1")
if (not response.ok):
    exit(1)
data = json.loads(response.content)
print(f'{data["Unix"]}: Export = {data["Export"]}')
print(f'{data["Unix"]}: Power = {data["Power"]}')
print(f'{data["Unix"]}: DCPower = {data["DCPower"]}')
