package utils

import (
	"encoding/base64"
	"fmt"
)

func AnyInputBodyEncodedToBytes(format string, body any) ([]byte, error) {

	switch format {
	case "base64":
		var base64Body string
		base64Body = body.(string)
		return base64.StdEncoding.DecodeString(base64Body)
	default:
		return nil, fmt.Errorf("formato no soportado: %s", format)
	}
}
