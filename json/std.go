//go:build !jsoniter && !go_json && !sonic

package json

import "encoding/json"

var Marshal = json.Marshal

var Unmarshal = json.Unmarshal
