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

type ComCompanyTeamSomeType struct {
	FieldOne ComCompanyTeamTypeOne `json:"fieldOne"`

	FieldTwo ComCompanySharedTypeTwo `json:"fieldTwo"`
}

const ComCompanyTeamSomeTypeAvroCRC64Fingerprint = "0\xc0\xc0\xb60N\xc0\xf6"

func NewComCompanyTeamSomeType() ComCompanyTeamSomeType {
	r := ComCompanyTeamSomeType{}
	r.FieldOne = NewComCompanyTeamTypeOne()

	r.FieldTwo = NewComCompanySharedTypeTwo()

	return r
}

func DeserializeComCompanyTeamSomeType(r io.Reader) (ComCompanyTeamSomeType, error) {
	t := NewComCompanyTeamSomeType()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeComCompanyTeamSomeTypeFromSchema(r io.Reader, schema string) (ComCompanyTeamSomeType, error) {
	t := NewComCompanyTeamSomeType()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeComCompanyTeamSomeType(r ComCompanyTeamSomeType, w io.Writer) error {
	var err error
	err = writeComCompanyTeamTypeOne(r.FieldOne, w)
	if err != nil {
		return err
	}
	err = writeComCompanySharedTypeTwo(r.FieldTwo, w)
	if err != nil {
		return err
	}
	return err
}

func (r ComCompanyTeamSomeType) Serialize(w io.Writer) error {
	return writeComCompanyTeamSomeType(r, w)
}

func (r ComCompanyTeamSomeType) Schema() string {
	return "{\"fields\":[{\"name\":\"fieldOne\",\"type\":{\"fields\":[{\"name\":\"type\",\"type\":{\"name\":\"SomeEnum\",\"symbols\":[\"A\",\"B\",\"C\"],\"type\":\"enum\"}}],\"name\":\"TypeOne\",\"namespace\":\"com.company.team\",\"type\":\"record\"}},{\"name\":\"fieldTwo\",\"type\":{\"fields\":[{\"name\":\"type\",\"type\":{\"name\":\"SomeEnum\",\"symbols\":[\"X\",\"Y\"],\"type\":\"enum\"}}],\"name\":\"TypeTwo\",\"namespace\":\"com.company.shared\",\"type\":\"record\"}}],\"name\":\"com.company.team.SomeType\",\"type\":\"record\"}"
}

func (r ComCompanyTeamSomeType) SchemaName() string {
	return "com.company.team.SomeType"
}

func (_ ComCompanyTeamSomeType) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) SetString(v string)   { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComCompanyTeamSomeType) Get(i int) types.Field {
	switch i {
	case 0:
		r.FieldOne = NewComCompanyTeamTypeOne()

		w := types.Record{Target: &r.FieldOne}

		return w

	case 1:
		r.FieldTwo = NewComCompanySharedTypeTwo()

		w := types.Record{Target: &r.FieldTwo}

		return w

	}
	panic("Unknown field index")
}

func (r *ComCompanyTeamSomeType) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ComCompanyTeamSomeType) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ComCompanyTeamSomeType) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) HintSize(int)                     { panic("Unsupported operation") }
func (_ ComCompanyTeamSomeType) Finalize()                        {}

func (_ ComCompanyTeamSomeType) AvroCRC64Fingerprint() []byte {
	return []byte(ComCompanyTeamSomeTypeAvroCRC64Fingerprint)
}

func (r ComCompanyTeamSomeType) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["fieldOne"], err = json.Marshal(r.FieldOne)
	if err != nil {
		return nil, err
	}
	output["fieldTwo"], err = json.Marshal(r.FieldTwo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ComCompanyTeamSomeType) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["fieldOne"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FieldOne); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fieldOne")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fieldTwo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FieldTwo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fieldTwo")
	}
	return nil
}
