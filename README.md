![go workflow](https://github.com/ring0-rootkit/video-streaming-in-go/actions/workflows/go.yml/badge.svg)

live server and vod server

live server is broadcast to everyone - for optimisation
vod server is for particular user - for ability to rewind


srt_server requires libsrt to be installed on your system check
https://github.com/Haivision/srt?tab=readme-ov-file

on debian/ubuntu you can install it using
```
sudo apt install libsrt
```
on linux mint
```
sudo apt install libsrt-openssl-dev
```
