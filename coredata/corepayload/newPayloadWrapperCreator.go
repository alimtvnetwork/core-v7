package corepayload

import (
	"fmt"

	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/isany"
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
		Deserialize.
		UsingBytes(
			rawBytes, empty)

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (it newPayloadWrapperCreator) CastOrDeserializeFrom(
	anyItem interface{},
) (*PayloadWrapper, error) {
	if isany.Null(anyItem) {
		return nil, errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs(
				"given any item is nil failed to convert to payload-wrapper")
	}

	toPayloadWrapper := &PayloadWrapper{}
	err := corejson.CastAny.FromToDefault(
		anyItem,
		toPayloadWrapper)

	return toPayloadWrapper, err
}

func (it newPayloadWrapperCreator) DeserializeToMany(
	rawBytes []byte,
) (payloadsSlice []*PayloadWrapper, err error) {
	err = corejson.
		Deserialize.
		UsingBytes(
			rawBytes,
			&payloadsSlice)

	if err != nil {
		return nil, err
	}

	return payloadsSlice, nil
}

func (it newPayloadWrapperCreator) DeserializeToCollection(
	rawBytes []byte,
) (payloadsSlice *PayloadsCollection, err error) {
	return New.
		PayloadsCollection.
		Deserialize(
			rawBytes)
}

func (it newPayloadWrapperCreator) DeserializeUsingJsonResult(
	jsonResult *corejson.Result,
) (*PayloadWrapper, error) {
	empty := it.Empty()

	err := corejson.
		Deserialize.
		Apply(jsonResult, empty)

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (it newPayloadWrapperCreator) UsingBytesCreateInstructionTypeStringer(
	createInstruction *BytesCreateInstructionStringer,
) *PayloadWrapper {
	return it.createInternalUsingBytes(
		createInstruction.Name,
		createInstruction.Identifier,
		createInstruction.TaskTypeName.String(),
		createInstruction.CategoryName.String(),
		createInstruction.EntityType,
		createInstruction.HasManyRecords,
		createInstruction.Payloads,
		createInstruction.Attributes,
		nil)
}

func (it newPayloadWrapperCreator) UsingBytesCreateInstruction(
	createInstruction *BytesCreateInstruction,
) *PayloadWrapper {
	return it.createInternalUsingBytes(
		createInstruction.Name,
		createInstruction.Identifier,
		createInstruction.TaskTypeName,
		createInstruction.CategoryName,
		createInstruction.EntityType,
		createInstruction.HasManyRecords,
		createInstruction.Payloads,
		createInstruction.Attributes,
		nil)
}

func (it newPayloadWrapperCreator) UsingCreateInstructionTypeStringer(
	createInstruction *PayloadCreateInstructionTypeStringer,
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(createInstruction.PayloadCreateInstruction())
}

func (it newPayloadWrapperCreator) UsingCreateInstruction(
	createInstruction *PayloadCreateInstruction,
) (*PayloadWrapper, error) {
	switch payloadCasted := createInstruction.Payloads.(type) {
	case []byte:
		return it.createInternalUsingBytes(
			createInstruction.Name,
			createInstruction.Identifier,
			createInstruction.TaskTypeName,
			createInstruction.CategoryName,
			createInstruction.EntityType,
			createInstruction.HasManyRecords,
			payloadCasted,
			createInstruction.Attributes,
			nil), nil
	case *[]byte:
		return it.createInternalUsingBytes(
			createInstruction.Name,
			createInstruction.Identifier,
			createInstruction.TaskTypeName,
			createInstruction.CategoryName,
			createInstruction.EntityType,
			createInstruction.HasManyRecords,
			converters.BytesPointerToBytes(payloadCasted),
			createInstruction.Attributes,
			nil), nil
	case string:
		return it.createInternalUsingBytes(
			createInstruction.Name,
			createInstruction.Identifier,
			createInstruction.TaskTypeName,
			createInstruction.CategoryName,
			createInstruction.EntityType,
			createInstruction.HasManyRecords,
			[]byte(payloadCasted),
			createInstruction.Attributes,
			nil), nil
	default: // any
		return it.createInternal(
			createInstruction.Name,
			createInstruction.Identifier,
			createInstruction.TaskTypeName,
			createInstruction.CategoryName,
			createInstruction.HasManyRecords,
			payloadCasted, // any
			createInstruction.Attributes,
		)
	}
}

func (it newPayloadWrapperCreator) UsingBytes(
	name, id, taskName,
	category, entityName string,
	payload []byte,
) *PayloadWrapper {
	payloadWrapper, err := it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   entityName,
			CategoryName: category,
			Payloads:     payload,
		})

	errcore.MustBeEmpty(err)

	return payloadWrapper
}

