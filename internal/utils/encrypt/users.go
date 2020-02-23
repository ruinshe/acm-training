package encrypt

import (
	"crypto/sha512"
	"encoding/hex"
)

const (
	prefix = "salt_for_uestc"
	suffix = "salt_for_acm_training"
)

// Pasword - entrypt the user's password and genrate a hash code for it.
func Pasword(password string) (entrypted string, err error) {
	hasher := sha512.New()
	if _, err = hasher.Write([]byte(prefix)); err != nil {
		return
	}
	if _, err = hasher.Write([]byte(password)); err != nil {
		return
	}
	if _, err = hasher.Write([]byte(suffix)); err != nil {
		return
	}
	entrypted = hex.EncodeToString(hasher.Sum(nil))
	return
}
