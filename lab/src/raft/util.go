package raft

import "log"

// Debugging
const Debug = 0

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug > 0 {
		log.Printf(format, a...)
	}
	return
}

// getLastLog returns the last log entry.
func getLastLog(log []LogEntry) *LogEntry {
	logSize := len(log)
	if (logSize == 0) {
		return nil
	}
	return &log[logSize - 1]
}
