"""
Tryhackme Solution: 
https://tryhackme.com/room/introtopython
"""
import base64

f = open("encodedflag.txt","r")
c = f.read()

for _ in range(5):
  c = base64.b16decode(c)

for _ in range(5):
  c = base64.b32decode(c)

for _ in range(5):
  c = base64.b64decode(c)

print(c)
f.close()
