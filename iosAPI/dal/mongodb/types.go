package mongodb

type UserInfo struct {
	Username string
	Password string
	Groups   []int
}

type Event struct {
	EventID  int    `json:"eventID"`
	TimeFrom string `json:"timeFrom"`
	TimeTo   string `json:"timeTo"`
	Theme    string `json:"theme"`
	Desc     string `json:"desc"`
}

type GroupInfo struct {
	GroupID   int
	Users     []string
	Events    []Event
	GroupName string
}

type EventListOpt struct {
	Data map[int][]Event `json:"data"`
}

type WSopt struct {
	GroupID int     `json:"groupID"`
	Events  []Event `json:"events"`
}
