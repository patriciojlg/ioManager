package utils

import (
	"encoding/base64"
	"fmt"
	"ioManager/models"
	"regexp"
	"strings"
)

func AddPaddingIfRequired(base64Body string) string {
	// Agregar padding si falta
	padding := len(base64Body) % 4
	if padding != 0 {
		base64Body += strings.Repeat("=", 4-padding)
	}
	return base64Body
}
func SanitizeBase64(input string) string {
	// Elimina espacios, saltos de línea, etc.
	re := regexp.MustCompile(`[^A-Za-z0-9+/=]`)
	return re.ReplaceAllString(input, "")
}
func AnyInputBodyEncodedToBytes(args models.SaveInputFileArgs) ([]byte, error) {

	switch args.Encoded {
	case "base64":
		var base64Body string
		base64Body = args.Body
		//base64Body = SanitizeBase64(base64Body)
		base64Body = AddPaddingIfRequired(base64Body)
		decoded, err := base64.StdEncoding.DecodeString(base64Body)
		if err != nil {
			return nil, fmt.Errorf("error decoding base64: %w", err)
		}
		return decoded, nil
	default:
		return nil, fmt.Errorf("encoded no soportado: %s", args.Format)
	}
}
