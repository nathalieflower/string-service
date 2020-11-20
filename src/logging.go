package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type logmw struct {
	logger log.Logger
	StringService
}

/*Logging middleware for StringService*/
func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next StringService) StringService {
		return logmw{logger, next}
	}
}

func (mw logmw) Uppercase(s string) (output string, err error) {
	/*Setup a defer function to log some metrics about the request. Will complete after final call in the method*/
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.StringService.Uppercase(s)
	return
}

func (mw logmw) Lowercase(s string) (output string, err error) {
	/*Setup a defer function to log some metrics about the request. Will complete after final call in the method*/
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Lowercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.StringService.Lowercase(s)
	return
}

func (mw logmw) Count(s string) (n int) {
	/*Setup a defer function to log some metrics about the request. Will complete after final call in the method*/
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.StringService.Count(s)
	return
}
