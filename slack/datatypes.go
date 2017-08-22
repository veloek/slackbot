package slack

type rtmResponse struct {
	Ok    bool     `json:"ok"`
	Error string   `json:"error"`
	Url   string   `json:"url"`
	Team  teamInfo `json:"team"`
	Self  userInfo `json:"self"`
}

type teamInfo struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Domain         string `json:"domain"`
	EnterpriseId   string `json:"enterprise_id"`
	EnterpriseName string `json:"enterprise_name"`
}

type userInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type event struct {
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	Timestamp string `json:"ts"`
}

type message struct {
	event
	Id      uint   `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
