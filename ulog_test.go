package ulog

import "testing"

func TestLog(t *testing.T) {
	SetWriteLog(func(loglevel int, msg string) {
		if !(loglevel == LogLevelDebug && msg == "debug: test 1") {
			t.Errorf("wrong" + msg)
			return
		}
	})
	Debugf("test %d", 1)
}
