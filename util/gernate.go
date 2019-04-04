package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func GetGernate() string {
	seed := time.Now().UnixNano()
	text := strconv.FormatInt(seed, 10)
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))

}
