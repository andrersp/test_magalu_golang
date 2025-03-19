package testhelper

import (
	"bytes"
	"encoding/json"
	"io"
)

func RequestToPayload(request any) io.Reader {
	btsReq, _ := json.Marshal(request)
	return bytes.NewReader(btsReq)
}
