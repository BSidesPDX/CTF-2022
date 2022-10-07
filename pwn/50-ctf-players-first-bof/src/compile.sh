#!/bin/sh
gcc main.c -fno-stack-protector -o 50-bof

cp ./50-bof ../distFiles