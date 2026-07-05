package hex2base64

import "cryptopals/utils"

func ConvertHexToBase64(h string) string {
	return utils.EncodeBytesToBase64(utils.DecodeHexToBytes(h))
}
