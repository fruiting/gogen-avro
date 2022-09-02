// Code generated by github.com/fruiting/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/fruiting/gogen-avro/v10/compiler"
	"github.com/fruiting/gogen-avro/v10/vm"
	"github.com/fruiting/gogen-avro/v10/vm/types"
)

type UnionNestedRecordNestedTestRecordTypeEnum int

const (
	UnionNestedRecordNestedTestRecordTypeEnumNestedRecord UnionNestedRecordNestedTestRecordTypeEnum = 0

	UnionNestedRecordNestedTestRecordTypeEnumNestedTestRecord UnionNestedRecordNestedTestRecordTypeEnum = 1
)

type UnionNestedRecordNestedTestRecord struct {
	NestedRecord     NestedRecord
	NestedTestRecord NestedTestRecord
	UnionType        UnionNestedRecordNestedTestRecordTypeEnum
}

func writeUnionNestedRecordNestedTestRecord(r UnionNestedRecordNestedTestRecord, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNestedRecordNestedTestRecordTypeEnumNestedRecord:
		return writeNestedRecord(r.NestedRecord, w)
	case UnionNestedRecordNestedTestRecordTypeEnumNestedTestRecord:
		return writeNestedTestRecord(r.NestedTestRecord, w)
	}
	return fmt.Errorf("invalid value for UnionNestedRecordNestedTestRecord")
}

func NewUnionNestedRecordNestedTestRecord() UnionNestedRecordNestedTestRecord {
	return UnionNestedRecordNestedTestRecord{}
}

func (r UnionNestedRecordNestedTestRecord) Serialize(w io.Writer) error {
	return writeUnionNestedRecordNestedTestRecord(r, w)
}

func DeserializeUnionNestedRecordNestedTestRecord(r io.Reader) (UnionNestedRecordNestedTestRecord, error) {
	t := NewUnionNestedRecordNestedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNestedRecordNestedTestRecordFromSchema(r io.Reader, schema string) (UnionNestedRecordNestedTestRecord, error) {
	t := NewUnionNestedRecordNestedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r UnionNestedRecordNestedTestRecord) Schema() string {
	return "[{\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"NestedRecord\",\"type\":\"record\"},{\"fields\":[{\"name\":\"OtherField\",\"type\":\"NestedRecord\"}],\"name\":\"NestedTestRecord\",\"type\":\"record\"}]"
}

func (_ UnionNestedRecordNestedTestRecord) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) SetInt(v int32)      { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNestedRecordNestedTestRecord) SetLong(v int64) {

	r.UnionType = (UnionNestedRecordNestedTestRecordTypeEnum)(v)
}

func (r *UnionNestedRecordNestedTestRecord) Get(i int) types.Field {

	switch i {
	case 0:
		r.NestedRecord = NewNestedRecord()
		return &types.Record{Target: (&r.NestedRecord)}
	case 1:
		r.NestedTestRecord = NewNestedTestRecord()
		return &types.Record{Target: (&r.NestedTestRecord)}
	}
	panic("Unknown field index")
}
func (_ UnionNestedRecordNestedTestRecord) NullField(i int)  { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) HintSize(i int)   { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) SetDefault(i int) { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ UnionNestedRecordNestedTestRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ UnionNestedRecordNestedTestRecord) Finalize()                {}

func (r UnionNestedRecordNestedTestRecord) MarshalJSON() ([]byte, error) {

	switch r.UnionType {
	case UnionNestedRecordNestedTestRecordTypeEnumNestedRecord:
		return json.Marshal(map[string]interface{}{"NestedRecord": r.NestedRecord})
	case UnionNestedRecordNestedTestRecordTypeEnumNestedTestRecord:
		return json.Marshal(map[string]interface{}{"NestedTestRecord": r.NestedTestRecord})
	}
	return nil, fmt.Errorf("invalid value for UnionNestedRecordNestedTestRecord")
}

func (r *UnionNestedRecordNestedTestRecord) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["NestedRecord"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.NestedRecord)
	}
	if value, ok := fields["NestedTestRecord"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.NestedTestRecord)
	}
	return fmt.Errorf("invalid value for UnionNestedRecordNestedTestRecord")
}
