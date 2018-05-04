package golog

// These flags define which text to prefix to each log entry generated by the Logger.
const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         LogFlag                          = 1 << iota // the date: 2009/01/23
	Ltime                                                      // the time: 01:23:23
	Lmicroseconds                                              // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                                                  // full file name and line number: /a/b/c/d.go:23
	Lshortfile                                                 // final file name element and line number: d.go:23. overrides Llongfile
	Llevel                                                     // 日志级别
	Lname                                                      // 日志名
	LstdFlags     = Ldate | Ltime | Llevel | Lname             // initial values for the standard logger
)

type LogFlag int

func (self LogFlag) Contains(flags LogFlag) bool {

	return int(self)&int(flags) != 0
}
