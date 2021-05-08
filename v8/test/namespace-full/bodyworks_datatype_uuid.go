// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
)

var _ = fmt.Printf

// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014
type BodyworksDatatypeUUID struct {
	Uuid string `json:"uuid"`
}

const BodyworksDatatypeUUIDAvroCRC64Fingerprint = "\xfc\xa43\x98\xee\xe0p\xe2"

func NewBodyworksDatatypeUUID() BodyworksDatatypeUUID {
	r := BodyworksDatatypeUUID{}
	r.Uuid = ""
	return r
}

func DeserializeBodyworksDatatypeUUID(r io.Reader) (BodyworksDatatypeUUID, error) {
	t := NewBodyworksDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBodyworksDatatypeUUIDFromSchema(r io.Reader, schema string) (BodyworksDatatypeUUID, error) {
	t := NewBodyworksDatatypeUUID()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBodyworksDatatypeUUID(r BodyworksDatatypeUUID, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Uuid, w)
	if err != nil {
		return err
	}
	return err
}

func (r BodyworksDatatypeUUID) Serialize(w io.Writer) error {
	return writeBodyworksDatatypeUUID(r, w)
}

func (r BodyworksDatatypeUUID) Schema() string {
	return "{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"bodyworks.datatype.UUID\",\"type\":\"record\"}"
}

func (r BodyworksDatatypeUUID) SchemaName() string {
	return "bodyworks.datatype.UUID"
}

func (_ BodyworksDatatypeUUID) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) SetInt(v int32)       { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) SetLong(v int64)      { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) SetString(v string)   { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BodyworksDatatypeUUID) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Uuid}
	}
	panic("Unknown field index")
}

func (r *BodyworksDatatypeUUID) SetDefault(i int) {
	switch i {
	case 0:
		r.Uuid = ""
		return
	}
	panic("Unknown field index")
}

func (r *BodyworksDatatypeUUID) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ BodyworksDatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ BodyworksDatatypeUUID) Finalize()                        {}

func (_ BodyworksDatatypeUUID) AvroCRC64Fingerprint() []byte {
	return []byte(BodyworksDatatypeUUIDAvroCRC64Fingerprint)
}

func (r BodyworksDatatypeUUID) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["uuid"], err = json.Marshal(r.Uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *BodyworksDatatypeUUID) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["uuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Uuid); err != nil {
			return err
		}
	} else {
		r.Uuid = ""
	}
	return nil
}
