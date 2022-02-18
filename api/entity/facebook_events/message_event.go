package facebook_events

type Event struct {
	Entry []Entry `json:"entry"`
}

type Entry struct {
	Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Message   Message   `json:"message"`
	Timestamp int       `json:"timestamp"`
}

type Sender struct {
	ID string `json:"id"`
}

type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	Text string `json:"text"`
}
