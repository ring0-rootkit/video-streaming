server:
	@go run ./cmd/main.go

client:
	@go run ./cmd/client/main.go

file:
	@ffmpeg -re -i srt://localhost:42069 -c:v h264 -c:a copy sample_videos/output.mp4
