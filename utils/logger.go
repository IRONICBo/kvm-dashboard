package utils

func LogWithInfo(args ...interface{}) {
	log.Info(args)
}

func LogWithError(args ...interface{}) {
	log.Error(args)
}

func LogWithWarn(args ...interface{}) {
	log.Warn(args)
}

func LogWithFatal(args ...interface{}) {
	log.Fatal(args)
}

func LogWithPanic(args ...interface{}) {
	log.Panic(args)
}
