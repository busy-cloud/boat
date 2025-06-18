//go:build sonic

package json

import "github.com/bytedance/sonic"

var Marshal = sonic.Marshal

var Unmarshal = sonic.Unmarshal
