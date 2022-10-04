# 100 - Hollywood more then meets the ear (Solution)

When looking at the mp3 file it looks like any normal mp3 when in fact it is a mp3 and a zip combined together. The zipped folder is appended after the mp3 footer so the song still plays but if you open the file with a compression tool you will get the zip

To get the flag first bump the song written by [0xdade](https://0xda.de)

Once you have finished bumping the song simply run `unzip hollywood_bsides.mp3`. You will be presented with a txt file called qwsrt that contains a base 64 encoded string. Run `echo "QnNpZGVzUERYe0gzYWQzcnNfMW5fZjFsM3NfYXIzX3ByM3R0eV9jMDBsfQo="|base64 -d` and you will get the flag `BsidesPDX{H3ad3rs_1n_f1l3s_ar3_pr3tty_c00l}`
