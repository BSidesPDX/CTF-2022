# 100 - Discordia

## Description

A simple trivia game via discord bot, where each question is enciphered with a different strategy.

1. ROT13
2. A1Z26
3. Morse Code
4. Multi-tap phone

## Deploy

This requires a discord bot token at launch time, for a bot that is connected to the BSides PDX discord server.

```
cd src && docker build -t discordia . && docker run -e DISCORD_TOKEN=<discord_token> discordia
```

## Challenge

Let's see how much you know about BSides PDX! I'm going to ask you four questions. But some wires got crossed and I can't find the original questions, only these weird codes. Can you give me the plaintext answers? Good luck!

Author: [aceroni](https://aceroni.com)


## Flag

Flag: `BSidesPDX{s0m3t1m3s_4_C7F_f33ls_l1k3_4_7r1v14l_pur5u17}`
