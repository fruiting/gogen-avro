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

type ArrayTestRecord struct {
	IntField []int32 `json:"IntField"`

	LongField []int64 `json:"LongField"`

	DoubleField []float64 `json:"DoubleField"`

	StringField []string `json:"StringField"`

	FloatField []float32 `json:"FloatField"`

	BoolField []bool `json:"BoolField"`

	BytesField []Bytes `json:"BytesField"`
}

const ArrayTestRecordAvroCRC64Fingerprint = "\"I\xbbnO\x1c#,"

func NewArrayTestRecord() ArrayTestRecord {
	r := ArrayTestRecord{}
	r.IntField = make([]int32, 0)

	r.IntField = make([]int32, 4)
	r.IntField[0] = 1
	r.IntField[1] = 2
	r.IntField[2] = 3
	r.IntField[3] = 4

	r.LongField = make([]int64, 0)

	r.LongField = make([]int64, 4)
	r.LongField[0] = 5
	r.LongField[1] = 6
	r.LongField[2] = 7
	r.LongField[3] = 8

	r.DoubleField = make([]float64, 0)

	r.DoubleField = make([]float64, 2)
	r.DoubleField[0] = 1.5
	r.DoubleField[1] = 2.4

	r.StringField = make([]string, 0)

	r.StringField = make([]string, 2)
	r.StringField[0] = "abc"
	r.StringField[1] = "def"

	r.FloatField = make([]float32, 0)

	r.FloatField = make([]float32, 2)
	r.FloatField[0] = 1.23
	r.FloatField[1] = 3.45

	r.BoolField = make([]bool, 0)

	r.BoolField = make([]bool, 2)
	r.BoolField[0] = true
	r.BoolField[1] = false

	r.BytesField = make([]Bytes, 0)

	r.BytesField = make([]Bytes, 2)
	r.BytesField[0] = []byte("abc")
	r.BytesField[1] = []byte("def")

	return r
}

