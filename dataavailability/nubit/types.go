package nubit

import (
	"encoding/json"
	"github.com/rollkit/go-da"
	"reflect"
)

type BatchDAData struct {
	ID []da.ID `json:"id,omitempty"`
}

// write a function that encode batchDAData struct into ABI-encoded bytes
func (b *BatchDAData) Encode() ([]byte, error) {
	return json.Marshal(b)
}
func (b *BatchDAData) Decode(data []byte) error {
	return json.Unmarshal(data, &b)
}
func (b *BatchDAData) IsEmpty() bool {
	return reflect.DeepEqual(b, BatchDAData{})
}
