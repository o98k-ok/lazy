package file

import (
	"encoding/json"
	"time"
)

type Codec interface {
	Encode(t any) error
	Decode(t any) error
}

type JsonCodec struct {
	*json.Encoder
	*json.Decoder
}

type Persistence[T any] struct {
	Time    time.Time
	Content T
}
