package corepubsubinf

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/coreinterface/loggerinf"
	"gitlab.com/auk-go/core/coreinterface/serializerinf"
)

type (
	SubscribeFunc func(
		subscriptionRecorder SubscriptionRecorder,
	)

	DirectSingleLogModelerSubscribeFunc func(modeler loggerinf.SingleLogModeler)

	JsonResultSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		jsonResult *corejson.Result,
	)

	BooleanSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		isResult bool,
	)

	HashmapSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		hashmap map[string]string,
	)
	BytesSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		rawBytes []byte,
	)
	ModelJsonSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		modelJson []byte,
	)
	AnyItemSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		anyItem interface{},
	)

	StringSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		message string,
	)

	SimpleBytesResulterSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		result serializerinf.SimpleBytesResulter,
	)

	BaseJsonResulterSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		result serializerinf.BaseJsonResulter,
	)

	JsonResulterSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		result serializerinf.JsonResulter,
	)

	DirectBytesSubscribeFunc func(
		rawBytes []byte,
	)

	DirectModelJsonSubscribeFunc func(
		modelJson []byte,
	)

	DirectAnyItemSubscribeFunc func(
		anyItem interface{},
	)

	DirectJsonResultSubscribeFunc func(
		jsonResult *corejson.Result,
	)

	DirectStringSubscribeFunc func(
		message string,
	)

	DirectBasicErrorSubscribeFunc func(
		basicErrorWrapper errcoreinf.BasicErrWrapper,
	)

	DirectBaseErrorOrCollectionWrapperSubscribeFunc func(
		basicErrorWrapper errcoreinf.BaseErrorOrCollectionWrapper,
	)

	DirectBooleanSubscribeFunc func(
		isResult bool,
	)

	DirectSimpleBytesResulterSubscribeFunc func(
		result serializerinf.SimpleBytesResulter,
	)

	DirectBaseJsonResulterSubscribeFunc func(
		result serializerinf.BaseJsonResulter,
	)

	DirectJsonResulterSubscribeFunc func(
		result serializerinf.JsonResulter,
	)

	SimpleSubscribeFunc                 func(communicate CommunicateModeler)
	FilterStringSubscribeFunc           func(communicate CommunicateModeler, currentStringValue string)
	LogSubscriberFunc                   func(logRecord BaseLogModeler)
	StartFunc                           func(subscriptionRecorder SubscriptionRecorder)
	CompletionFunc                      func(subscriptionRecorder SubscriptionRecorder)
	StartEndSubscriptionFunc            func(isStart bool, subscriptionRecorder SubscriptionRecorder)
	SimpleCompletionFunc                func(communicate CommunicateModeler)
	CategoryNameAnyItemSubscriptionFunc func(categoryName string, anyItem interface{})
	FilterAnyItemSubscriptionFunc       func(filter string, anyItem interface{})
	FilterBytesSubscriptionFunc         func(filter string, rawBytes []byte)
	MatcherFunc                         func() (isMatch bool)
)
