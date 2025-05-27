package domain

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Код игры, прячется в URL https://believe-or-not.com/g/{GameCode}
type GameCode string

var (
	mu            sync.Mutex
	lastTimestamp uint64
	seq           uint64
	baseEpoch     = time.Date(2024, time.May, 21, 0, 0, 0, 0, time.UTC).Unix()
)

func GenerateGameCode() GameCode {
	const (
		seqBits        = 6
		seqMask uint64 = (1 << seqBits) - 1
		codeLen        = 6
	)

	now := time.Now().Unix()
	delta := uint64(now - baseEpoch)

	mu.Lock()
	if delta == lastTimestamp {
		seq = (seq + 1) & seqMask
	} else {
		lastTimestamp = delta
		seq = 0
	}
	localSeq := seq
	mu.Unlock()

	id := (delta << seqBits) | localSeq

	code := strings.ToUpper(strconv.FormatUint(id, 36))
	if len(code) < codeLen {
		code = strings.Repeat("0", codeLen-len(code)) + code
	}

	return GameCode(fmt.Sprintf("%s-%s-%s", code[0:2], code[2:4], code[4:6]))
}
