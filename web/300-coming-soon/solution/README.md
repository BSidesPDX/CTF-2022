# Solution

## Viewing Page in browser

Viewing page in browser looks pretty plain. View Source shows `js/scripts.js`, which could be interesting. We also see that the body tag has an onload hook:

```
<body onload="init()">
```

Let's look at the javascript:

```
$ curl -s localhost:1337/js/scripts.js
/*!
* Start Bootstrap - Coming Soon v6.0.6 (https://startbootstrap.com/theme/coming-soon)
* Copyright 2013-2022 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-coming-soon/blob/master/LICENSE)
*/

class Message {
    constructor(opcode, argument) {
        this.opcode = opcode;
        this.argument = argument;
    }
}

function twit() {
    let twitMsg = new Message("twit", null);
    let data = JSON.stringify(twitMsg);
    socket.send(data);
}

function init() {
    url = new URL(window.location.href);
    url.protocol = "ws";
    url.pathname = "/ws";

    var socket = new WebSocket(url.href);
    socket.onopen = () => {
        console.log("Successfully Connected");
        setInterval(function() {
            let heartbeat = new Message("heartbeat", null);
            let data = JSON.stringify(heartbeat);
            socket.send(data);
        }, 2000);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

    socket.onmessage = (event) => {
        // TODO: build out the dynamic content updater
        console.log(event.data);
    }
}
```
The code looks very simple, and it seems to establish a websocket, send stuff to it every 2 seconds, and print anything that comes across it

Looking in the console:
```
############################# Start Bootstrap Forms ################################                                                                           ######    To enable this form using Start Bootstrap Forms, simply sign up at:    ######                                                                           ######               https://startbootstrap/solution/contact-forms               ######                                                                           ####################################################################################
scripts.js:27 Successfully Connected
scripts.js:45 FIXME: Jim- fix the selector race
```

This gives a clue that the server is sending us messages and that there may be a user named jim as well as some kind of race condition?

If we wait log enough, we do see that the server hands out the flag:

```
BSidesPDX{DEFAULT_FLAG}
```

## Wrong Turns
### login

We easily find the `/login` path, we guess it takes form encoded post,

```
$ curl -X POST localhost:1337/login -d user=user,password=password
{"error":"Invalid request payload"}
```

Perhaps it takes json?

```
$ http localhost:1337/login user=admin password=admin
HTTP/1.1 403 Forbidden
Content-Length: 24
Content-Type: application/json
Date: Mon, 26 Sep 2022 03:17:58 GMT

{
    "error": "login failed"
}
```

Okay, maybe we need to guess the creds. As we saw from console, there is perhaps as user named `jim`. Let's try that as the username:

```
$ http localhost:1337/login user=jim password=admin
HTTP/1.1 302 Found
Content-Length: 19
Content-Type: application/json
Cookie: {"authentication":"successful","role":"user"};06072a903d7757beb104f1bc4cc35a70
Date: Mon, 26 Sep 2022 03:32:36 GMT
Location: admin

{
    "login": "success"
}
```

This looks interesting. Let's try that `/admin` path:

```
$ http localhost:1337/admin 'Cookie: {"authentication":"successful","role":"user"};06072a903d7757beb104f1bc4cc35a70'
HTTP/1.1 403 Forbidden
Content-Length: 29
Content-Type: application/json
Date: Mon, 26 Sep 2022 03:33:19 GMT

{
    "error": "permission denied"
}
```

Perhaps we need to change the role from `user` to `admin`?

```
$ http localhost:1337/admin 'Cookie: {"authentication":"successful","role":"admin"};06072a903d7757beb104f1bc4cc35a70'
HTTP/1.1 400 Bad Request
Content-Length: 20
Content-Type: application/json
Date: Mon, 26 Sep 2022 03:34:06 GMT

{
    "error": "mismatch"
}

```

This error leads us to believe that trailing value is used to validate? What kind of value is it? It looks like an md5 hash.

```
$ echo -n '{"authentication":"successful","role":"user"}'|md5sum
06072a903d7757beb104f1bc4cc35a70  -
```

What if we recalculate?

