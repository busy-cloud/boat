//go:build jsoniter

package json

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var Marshal = json.Marshal

var Unmarshal = json.Unmarshal
