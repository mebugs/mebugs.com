package logUtil

func (s *Standard) Debug(m ...interface{}) {
	s.TPrintF(globalLevel, DebugLevel, "", "", m...)
}
func (s *Standard) DebugF(format string, m ...interface{}) {
	s.TPrintF(globalLevel, DebugLevel, "", format, m...)
}
func (s *Standard) Info(m ...interface{}) {
	s.TPrintF(globalLevel, InfoLevel, "", "", m...)
}

func (s *Standard) InfoF(format string, m ...interface{}) {
	s.TPrintF(globalLevel, InfoLevel, "", format, m...)
}

func (s *Standard) Warn(m ...interface{}) {
	s.TPrintF(globalLevel, WarnLevel, "", "", m...)
}
func (s *Standard) WarnF(format string, m ...interface{}) {
	s.TPrintF(globalLevel, WarnLevel, "", format, m...)
}

func (s *Standard) Error(m ...interface{}) {
	s.TPrintF(globalLevel, ErrorLevel, "", "", m...)
}

func (s *Standard) Panic(m ...interface{}) {
	s.TPrintF(globalLevel, PanicLevel, "", "", m...)
}

func (s *Standard) Fatal(m ...interface{}) {
	s.TPrintF(globalLevel, FatalLevel, "", "", m...)
}

func (s *Standard) Print(m ...interface{}) {
	s.TPrintF(globalLevel, PrintLevel, "", "", m...)
}

// ---

func (s *Standard) ErrorF(format string, m ...interface{}) {
	s.TPrintF(globalLevel, ErrorLevel, "", format, m...)
}

func (s *Standard) PanicF(format string, m ...interface{}) {
	s.TPrintF(globalLevel, PanicLevel, "", format, m...)
}

func (s *Standard) FatalF(format string, m ...interface{}) {
	s.TPrintF(globalLevel, FatalLevel, "", format, m...)
}

func (s *Standard) PrintF(format string, m ...interface{}) {
	s.TPrintF(globalLevel, PrintLevel, "", format, m...)
}

func (s *Standard) DebugT(traceID string, m ...interface{}) {
	s.TPrintTF(globalLevel, DebugLevel, "", traceID, "", m...)
}

func (s *Standard) InfoT(traceID string, m ...interface{}) {
	s.TPrintTF(globalLevel, InfoLevel, "", traceID, "", m...)
}

func (s *Standard) WarnT(traceID string, m ...interface{}) {
	s.TPrintTF(globalLevel, WarnLevel, "", traceID, "", m...)
}

func (s *Standard) ErrorT(traceID string, m ...interface{}) {
	s.TPrintTF(globalLevel, ErrorLevel, "", traceID, "", m...)
}

func (s *Standard) PanicT(traceID string, m ...interface{}) {
	s.TPrintTF(globalLevel, PanicLevel, "", traceID, "", m...)
}

func (s *Standard) FatalT(traceID string, m ...interface{}) {
	s.TPrintTF(globalLevel, FatalLevel, "", traceID, "", m...)
}

func (s *Standard) PrintT(traceID string, m ...interface{}) {
	s.TPrintTF(globalLevel, PrintLevel, "", traceID, "", m...)
}

// ---

func (s *Standard) DebugTF(traceID string, format string, m ...interface{}) {
	s.TPrintTF(globalLevel, DebugLevel, "", traceID, format, m...)
}

func (s *Standard) InfoTF(traceID string, format string, m ...interface{}) {
	s.TPrintTF(globalLevel, InfoLevel, "", traceID, format, m...)
}

func (s *Standard) WarnTF(traceID string, format string, m ...interface{}) {
	s.TPrintTF(globalLevel, WarnLevel, "", traceID, format, m...)
}

func (s *Standard) ErrorTF(traceID string, format string, m ...interface{}) {
	s.TPrintTF(globalLevel, ErrorLevel, "", traceID, format, m...)
}

func (s *Standard) PanicTF(traceID string, format string, m ...interface{}) {
	s.TPrintTF(globalLevel, PanicLevel, "", traceID, format, m...)
}

func (s *Standard) FatalTF(traceID string, format string, m ...interface{}) {
	s.TPrintTF(globalLevel, FatalLevel, "", traceID, format, m...)
}

func (s *Standard) PrintTF(traceID string, format string, m ...interface{}) {
	s.TPrintTF(globalLevel, PrintLevel, "", traceID, format, m...)
}
