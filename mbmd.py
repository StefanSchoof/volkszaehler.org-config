#!/usr/bin/env python3
import requests
import json

response = requests.get("http://localhost:8080/api/last/SUNSPEC1.1")
if (not response.ok):
    exit(1)
data = json.loads(response.content)
print(f'{data["Unix"]}: Export = {data["Export"]}')
