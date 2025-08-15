package log

type lockDistributorLogger struct {
	logger Logger
}

func (ldl *lockDistributorLogger) Println(v ...any) {
	// Log as debug to avoid to much log
	ldl.logger.Debug(v...)
}
