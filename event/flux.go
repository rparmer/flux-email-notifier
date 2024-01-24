package event

import (
	"encoding/json"

	eventv1 "github.com/fluxcd/pkg/apis/event/v1beta1"
)

func FromJson(b []byte) (*eventv1.Event, error) {
	var e eventv1.Event
	err := json.Unmarshal(b, &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func ToJsonIndent(e *eventv1.Event) ([]byte, error) {
	b, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		return nil, err
	}
	return b, nil
}
