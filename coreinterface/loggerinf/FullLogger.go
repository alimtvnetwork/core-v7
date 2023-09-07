package loggerinf

import "gitlab.com/auk-go/core/coreinterface/enuminf"

type FullLogger interface {
	LogModel(
		stackSkip int,
		logType LoggerTyper,
		title string,
		model SingleLogModeler,
	)

	LogModelNoTitle(
		stackSkip int,
		logType LoggerTyper,
		model SingleLogModeler,
	)

	LogString(
		stackSkip int,
		logType LoggerTyper,
		title string,
		attr string,
	)

	LogStringType(
		stackSkip int,
		logType LoggerTyper,
		filterType enuminf.BasicEnumer,
		levelType enuminf.LogLevelTyper,
		title string,
		attr string,
		model interface{},
	)

	LogAll(
		stackSkip int,
		logType LoggerTyper,
		message, attributes string,
	)

	LogAllUsingStackSkip(
		stackSkipIndex int,
		logType LoggerTyper,
		message, attributes string,
	)
	LogAllUsingConfig(
		config Configurer,
		message, attributes string,
	)
}
