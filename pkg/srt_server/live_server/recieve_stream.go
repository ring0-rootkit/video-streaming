package live_server

import (
	"github.com/ring0-rootkit/video-streaming-in-go/pkg/live"
)

func (s *SRTServer) recieveStream() {
	for !s.streamer.Connected || !s.isRunning {
	}
	s.log.Info("recieving data from streamer")

	for s.isRunning {
		buf := make([]byte, live.BufferSize)
		n, err := s.streamer.Socket.Read(buf)
		if err != nil {
			s.log.Error(err.Error())
			s.Close()
			return
		}
		_, err = s.readWriter.Write(buf[:n])
		if err != nil {
			s.log.Error(err.Error())
			s.Close()
			return
		}
	}
	s.log.Info("stream ended, streamer is disconected")
}
