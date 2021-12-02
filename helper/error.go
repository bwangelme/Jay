package helper

import "qae/logger"

func EnsureOK(err error, fields map[string]interface{}) {
	if err != nil {
		fields["err"] = err
		logger.WithField(fields).Panicln()
	}
}
