package issetter

import (
	"gitlab.com/evatix-go/core/simplewrap"
)

var (
	values = []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}

	jsonValuesMap = map[string]Value{
		simplewrap.WithDoubleQuote("0"):        Uninitialized,
		simplewrap.WithDoubleQuote(""):         Uninitialized,
		simplewrap.WithDoubleQuote("-"):        Uninitialized,
		simplewrap.WithDoubleQuote("-1"):       Uninitialized,
		simplewrap.WithDoubleQuote("1"):        True,
		simplewrap.WithDoubleQuote("yes"):      True,
		simplewrap.WithDoubleQuote("Yes"):      True,
		simplewrap.WithDoubleQuote("true"):     True,
		simplewrap.WithDoubleQuote("True"):     True,
		simplewrap.WithDoubleQuote("no"):       False,
		simplewrap.WithDoubleQuote("No"):       False,
		simplewrap.WithDoubleQuote("Nop"):      False,
		simplewrap.WithDoubleQuote("None"):     False,
		simplewrap.WithDoubleQuote("false"):    False,
		simplewrap.WithDoubleQuote("False"):    False,
		simplewrap.WithDoubleQuote("set"):      Set,
		simplewrap.WithDoubleQuote("Set"):      Set,
		simplewrap.WithDoubleQuote("Unset"):    Unset,
		simplewrap.WithDoubleQuote("unset"):    Unset,
		simplewrap.WithDoubleQuote("*"):        Wildcard,
		simplewrap.WithDoubleQuote("%"):        Wildcard,
		simplewrap.WithDoubleQuote("Wildcard"): Wildcard,
		simplewrap.WithDoubleQuote("WildCard"): Wildcard,
		simplewrap.WithDoubleQuote("wildcard"): Wildcard, // all small
		"0":                                    Uninitialized,
		"":                                     Uninitialized,
		"-":                                    Uninitialized,
		"true":                                 True,
		"True":                                 True,
		"yes":                                  True,
		"Yes":                                  True,
		"y":                                    True,
		"Y":                                    True,
		"false":                                False,
		"False":                                False,
		"no":                                   False,
		"No":                                   False,
		"n":                                    False,
		"N":                                    False,
		"*":                                    Wildcard,
		"Wildcard":                             Wildcard,
		"wildcard":                             Wildcard,
		"%":                                    Wildcard,
		"set":                                  Set,
		"Set":                                  Set,
		"Unset":                                Unset,
		"unset":                                Unset,
	}

	valuesToJsonBytesMap = map[Value][]byte{
		Uninitialized: jsonBytes("Uninitialized"),
		True:          jsonBytes("True"),
		False:         jsonBytes("False"),
		Unset:         jsonBytes("Unset"),
		Set:           jsonBytes("Set"),
		Wildcard:      jsonBytes("Wildcard"),
	}

	undefinedMap = map[Value]bool{
		Uninitialized: true,
		Wildcard:      true,
	}

	falseMap = map[Value]bool{
		False: true,
		Unset: true,
	}

	trueMap = map[Value]bool{
		True: true,
		Set:  true,
	}

	valuesToNameMap = map[Value]string{
		Uninitialized: "Uninitialized",
		True:          "True",
		False:         "False",
		Unset:         "Unset",
		Set:           "Set",
		Wildcard:      "Wildcard",
	}

	lowerCaseYesNoNames = map[Value]string{
		Uninitialized: "-",
		True:          "yes",
		False:         "no",
		Set:           "yes",
		Unset:         "no",
		Wildcard:      "*",
	}

	yesNoNames = map[Value]string{
		Uninitialized: "-",
		True:          "Yes",
		False:         "No",
		Set:           "Yes",
		Unset:         "No",
		Wildcard:      "*",
	}

	trueFalseNames = map[Value]string{
		Uninitialized: "-",
		True:          "True",
		False:         "False",
		Set:           "True",
		Unset:         "False",
		Wildcard:      "*",
	}

	trueFalseLowerNames = map[Value]string{
		Uninitialized: "-",
		True:          "true",
		False:         "false",
		Set:           "true",
		Unset:         "false",
		Wildcard:      "*",
	}

	setUnsetLowerNames = map[Value]string{
		Uninitialized: "-",
		True:          "set",
		False:         "unset",
		Set:           "set",
		Unset:         "unset",
		Wildcard:      "*",
	}

	convSetUnsetToTrueFalseMap = map[Value]Value{
		Uninitialized: Uninitialized,
		True:          True,
		False:         False,
		Set:           True,
		Unset:         False,
		Wildcard:      Wildcard,
	}

	convTrueFalseToSetUnsetMap = map[Value]Value{
		Uninitialized: Uninitialized,
		True:          Set,
		False:         Unset,
		Set:           Set,
		Unset:         Unset,
		Wildcard:      Wildcard,
	}
)
