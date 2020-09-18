package util

import (
	"blogBackend/config"
	"github.com/rs/xid"
)

func GenerateId() string {
	return xid.New().String()
}

func GenerateSummary(content string) string {

	runes := []rune(content)
	length := len(runes)
	if length > config.Config.Summary.Length{
		length = config.Config.Summary.Length
	}

	return string(runes[:length])
}
