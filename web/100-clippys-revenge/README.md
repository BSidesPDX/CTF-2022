# 100 - Clippy's Revenge

## Description

This challenge aims to give out a free flag, hence the 100 points. If the player visits the site in Chrome (as of the time of writing) then the flag is immediately written to the player's clipboard. But there is a fake flag on the page in a little popup, which will tempt the player to put it on their clipboard. They will get a _different_ flag when they attempt to copy it, but that flag is still not right!

## Deploy

```
cd src && docker build -t clippy . && docker run -p 8000:8000 clippy
```

## Challenge

Dade wanted to hand out a free flag this year, he's gotten really lazy with the challenges. But Clippy had other ideas. Can you find the flag before you get board?

Author: [0xdade](https://0xda.de)
Flag: `BSidesPDX{th3_r34l_fl4g_w4s_th3_fr13nds_w3_m4d3_4l0ng_th3_w4y}`
