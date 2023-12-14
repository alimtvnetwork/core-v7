package errcore

import (
	"errors"
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

type stackTraceEnhance struct{}

func (it stackTraceEnhance) Error(err error) error {
	return it.ErrorSkip(1, err)
}

func (it stackTraceEnhance) ErrorSkip(skip int, err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) MsgToErrSkip(skip int, msg string) error {
	if len(msg) == 0 {
		return nil
	}

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) FmtSkip(skip int, format string, v ...interface{}) error {
	if len(format) == 0 {
		return nil
	}

	msg := fmt.Sprintf(format, v...)

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) Msg(msg string) string {
	if len(msg) == 0 {
		return ""
	}

	return it.MsgSkip(1, msg)
}

func (it stackTraceEnhance) MsgSkip(skip int, msg string) string {
	if len(msg) == 0 {
		return ""
	}

	if strings.Contains(msg, "Stack-Trace") {
		return msg
	}

	fullMessage := fmt.Sprintf(
		"%s - %s\n  - %s",
		reflectinternal.CodeStack.MethodName(1+skip),
		msg,
		reflectinternal.CodeStack.SingleStack(2+skip),
	)

	return fullMessage
}

func (it stackTraceEnhance) MsgErrorSkip(skip int, msg string, err error) string {
	if err == nil {
		return ""
	}

	compiledMsg := fmt.Sprintf(
		"%s %s",
		msg,
		err,
	)

	if strings.Contains(compiledMsg, "Stack-Trace") {
		return compiledMsg
	}

	fullMessage := fmt.Sprintf(
		"%s - %s\n  - %s",
		reflectinternal.CodeStack.MethodName(1+skip),
		compiledMsg,
		reflectinternal.CodeStack.SingleStack(2+skip),
	)

	return fullMessage
}

func (it stackTraceEnhance) MsgErrorToErrSkip(skip int, msg string, err error) error {
	if err == nil {
		return nil
	}

	return errors.New(it.MsgErrorSkip(1+skip, msg, err))
}
