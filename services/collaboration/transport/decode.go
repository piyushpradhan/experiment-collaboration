package collaboration

import (
	"collaboration/types"
	"context"
	"encoding/json"
	"io"
)

func decodeMessageRequest(_ context.Context, r io.Reader) (interface{}, error) {
	var msg types.Message
	err := json.NewDecoder(r).Decode(&msg)
	if err != nil {
		// TODO: Add log
		return nil, err
	}
	return msg, nil
}