```
$ http localhost:1337/admin 'Cookie: {"authentication":"successful","role":"admin"};af1e79f14f245b1f9233f286e20b64ad'
HTTP/1.1 200 OK
Content-Length: 14
Content-Type: application/json
Date: Mon, 26 Sep 2022 03:54:38 GMT

{
    "good": "job"
}
```

This doesn't lead us to the flag.

### bg.mp4
Looking at bg.mp4, we see something in strings:

```
$ strings bg.mp4 | grep '.\{14\}'
isomiso2avc1mp41
x264 - core 157 r2935 545de2f - H.264/MPEG-4 AVC codec - Copyleft 2003-2018 - http://www.videolan.org/x264.html - options: cabac=1 ref=2 deblock=1:0:0 analyse=0x1:0x111 me=hex subme=6 psy=1 psy_rd=1.00:0.00 mixed_ref=1 me_range=16 chroma_me=1 trellis=1 8x8dct=0 cqm=0 deadzone=21,11 fast_pskip=1 chroma_qp_offset=-2 threads=12 lookahead_threads=2 sliced_threads=0 nr=0 decimate=1 interlaced=0 bluray_compat=0 constrained_intra=0 bframes=3 b_pyramid=2 b_adapt=1 b_bias=0 direct=1 weightb=1 open_gop=0 weightp=1 keyint=240 keyint_min=24 scenecut=40 intra_refresh=0 rc_lookahead=30 rc=crf mbtree=1 crf=22.0 qcomp=0.60 qpmin=0 qpmax=69 qpstep=4 vbv_maxrate=20000 vbv_bufsize=25000 crf_max=0.0 nal_hrd=none filler=0 ip_ratio=1.40 aq=1:1.00
B{~W&tmDGhe_|&
8@v?ci#S%[ICw`X
=pFTvfMwk-GP%R
`e>	C18cjJhNbcEn
k'h{{$Wm4[Ci\FN1#GH
NM$''f/xYN}	<(
1PeF>fOox,7t{z
;mInW\1#%c<]&{
*b&T%NaJUhy&EF
7~:fU]|Yt4(u`d
HandBrake 1.3.3 2020061300
This video is about My MovieH4sIEFwALGMA/2NvdWxkIHRoZSBmbGFnIGJlIGluIGhlcmU/AA3HsQ0AMQgDwFW8WlxgiSJ5WQ9FpiflqYzjEFS9/sDHa4iPDCOLyN0akKZZ7iYAAAA=
```

This looks like base64:
```
$ echo 'H4sIEFwALGMA/2NvdWxkIHRoZSBmbGFnIGJlIGluIGhlcmU/AA3HsQ0AMQgDwFW8WlxgiSJ5WQ9FpiflqYzjEFS9/sDHa4iPDCOLyN0akKZZ7iYAAAA='|base64 -d|hd
00000000  1f 8b 08 10 5c 00 2c 63  00 ff 63 6f 75 6c 64 20  |....\.,c..could |
00000010  74 68 65 20 66 6c 61 67  20 62 65 20 69 6e 20 68  |the flag be in h|
00000020  65 72 65 3f 00 0d c7 b1  0d 00 31 08 03 c0 55 bc  |ere?......1...U.|
00000030  5a 5c 60 89 22 79 59 0f  45 a6 27 e5 a9 8c e3 10  |Z\`."yY.E.'.....|
00000040  54 bd fe c0 c7 6b 88 8f  0c 23 8b c8 dd 1a 90 a6  |T....k...#......|
00000050  59 ee 26 00 00 00                                 |Y.&...|
00000056
```

We're getting somewhere, but what is that stuff?

