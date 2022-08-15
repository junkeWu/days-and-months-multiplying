package log

import (
	"log"

	"gin-demo/errs"
)

// LogPanicData logs the panic data with stacktrace and return an
// error with the panic message. This function is separated from
// LogAndPanic so that unwanted panics can still be logged with
// this function.
func LogPanicData(panicData interface{}) error {
	err := errs.WrapPanicValue(panicData)
	log.Error("Flora panicked", "err", err)
	return err
}
