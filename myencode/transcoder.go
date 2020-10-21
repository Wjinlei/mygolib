package myencode

import "github.com/axgle/mahonia"

func GetDecoder(coding string) mahonia.Decoder {
	return mahonia.NewDecoder(coding)
}

func GetEncoder(coding string) mahonia.Encoder {
	return mahonia.NewEncoder(coding)
}
