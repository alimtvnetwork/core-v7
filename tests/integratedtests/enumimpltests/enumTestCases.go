package enumimpltests

import (
	"gitlab.com/evatix-go/core/coredata/corerange"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl/enumtype"
)

var enumTestCases = []TestWrapper{
	{
		Header: "Byte enum example min 0, max 10",
		ExpectedMinMax: corerange.MinMaxInt64{
			Min: 0,
			Max: 10,
		},
		EnumMap: map[string]interface{}{
			"Invalid":   0,
			"A":         -2,
			"B":         8,
			"C":         5,
			"Something": 10,
		},
		EnumType: enumtype.Byte,
	},
	{
		Header: "Integer8 enum example min -2, max 12",
		ExpectedMinMax: corerange.MinMaxInt64{
			Min: -2,
			Max: 12,
		},
		EnumMap: map[string]interface{}{
			"Invalid":   -2,
			"A":         -2,
			"B":         8,
			"C":         5,
			"Something": 12,
		},
		EnumType: enumtype.Integer8,
	},
	{
		Header: "Integer16 enum example min -3, max 14",
		ExpectedMinMax: corerange.MinMaxInt64{
			Min: -3,
			Max: 14,
		},
		EnumMap: map[string]interface{}{
			"Invalid":   -3,
			"A":         -2,
			"B":         -3,
			"C":         5,
			"Something": 14,
		},
		EnumType: enumtype.Integer16,
	},
	{
		Header: "Integer32 enum example min -4, max 15",
		ExpectedMinMax: corerange.MinMaxInt64{
			Min: -4,

			Max: 15,
		},
		EnumMap: map[string]interface{}{
			"Invalid":   -4,
			"A":         -2,
			"B":         -3,
			"C":         5,
			"Something": 15,
		},
		EnumType: enumtype.Integer16,
	},

	{
		Header: "UnsignedInteger16 enum example min 0, max 20",
		ExpectedMinMax: corerange.MinMaxInt64{
			Min: 0,
			Max: 20,
		},
		EnumMap: map[string]interface{}{
			"Invalid":    0,
			"Something2": 15,
			"B":          15,
			"Something":  20,
		},
		EnumType: enumtype.UnsignedInteger16,
	},
	{
		Header: "String enum example min 0, max 20",
		ExpectedMinMax: corerange.MinMaxInt64{
			Min: 0,
			Max: 20,
		},
		EnumMap: map[string]interface{}{
			"Invalid":    0,
			"Something2": 15,
			"B":          15,
			"Something":  20,
		},
		EnumType: enumtype.String,
	},
}
