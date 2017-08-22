package slack

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
)

const apiUrl = "https://slack.com/api/rtm.connect?token=%s"
const protocol = ""
const origin = "https://api.slack.com/"

type Connection struct {
	conn *websocket.Conn
}

func (c *Connection) Init(token string) error {
	url := fmt.Sprintf(apiUrl, token)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("API failed with status code: %d", resp.StatusCode)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	var respObj rtmResponse
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return err
	}

	if !respObj.Ok {
		err = fmt.Errorf("Slack error: %s", respObj.Error)
		return err
	}

	c.conn, err = websocket.Dial(respObj.Url, protocol, origin)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) Close() error {
	err := c.conn.Close()
	return err
}

func (c *Connection) GetMessage() string {
	var e event
	for {
		buf := make([]byte, 2048)
		n, err := c.conn.Read(buf)
		if err != nil || n == 0 {
			continue
		}
		data := buf[:n]

		err = json.Unmarshal(data, &e)
		if err != nil {
			fmt.Printf("Error while unmarshaling event: %s\n", err)
			continue
		}
		//fmt.Printf("DEBUG: Event type = %s\n", e.Type)

		if e.Type == "message" && e.Subtype == "" {
			var m message
			err = json.Unmarshal(data, &m)
			if err != nil {
				fmt.Printf("Error while unmarshaling message: %s\n", err)
				continue
			}
			return m.Text
		}
	}
}
