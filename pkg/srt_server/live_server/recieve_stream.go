package live_server

import (
	"fmt"

	"github.com/ring0-rootkit/video-streaming-in-go/pkg/live"
)

func (s *SRTServer) recieveStream() {
	for !s.streamer.Connected || !s.isRunning {
	}
	s.log.Info("recieving data from streamer")

	s.recieveStreamHeader()
	for s.isRunning {
		buf := make([]byte, live.BufferSize)
		_, err := s.streamer.Socket.Read(buf)
		if err != nil {
			s.log.Error(err.Error())
			s.Close()
			return
		}
		_, err = s.readWriter.Write(buf[:])
		if err != nil {
			s.log.Error(err.Error())
			s.Close()
			return
		}
		// s.log.Info(fmt.Sprintf("streamer sent %d bytes of data", n))
	}
	s.log.Info("stream ended, streamer is disconected")
}

func (s *SRTServer) recieveStreamHeader() {
	var err error
	var n = 0
	buf := make([]byte, live.BufferSize)
	for n == 0 {
		n, err = s.streamer.Socket.Read(buf[:live.BufferSize])
		if err != nil {
			s.log.Error(err.Error())
			s.Close()
			return
		}
	}
	var tmp []byte = make([]byte, live.BufferSize)
	copy(tmp, buf[:n])
	s.streamHeader = append(s.streamHeader, tmp)
	for range PacketsInHeader + 1 {
		n, err = s.streamer.Socket.Read(buf)
		if err != nil {
			s.log.Error(err.Error())
			s.Close()
			return
		}
		var tmp []byte = make([]byte, live.BufferSize)
		copy(tmp, buf[:n])
		s.streamHeader = append(s.streamHeader, tmp)
	}

	s.isStreamHeaderSet = true
	s.log.Info(fmt.Sprintf("header recieved"))
}