func DeserializeArrayTestRecord(r io.Reader) (ArrayTestRecord, error) {
	t := NewArrayTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArrayTestRecordFromSchema(r io.Reader, schema string) (ArrayTestRecord, error) {
	t := NewArrayTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArrayTestRecord(r ArrayTestRecord, w io.Writer) error {
	var err error
	err = writeArrayInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = writeArrayLong(r.LongField, w)
	if err != nil {
		return err
	}
	err = writeArrayDouble(r.DoubleField, w)
	if err != nil {
		return err
	}
	err = writeArrayString(r.StringField, w)
	if err != nil {
		return err
	}
	err = writeArrayFloat(r.FloatField, w)
	if err != nil {
		return err
	}
	err = writeArrayBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = writeArrayBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r ArrayTestRecord) Serialize(w io.Writer) error {
	return writeArrayTestRecord(r, w)
}

func (r ArrayTestRecord) Schema() string {
	return "{\"fields\":[{\"default\":[1,2,3,4],\"name\":\"IntField\",\"type\":{\"items\":\"int\",\"type\":\"array\"}},{\"default\":[5,6,7,8],\"name\":\"LongField\",\"type\":{\"items\":\"long\",\"type\":\"array\"}},{\"default\":[1.5,2.4],\"name\":\"DoubleField\",\"type\":{\"items\":\"double\",\"type\":\"array\"}},{\"default\":[\"abc\",\"def\"],\"name\":\"StringField\",\"type\":{\"items\":\"string\",\"type\":\"array\"}},{\"default\":[1.23,3.45],\"name\":\"FloatField\",\"type\":{\"items\":\"float\",\"type\":\"array\"}},{\"default\":[true,false],\"name\":\"BoolField\",\"type\":{\"items\":\"boolean\",\"type\":\"array\"}},{\"default\":[\"abc\",\"def\"],\"name\":\"BytesField\",\"type\":{\"items\":\"bytes\",\"type\":\"array\"}}],\"name\":\"ArrayTestRecord\",\"type\":\"record\"}"
}

func (r ArrayTestRecord) SchemaName() string {
	return "ArrayTestRecord"
}

func (_ ArrayTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ArrayTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.IntField = make([]int32, 0)

		w := ArrayIntWrapper{Target: &r.IntField}

		return w

	case 1:
		r.LongField = make([]int64, 0)

		w := ArrayLongWrapper{Target: &r.LongField}

		return w

	case 2:
		r.DoubleField = make([]float64, 0)

		w := ArrayDoubleWrapper{Target: &r.DoubleField}

		return w

	case 3:
		r.StringField = make([]string, 0)

		w := ArrayStringWrapper{Target: &r.StringField}

		return w

	case 4:
		r.FloatField = make([]float32, 0)

		w := ArrayFloatWrapper{Target: &r.FloatField}

		return w

	case 5:
		r.BoolField = make([]bool, 0)

		w := ArrayBoolWrapper{Target: &r.BoolField}

		return w

	case 6:
		r.BytesField = make([]Bytes, 0)

		w := ArrayBytesWrapper{Target: &r.BytesField}

		return w

	}
	panic("Unknown field index")
}

func (r *ArrayTestRecord) SetDefault(i int) {
	switch i {
	case 0:
		r.IntField = make([]int32, 4)
		r.IntField[0] = 1
		r.IntField[1] = 2
		r.IntField[2] = 3
		r.IntField[3] = 4

		return
	case 1:
		r.LongField = make([]int64, 4)
		r.LongField[0] = 5
		r.LongField[1] = 6
		r.LongField[2] = 7
		r.LongField[3] = 8

		return
	case 2:
		r.DoubleField = make([]float64, 2)
		r.DoubleField[0] = 1.5
		r.DoubleField[1] = 2.4

		return
	case 3:
		r.StringField = make([]string, 2)
		r.StringField[0] = "abc"
		r.StringField[1] = "def"

		return
	case 4:
		r.FloatField = make([]float32, 2)
		r.FloatField[0] = 1.23
		r.FloatField[1] = 3.45

		return
	case 5:
		r.BoolField = make([]bool, 2)
		r.BoolField[0] = true
		r.BoolField[1] = false

		return
	case 6:
		r.BytesField = make([]Bytes, 2)
		r.BytesField[0] = []byte("abc")
		r.BytesField[1] = []byte("def")

		return
	}
	panic("Unknown field index")
}

func (r *ArrayTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ArrayTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ArrayTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ ArrayTestRecord) Finalize()                        {}

func (_ ArrayTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(ArrayTestRecordAvroCRC64Fingerprint)
}

func (r ArrayTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IntField"], err = json.Marshal(r.IntField)
	if err != nil {
		return nil, err
	}
	output["LongField"], err = json.Marshal(r.LongField)
	if err != nil {
		return nil, err
	}
	output["DoubleField"], err = json.Marshal(r.DoubleField)
	if err != nil {
		return nil, err
	}
	output["StringField"], err = json.Marshal(r.StringField)
	if err != nil {
		return nil, err
	}
	output["FloatField"], err = json.Marshal(r.FloatField)
	if err != nil {
		return nil, err
	}
	output["BoolField"], err = json.Marshal(r.BoolField)
	if err != nil {
		return nil, err
	}
	output["BytesField"], err = json.Marshal(r.BytesField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ArrayTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IntField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IntField); err != nil {
			return err
		}
	} else {
		r.IntField = make([]int32, 0)

		r.IntField = make([]int32, 4)
		r.IntField[0] = 1
		r.IntField[1] = 2
		r.IntField[2] = 3
		r.IntField[3] = 4

	}
	val = func() json.RawMessage {
		if v, ok := fields["LongField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LongField); err != nil {
			return err
		}
	} else {
		r.LongField = make([]int64, 0)

		r.LongField = make([]int64, 4)
		r.LongField[0] = 5
		r.LongField[1] = 6
		r.LongField[2] = 7
		r.LongField[3] = 8

	}
	val = func() json.RawMessage {
		if v, ok := fields["DoubleField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DoubleField); err != nil {
			return err
		}
	} else {
		r.DoubleField = make([]float64, 0)

		r.DoubleField = make([]float64, 2)
		r.DoubleField[0] = 1.5
		r.DoubleField[1] = 2.4

	}
	val = func() json.RawMessage {
		if v, ok := fields["StringField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StringField); err != nil {
			return err
		}
	} else {
		r.StringField = make([]string, 0)

		r.StringField = make([]string, 2)
		r.StringField[0] = "abc"
		r.StringField[1] = "def"

	}
	val = func() json.RawMessage {
		if v, ok := fields["FloatField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FloatField); err != nil {
			return err
		}
	} else {
		r.FloatField = make([]float32, 0)

		r.FloatField = make([]float32, 2)
		r.FloatField[0] = 1.23
		r.FloatField[1] = 3.45

	}
	val = func() json.RawMessage {
		if v, ok := fields["BoolField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BoolField); err != nil {
			return err
		}
	} else {
		r.BoolField = make([]bool, 0)

		r.BoolField = make([]bool, 2)
		r.BoolField[0] = true
		r.BoolField[1] = false

	}
	val = func() json.RawMessage {
		if v, ok := fields["BytesField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BytesField); err != nil {
			return err
		}
	} else {
		r.BytesField = make([]Bytes, 0)

		r.BytesField = make([]Bytes, 2)
		r.BytesField[0] = []byte("abc")
		r.BytesField[1] = []byte("def")

	}
	return nil
}
