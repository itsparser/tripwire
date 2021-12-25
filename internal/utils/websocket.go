package utils

import (
	"errors"
	"time"

	"golang.org/x/net/websocket"
)

// Dialer is the connection structure.
type Dialer struct {
	OnMessage     func([]byte)
	OnError       func(error)
	OnConnected   func()
	MatchMsg      func([]byte, []byte) bool
	Reconnect     bool
	PingMsg       []byte
	PingInterval  int
	ws            *websocket.Conn
	url, protocol string
	closed        bool
	msgQueue      []Msg
	pingTimer     time.Time
}

// Msg is the message structure.
type Msg struct {
	Body     []byte
	Callback func([]byte, *Dialer)
}

// Dial sets up the connection with the remote
// host provided in the url parameter.
// Note that all the parameters of the structure
// must have been set before calling it.
func (c *Dialer) Dial(url, protocol string) error {
	c.closed = true
	c.url = url
	c.protocol = protocol
	c.msgQueue = []Msg{}
	var err error
	c.ws, err = websocket.Dial(url, protocol, "http://localhost/")
	if err != nil {
		return err
	}
	c.closed = false
	if c.OnConnected != nil {
		go c.OnConnected()
	}

	go func() {
		defer c.close()

		for {
			var msg = make([]byte, 512)
			var n int
			if n, err = c.ws.Read(msg); err != nil {
				if c.OnError != nil {
					c.OnError(err)
				}
			}
			c.onMsg(msg[:n])
		}
	}()

	c.setupPing()

	return nil
}

// Send sends a message through the connection.
func (c *Dialer) Send(msg Msg) error {
	if c.closed {
		return errors.New("closed connection")
	}
	if _, err := c.ws.Write(msg.Body); err != nil {
		c.close()
		if c.OnError != nil {
			c.OnError(err)
		}
		return err
	}
	//if c.PingInterval > 0 && c.PingMsg != nil {
	//	c.pingTimer = time.Now().Add(time.Second * time.Duration(c.PingInterval))
	//}
	if msg.Callback != nil {
		c.msgQueue = append(c.msgQueue, msg)
	}

	return nil
}

// IsConnected tells wether the connection is
// opened or closed.
func (c *Dialer) IsConnected() bool {
	return !c.closed
}

func (c *Dialer) onMsg(msg []byte) {
	if c.MatchMsg != nil {
		for i, m := range c.msgQueue {
			if m.Callback != nil && c.MatchMsg(msg, m.Body) {
				go m.Callback(msg, c)
				// Delete this element from the queue
				c.msgQueue = append(c.msgQueue[:i], c.msgQueue[i+1:]...)
				break
			}
		}
	}
	// Fire OnMessage every time.
	if c.OnMessage != nil {
		go c.OnMessage(msg)
	}
}

func (c *Dialer) close() {
	err := c.ws.Close()
	if err != nil {
		return
	}
	c.closed = true
	if c.Reconnect {
		for {
			if err := c.Dial(c.url, c.protocol); err == nil {
				break
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func (c *Dialer) setupPing() {
	if c.PingInterval > 0 && len(c.PingMsg) > 0 {
		c.pingTimer = time.Now().Add(time.Second * time.Duration(c.PingInterval))
		go func() {
			for {
				if !time.Now().After(c.pingTimer) {
					time.Sleep(time.Millisecond * 100)
					continue
				}
				if err := c.Send(Msg{c.PingMsg, nil}); err != nil {
					if c.OnError != nil {
						c.OnError(err)
					}
					return
				}
				c.pingTimer = time.Now().Add(time.Second * time.Duration(c.PingInterval))
			}
		}()
	}
}
