package websocket

import (
	"encoding/json"
	"log"
)

const (
	CreateProject string = "create/project"
	CreateTask    string = "create/task"
	EditTask      string = "edit/task"
)

type Event struct {
	EventName string      `json:"event_name"`
	Data      interface{} `json:"data"`
}

func NewEventFromRaw(rawData []byte) (*Event, error) {
	event := new(Event)
	log.Println(rawData)
	err := json.Unmarshal(rawData, &event)

	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
