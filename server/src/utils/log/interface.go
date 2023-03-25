package log

import (
	"io"
)

// Logger 定义Logger接口
type Logger interface {
	io.Writer

	Printer

	Debug(v ...interface{})
	DebugF(format string, v ...interface{})

	Info(v ...interface{})
	InfoF(format string, v ...interface{})

	Warn(v ...interface{})
	WarnF(format string, v ...interface{})

	Error(v ...interface{})
	ErrorF(format string, v ...interface{})

	Fatal(v ...interface{})
	FatalF(format string, v ...interface{})

	Panic(v ...interface{})
	PanicF(format string, v ...interface{})

	Print(v ...interface{})
	PrintF(format string, v ...interface{})

	DebugT(traceID string, v ...interface{})
	DebugTF(traceID string, format string, v ...interface{})

	InfoT(traceID string, v ...interface{})
	InfoTF(traceID string, format string, v ...interface{})

	WarnT(traceID string, v ...interface{})
	WarnTF(traceID string, format string, v ...interface{})

	ErrorT(traceID string, v ...interface{})
	ErrorTF(traceID string, format string, v ...interface{})

	FatalT(traceID string, v ...interface{})
	FatalTF(traceID string, format string, v ...interface{})

	PanicT(traceID string, v ...interface{})
	PanicTF(traceID string, format string, v ...interface{})

	PrintT(traceID string, v ...interface{})
	PrintTF(traceID string, format string, v ...interface{})
}

type Printer interface {
	// TPrintF 所有方法最终归为这个方法，真正打印日志
	TPrintF(v, l levelIndex, tag string, format string, m ...interface{})

	// TPrintTF 所有方法最终归为这个方法，真正打印日志, 带traceID
	TPrintTF(v, l levelIndex, tag string, traceID string, format string, m ...interface{})

	// ChangeFormat 改变日志格式
	ChangeFormat(format string)

	// ChangeWriter 改变输出流
	ChangeWriter(w io.Writer)
}
