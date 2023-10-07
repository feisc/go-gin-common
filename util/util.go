package util

import (
	"context"
	"encoding/base64"
	"strconv"
	"unicode/utf8"

	"github.com/pkg/errors"
)

// request id
const RidFlag = "request_id"

func GetRequestIdFromCtx(ctx context.Context) interface{} {
	return ctx.Value(RidFlag)
}
func Base64Decode(s string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringToBool(s string) (bool, error) {
	if s == "false" {
		return false, nil
	} else if s == "true" {
		return true, nil
	}
	return false, errors.New("params error: should be true or false")
}

// ToValidUTF8 treats s as UTF-8-encoded bytes and returns a copy with each run of bytes
// representing invalid UTF-8 replaced with the bytes in replacement, which may be empty.
func ToValidUTF8(s, replacement []byte) []byte {
	b := make([]byte, 0, len(s)+len(replacement))
	invalid := false // previous byte was from an invalid UTF-8 sequence
	for i := 0; i < len(s); {
		c := s[i]
		if c < utf8.RuneSelf {
			i++
			invalid = false
			b = append(b, byte(c))
			continue
		}
		_, wid := utf8.DecodeRune(s[i:])
		if wid == 1 {
			i++
			if !invalid {
				invalid = true
				b = append(b, replacement...)
			}
			continue
		}
		invalid = false
		b = append(b, s[i:i+wid]...)
		i += wid
	}
	return b
}

func IsStringSliceContains(s []string, v string) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}

func MergeStringSliceAsSet(origin, input []string) []string {
	output := origin
	for _, s := range input {
		if !IsStringSliceContains(output, s) {
			output = append(output, s)
		}
	}

	return output
}