func (it newPayloadWrapperCreator) Create(
	name, id, taskName, category string,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   reflectinternal.SafeTypeName(record),
			CategoryName: category,
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) CreateUsingTypeStringer(
	name, id string,
	taskNameStringer, categoryStringer fmt.Stringer,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskNameStringer.String(),
			EntityType:   reflectinternal.SafeTypeName(record),
			CategoryName: categoryStringer.String(),
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdCategory(
	name, id, category string,
	record interface{},
) (*PayloadWrapper, error) {
	entity := reflectinternal.SafeTypeName(
		record)

	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: entity,
			EntityType:   entity,
			CategoryName: category,
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdCategoryStringer(
	name, id string,
	categoryStringer fmt.Stringer,
	record interface{},
) (*PayloadWrapper, error) {
	entity := reflectinternal.SafeTypeName(
		record)

	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: entity,
			EntityType:   entity,
			CategoryName: categoryStringer.String(),
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) Records(
	name, id, taskName, category string,
	records interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
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

func (it newPayloadWrapperCreator) RecordsTypeStringer(
	name, id string,
	taskNameStringer, categoryStringer fmt.Stringer,
	records interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskNameStringer.String(),
			EntityType: reflectinternal.SafeTypeNameOfSliceOrSingle(
				false, records),
			CategoryName:   categoryStringer.String(),
			HasManyRecords: true,
			Payloads:       records,
		})
}

func (it newPayloadWrapperCreator) Record(
	name, id, taskName, category string,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType: reflectinternal.SafeTypeName(
				record),
			CategoryName: category,
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) RecordTypeStringer(
	name, id string,
	taskNameStringer, categoryStringer fmt.Stringer,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskNameStringer.String(),
			EntityType: reflectinternal.SafeTypeName(
				record),
			CategoryName: categoryStringer.String(),
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdTaskRecord(
	name, id, taskName string,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskName,
			EntityType:   reflectinternal.SafeTypeName(record),
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdTaskStringerRecord(
	name, id string,
	taskNameStringer fmt.Stringer,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:         name,
			Identifier:   id,
			TaskTypeName: taskNameStringer.String(),
			EntityType:   reflectinternal.SafeTypeName(record),
			Payloads:     record,
		})
}

func (it newPayloadWrapperCreator) NameIdRecord(
	name, id string,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
			Name:       name,
			Identifier: id,
			EntityType: reflectinternal.SafeTypeName(record),
			Payloads:   record,
		})
}

func (it newPayloadWrapperCreator) NameTaskNameRecord(
	id, taskName string,
	record interface{},
) (*PayloadWrapper, error) {
	return it.UsingCreateInstruction(
		&PayloadCreateInstruction{
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
) (*PayloadWrapper, error) {
	jsonResult := corejson.
		Serialize.
		UsingAny(records)

	entityTypeName := reflectinternal.SafeTypeNameOfSliceOrSingle(
		!hasManyRecords,
		records)

	return &PayloadWrapper{
		Name:           name,
		Identifier:     id,
		TaskTypeName:   taskName,
		HasManyRecords: hasManyRecords,
		EntityType:     entityTypeName,
		CategoryName:   category,
		Payloads:       jsonResult.SafeBytes(),
		Attributes:     attr,
	}, jsonResult.MeaningfulError()
}

func (it newPayloadWrapperCreator) createInternalUsingBytes(
	name, id, taskName,
	category, entityName string,
	hasManyRecords bool,
	payloads []byte,
	attr *Attributes,
	basicErr errcoreinf.BasicErrWrapper, // will be mutated inside the attr error
) *PayloadWrapper {
	if attr == nil {
		attr = New.Attributes.UsingBasicError(basicErr)
	} else {
		attr.SetBasicErr(basicErr)
	}

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
) (*PayloadWrapper, error) {
	jsonResult := corejson.
		Serialize.
		UsingAny(records)

	return &PayloadWrapper{
		Name:         name,
		Identifier:   id,
		TaskTypeName: taskName,
		EntityType:   reflectinternal.SafeTypeName(records),
		CategoryName: category,
		Payloads:     jsonResult.SafeBytes(),
	}, jsonResult.MeaningfulError()
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

func (it newPayloadWrapperCreator) AllUsingStringer(
	name, id string,
	taskNameStringer,
	categoryStringer fmt.Stringer,
	entityTypeName string,
	hasManyRecords bool,
	attributes *Attributes,
	payloads []byte,
) *PayloadWrapper {
	return &PayloadWrapper{
		Name:           name,
		Identifier:     id,
		TaskTypeName:   taskNameStringer.String(),
		EntityType:     entityTypeName,
		CategoryName:   categoryStringer.String(),
		HasManyRecords: hasManyRecords,
		Payloads:       payloads,
		Attributes:     attributes,
	}
}

func (it newPayloadWrapperCreator) AllUsingExpander(
	name, id string,
	typeExpander PayloadTypeExpander,
	entityTypeName string,
	hasManyRecords bool,
	attributes *Attributes,
	payloads []byte,
) *PayloadWrapper {
	return &PayloadWrapper{
		Name:           name,
		Identifier:     id,
		TaskTypeName:   typeExpander.TaskTypeStringer.String(),
		EntityType:     entityTypeName,
		CategoryName:   typeExpander.CategoryStringer.String(),
		HasManyRecords: hasManyRecords,
		Payloads:       payloads,
		Attributes:     attributes,
	}
}
