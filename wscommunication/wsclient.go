package wscommunication

import (
	"log"

	"golang.org/x/net/websocket"
)

const localhost string = "htt://localhost"

// ReadFromC1 connects to a given url received as parameters
// and sends the data through a channel
func ReadFromC1(url string, chn chan []byte) {
	con, error := websocket.Dial(url, "ws", localhost)
	if error != nil {
		log.Println(error)
	}
	go func() {
		defer close(chn)
		defer con.Close()

		for {
			msg := make([]byte, 1024)
			lenght, err := con.Read(msg)
			if err != nil {
				log.Println(err)
				break
			}
			if lenght != 0 {
				msg = msg[:lenght]
				chn <- msg
			}
		}
	}()
}
