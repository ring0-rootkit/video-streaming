![go workflow](https://github.com/ring0-rootkit/video-streaming-in-go/actions/workflows/go.yml/badge.svg)

# Video streaming

This project is live streaming webapp inspired by [Twitch](twitch.tv). This project We made for fun 😋.

# Usage

You can install dependencies using
```
make build
```
And start the server using 
```
make server
```

After that webserver will be launched.

Now you can go to `localhost://8080/register`, create an account and after you logged in go to you dashboard.

In dashboard press `Your StreamKey` and then copy your stream key from that page.

Now you can open [OBS-studio](https://obsproject.com/),
go to settings -> stream, here select Custom server, and in url enter `srt://localhost:42069`,
as stream key past the key you just copied.

And you all set 

Now press `start stream` in OBS-Studio, and then go to `localhost:8080/{yourusername}` (without curly brackets)

Enjoy!

# Technologies (Backend) - by [ring0-rootkit](https://github.com/ring0-rootkit)

<Design scheme will be soon>

That app consists of: 

`Live server` (for recieving live stream from streamer, save it to database and send it to users),

`VOD` (Video On Demand) server (sends VOD to users),

`WebServer` (handles http connections, sends html/css/js to users)

Live and VOD servers are using SRT protocol and h264 codec to send videos.

## Stack

[Golang](https://go.dev/)
[srtgo-library](https://github.com/Haivision/srtgo)
[Postgres](https://www.postgresql.org/)
[sqlx](https://github.com/jmoiron/sqlx)

# Technologies (Frontend) - by TBD
