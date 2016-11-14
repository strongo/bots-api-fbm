package fbm_bot_api

import (
	"github.com/strongo/bots-framework/core"
	"time"
)

// ReceivedMessage ...
type ReceivedMessage struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

// Entry ...
type Entry struct {
	ID        string       `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

func (e Entry) GetID() interface{} {
	return e.ID
}

func (e Entry) GetTime() time.Time {
	return time.Unix(e.Time, 0)
}

// Messaging ...
type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   *Message  `json:"message"`
	Postback  *Postback `json:"postback"`
	Delivery  *Delivery `json:"delivery"`
}


type Actor struct {
	ID string `json:"id"`
}

func (a Actor) GetID() interface{} {
	return a.ID
}

func (a Actor) GetFirstName() string {
	panic("Not supported")
}

func (a Actor) GetLastName() string {
	panic("Not supported")
}

func (a Actor) GetUserName() string {
	panic("Not supported")
}

func (a Actor) Platform() string {
	return "fbm"
}

// Sender ...
type Sender struct {
	Actor
}

// Recipient ...
type Recipient struct {
	Actor
}

type Postback struct {
	Payload Payload `json:"payload"`
}

type Delivery struct {
	Watermark  int64      `json:"watermark"`
	MessageIDs MessageIDs `json:"mids"`
}

type MessageID string

type MessageIDs []MessageID

// Message ...
type Message struct {
	MID         string       `json:"mid"`
	Seq         int          `json:"seq"`
	MText       string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"paylaod"`
}

type Payload struct {
	Url string `json:"url"`
}

// SendMessage ...
type SendMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   struct {
		Text string `json:"text"`
	} `json:"message"`
}


func (m Message) IntID() int64 {
	panic("Not supported")
}

func (m Message) StringID() string {
	return m.MID
}

func (m Message) Sequence() int {
	return m.Seq
}

func (m Message) Text() string {
	return m.MText
}

func (m Message) Contact() bots.WebhookContactMessage {
	panic("Not supported")
}

