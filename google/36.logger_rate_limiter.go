package google

type Logger struct {
	tumblingWindowEnd map[string]int
}

//func Constructor() Logger {
//	return Logger{
//		tumblingWindowEnd: make(map[string]int),
//	}
//}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	end, ok := l.tumblingWindowEnd[message]
	if ok {
		if timestamp >= end {
			l.tumblingWindowEnd[message] = timestamp + 10
			return true
		} else {
			return false
		}
	} else {
		l.tumblingWindowEnd[message] = timestamp + 10
		return true
	}
}