```
$ echo 'H4sIEFwALGMA/2NvdWxkIHRoZSBmbGFnIGJlIGluIGhlcmU/AA3HsQ0AMQgDwFW8WlxgiSJ5WQ9FpiflqYzjEFS9/sDHa4iPDCOLyN0akKZZ7iYAAAA='|base64 -d>f && file f
f: gzip compressed data, has comment, last modified: Thu Sep 22 06:27:40 2022
````

What is it uncompressed?

```
$ echo 'H4sIEFwALGMA/2NvdWxkIHRoZSBmbGFnIGJlIGluIGhlcmU/AA3HsQ0AMQgDwFW8WlxgiSJ5WQ9FpiflqYzjEFS9/sDHa4iPDCOLyN0akKZZ7iYAAAA='|base64 -d|gzip -d
gur orfg guvatf pbzr gb gubfr jub jnvg
```

That looks like rot13:

```
$ echo 'H4sIEFwALGMA/2NvdWxkIHRoZSBmbGFnIGJlIGluIGhlcmU/AA3HsQ0AMQgDwFW8WlxgiSJ5WQ9FpiflqYzjEFS9/sDHa4iPDCOLyN0akKZZ7iYAAAA='|base64 -d|gzip -d|rot13
the best things come to those who wait
```

This is points us back to the idea that something is supposed to happen after a while

## .well-known/sshfp

We find that ".well-known/sshfp" gives a 200 response. From a quick googling, we see that this is standardized by IANA https://www.iana.org/assignments/well-known-uris/well-known-uris.xhtml, leading us to the structure. The response looks a little funny. Perhaps there's a flag there:

```
{
    "hosts": {
        "prod": [
            {
                "algo": "ecdsa-sha2-nistp256",
                "public_key": "VGhlIGZsYWcgY291bGQgYmUgaGlkZGVuIGluIHRoaXMgYmFzZTY0IGdhcmJhZ2UsIGJ1dCBpdCBpcyBub3QuIE1heWJlIGxhdGVyPwo="
            },
            {
                "algo": "ssh-rsa",
                "fp": "SHA256:cSQ6NTZEIXMpTDpEIEU5NiA0OUBENj8gQT0yOj9FSUU=",
                "port": 22
            }
        ]
    }
}
```

The host "prod" maybe tells us something. Let's save that for later. That public key doesn't look right. When looking at the example at https://cynthia.re/.well-known/sshfp, we can see the public key starts with A's and decoded reads:

```
$ echo AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAD+WrsYTcgGfmSZONzVScwk9+illAYwbZPewX3ihhajSJXVXrHYbHqiGFFSQZTYc2fKuq6Pgl5ed5lwvhzYtyQ=|base64 -d | hd
00000000  00 00 00 13 65 63 64 73  61 2d 73 68 61 32 2d 6e  |....ecdsa-sha2-n|
00000010  69 73 74 70 32 35 36 00  00 00 08 6e 69 73 74 70  |istp256....nistp|
00000020  32 35 36 00 00 00 41 04  00 fe 5a bb 18 4d c8 06  |256...A...Z..M..|
00000030  7e 64 99 38 dc d5 49 cc  24 f7 e8 a5 94 06 30 6d  |~d.8..I.$.....0m|
00000040  93 de c1 7d e2 86 16 a3  48 95 d5 5e b1 d8 6c 7a  |...}....H..^..lz|
00000050  a2 18 51 52 41 94 d8 73  67 ca ba ae 8f 82 5e 5e  |..QRA..sg.....^^|
00000060  77 99 70 be 1c d8 b7 24                           |w.p....$|
00000068
```

What does our's read as when decoded?

```
echo 'VGhlIGZsYWcgY291bGQgYmUgaGlkZGVuIGluIHRoaXMgYmFzZTY0IGdhcmJhZ2UsIGJ1dCBpdCBpcyBub3QuIE1heWJlIGxhdGVyPwo='|base64 -d
The flag could be hidden in this base64 garbage, but it is not. Maybe later?
```

This seems like a bit of a troll, perhaps there is a time based component to this challenge?

Let's look at the fp, to be thorough. Decoding it doesn't look right:

```
echo 'cSQ6NTZEIXMpTDpEIEU5NiA0OUBENj8gQT0yOj9FSUU='|base64 -d
q$:56D!s)L:D E96 49@D6? A=2:?EIE
```

Heading over to cyberchef, we try the rot47brute force with chosen plaintext of `BSidesPDX{`, which is our flag format. Sure, enough, that leads to a decode: `BSidesPDX{is the chosen plaintxt` 

This may have some clues, but is not the flag.
