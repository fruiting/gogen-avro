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

var _ = fmt.Printf

// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014
type DatatypeUUID struct {
	Uuid string `json:"uuid"`
}

const DatatypeUUIDAvroCRC64Fingerprint = "\xfc\xa43\x98\xee\xe0p\xe2"

func NewDatatypeUUID() DatatypeUUID {
	r := DatatypeUUID{}
	r.Uuid = ""
	return r
}

func DeserializeDatatypeUUID(r io.Reader) (DatatypeUUID, error) {
	t := NewDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatatypeUUIDFromSchema(r io.Reader, schema string) (DatatypeUUID, error) {
	t := NewDatatypeUUID()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatatypeUUID(r DatatypeUUID, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Uuid, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatatypeUUID) Serialize(w io.Writer) error {
	return writeDatatypeUUID(r, w)
}

func (r DatatypeUUID) Schema() string {
	return "{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"bodyworks.datatype.UUID\",\"type\":\"record\"}"
}

func (r DatatypeUUID) SchemaName() string {
	return "bodyworks.datatype.UUID"
}

func (_ DatatypeUUID) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatatypeUUID) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatatypeUUID) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatatypeUUID) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatatypeUUID) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatatypeUUID) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatatypeUUID) SetString(v string)   { panic("Unsupported operation") }
func (_ DatatypeUUID) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatatypeUUID) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Uuid}

		return w

	}
	panic("Unknown field index")
}

func (r *DatatypeUUID) SetDefault(i int) {
	switch i {
	case 0:
		r.Uuid = ""
		return
	}
	panic("Unknown field index")
}

func (r *DatatypeUUID) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatatypeUUID) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatatypeUUID) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatatypeUUID) Finalize()                        {}

func (_ DatatypeUUID) AvroCRC64Fingerprint() []byte {
	return []byte(DatatypeUUIDAvroCRC64Fingerprint)
}

func (r DatatypeUUID) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["uuid"], err = json.Marshal(r.Uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatatypeUUID) UnmarshalJSON(data []byte) error {
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
