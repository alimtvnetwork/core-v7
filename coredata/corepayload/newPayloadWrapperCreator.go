package corepayload

import (
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type newPayloadWrapperCreator struct{}

func (it newPayloadWrapperCreator) Empty() *PayloadWrapper {
	return &PayloadWrapper{
		Payloads:   []byte{},
		Attributes: New.Attributes.Empty(),
	}
}

func (it newPayloadWrapperCreator) Deserialize(
	rawBytes []byte,
) (*PayloadWrapper, error) {
	empty := it.Empty()

	err := corejson.
		Deserialize.UsingBytes(
		rawBytes, empty)

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (it newPayloadWrapperCreator) UsingCreateStruct(
	create *PayloadCreate,
) *PayloadWrapper {
	switch payloadCasted := create.Payloads.(type) {
	case []byte:
		return it.createInternalUsingBytes(
			create.Name,
			create.Identifier,
			create.TaskTypeName,
			create.CategoryName,
			create.EntityType,
			create.HasManyRecords,
			payloadCasted,
			create.Attributes,
			nil)
	case *[]byte:
		return it.createInternalUsingBytes(
			create.Name,
			create.Identifier,
			create.TaskTypeName,
			create.CategoryName,
			create.EntityType,
			create.HasManyRecords,
			converters.BytesPointerToBytes(payloadCasted),
			create.Attributes,
			nil)
	case string:
		return it.createInternalUsingBytes(
			create.Name,
			create.Identifier,
			create.TaskTypeName,
			create.CategoryName,
			create.EntityType,
			create.HasManyRecords,
			[]byte(payloadCasted),
			create.Attributes,
			nil)
	default:
		return it.createInternal(
			create.Name,
			create.Identifier,
			create.TaskTypeName,
			create.CategoryName,
			create.HasManyRecords,
			payloadCasted, // any
			create.Attributes,
		)
	}
}

func (it newPayloadWrapperCreator) UsingBytes(
	name, id, taskName,
	category, entityName string,
	payload []byte,
) *PayloadWrapper {
	return it.UsingCreateStruct(
		&PayloadCreate{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   entityName,
			CategoryName: category,
			Payloads:     payload,
		})
}

func (it newPayloadWrapperCreator) Create(
	name, id, taskName, category string,
	record interface{},
) *PayloadWrapper {
	return it.UsingCreateStruct(
		&PayloadCreate{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   reflectinternal.SafeTypeName(record),
			CategoryName: category,
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdCategory(
	name, id, category string,
	record interface{},
) *PayloadWrapper {
	entity := reflectinternal.SafeTypeName(
		record)

	return it.UsingCreateStruct(
		&PayloadCreate{
			Name:         name,
			Identifier:   id,
			TaskTypeName: entity,
			EntityType:   entity,
			CategoryName: category,
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) Records(
	name, id, taskName, category string,
	records interface{},
) *PayloadWrapper {
	return it.UsingCreateStruct(
		&PayloadCreate{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType: reflectinternal.SafeTypeNameOfSliceOrSingle(
				false, records),
			CategoryName:   category,
			HasManyRecords: true,
			Payloads:       records,
		})
}

func (it newPayloadWrapperCreator) Record(
	name, id, taskName, category string,
	record interface{},
) *PayloadWrapper {
	return it.UsingCreateStruct(
		&PayloadCreate{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType: reflectinternal.SafeTypeName(
				record),
			CategoryName: category,
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdTaskRecord(
	name, id, taskName string,
	record interface{},
) *PayloadWrapper {
	return it.UsingCreateStruct(
		&PayloadCreate{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   reflectinternal.SafeTypeName(record),
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdRecord(
	name, id string,
	record interface{},
) *PayloadWrapper {
	return it.UsingCreateStruct(
		&PayloadCreate{
			Name:       name,
			Identifier: id,
			EntityType: reflectinternal.SafeTypeName(record),
			Payloads:   record,
		})
}

func (it newPayloadWrapperCreator) NameTaskNameRecord(
	id, taskName string,
	record interface{},
) *PayloadWrapper {
	return it.UsingCreateStruct(
		&PayloadCreate{
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   reflectinternal.SafeTypeName(record),
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) createInternal(
	name, id, taskName, category string,
	hasManyRecords bool,
	records interface{},
	attr *Attributes,
) *PayloadWrapper {
	jsonResult := corejson.
		Serialize.
		UsingAny(records)

	entityTypeName := reflectinternal.SafeTypeNameOfSliceOrSingle(
		!hasManyRecords,
		records)

	if jsonResult.HasError() {
		attr = attr.AttachOrAppendErrorMessage(
			jsonResult.MeaningfulErrorMessage())

		return &PayloadWrapper{
			Name:           name,
			Identifier:     id,
			TaskTypeName:   taskName,
			EntityType:     entityTypeName,
			HasManyRecords: hasManyRecords,
			CategoryName:   category,
			Payloads:       jsonResult.SafeBytes(),
			Attributes:     attr,
		}
	}

	return &PayloadWrapper{
		Name:           name,
		Identifier:     id,
		TaskTypeName:   taskName,
		HasManyRecords: hasManyRecords,
		EntityType:     entityTypeName,
		CategoryName:   category,
		Payloads:       jsonResult.SafeBytes(),
		Attributes:     attr,
	}
}

func (it newPayloadWrapperCreator) createInternalUsingBytes(
	name, id, taskName,
	category, entityName string,
	hasManyRecords bool,
	payloads []byte,
	attr *Attributes,
	err error,
) *PayloadWrapper {
	attr = attr.AttachOrAppendErrorMessage(
		errcore.ToString(err))

	return &PayloadWrapper{
		Name:           name,
		Identifier:     id,
		TaskTypeName:   taskName,
		EntityType:     entityName,
		HasManyRecords: hasManyRecords,
		CategoryName:   category,
		Payloads:       payloads,
		Attributes:     attr,
	}
}

func (it newPayloadWrapperCreator) ManyRecords(
	name, id, taskName, category string,
	records interface{},
) *PayloadWrapper {
	jsonResult := corejson.
		Serialize.
		UsingAny(records)

	if jsonResult.HasError() {
		return &PayloadWrapper{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   reflectinternal.SafeTypeName(records),
			CategoryName: category,
			Payloads:     jsonResult.SafeBytes(),
			Attributes: &Attributes{
				ErrorMessage: jsonResult.MeaningfulErrorMessage(),
			},
		}
	}

	return &PayloadWrapper{
		Name:         name,
		Identifier:   id,
		TaskTypeName: taskName,
		EntityType:   reflectinternal.SafeTypeName(records),
		CategoryName: category,
		Payloads:     jsonResult.SafeBytes(),
	}
}

func (it newPayloadWrapperCreator) All(
	name, id, taskName,
	category, entityTypeName string,
	hasManyRecords bool,
	attributes *Attributes,
	payloads []byte,
) *PayloadWrapper {
	return &PayloadWrapper{
		Name:           name,
		Identifier:     id,
		TaskTypeName:   taskName,
		EntityType:     entityTypeName,
		CategoryName:   category,
		HasManyRecords: hasManyRecords,
		Payloads:       payloads,
		Attributes:     attributes,
	}
}
