package Miner

import (
	"encoding/hex"
	"time"
)


func truncateAndFormat(data []byte) string {
    if len(data) == 0 {
        return "N/A"
    }
    hexStr := hex.EncodeToString(data)
    if len(hexStr) > 10 {
        return hexStr[:10] + "..."
    }
    return hexStr
}

func formatTime(timestamp uint64) string {
    t := time.Unix(int64(timestamp), 0)
    return t.Format("2006-01-02 15:04:05")
}