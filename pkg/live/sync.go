package live

import (
// "time"
)

// returns the channel
// this functions runs goroutine that writes true to that channel if next chunk of data is ready to read
// or writes false to chanel if reached EOF or encountered an error
// func (r *ReadWriter) StartLiveSync() chan bool {
// 	ch := make(chan bool)
// 	go func() {
// 		for {
// 			var buf [BufSize]byte
// 			_, err := r.Read(buf)
// 			if err != nil {
// 				ch <- false
// 				if err.Error() == "EOF" {
// 					return
// 				}
// 				// TODO error handling
// 				panic(err)
// 			}
// 			copy(r.curBuf, buf)
// 			ch <- true
// 			time.Sleep(time.Millisecond * time.Duration(r.MillisecondsPerPack()))
// 			continue
// 		}
// 	}()
// 	return ch
// }
