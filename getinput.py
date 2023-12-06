#!/usr/bin/python
import requests
import sys

cookies = {
    'session': 'COOKIE',
}

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0',
    'Accept': 'text/html',
    'Connection': 'keep-alive',
}

dayNo = sys.argv[1]
#print(dayNo)

response = requests.get(f"https://adventofcode.com/2023/day/{dayNo}/input", cookies=cookies, headers=headers)
print(response.text.rstrip())
