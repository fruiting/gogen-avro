// Code generated by github.com/fruiting/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     superhero.avsc
 */
package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/fruiting/gogen-avro/v10/compiler"
	"github.com/fruiting/gogen-avro/v10/vm"
	"github.com/fruiting/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Superhero struct {
	Id int32 `json:"id"`

	Affiliation_id int32 `json:"affiliation_id"`

	Name string `json:"name"`

	Life float32 `json:"life"`

	Energy float32 `json:"energy"`

	Powers []Superpower `json:"powers"`
}

const SuperheroAvroCRC64Fingerprint = "\x1e\x8e\x9cj\xdbGd\xaf"

func NewSuperhero() Superhero {
	r := Superhero{}
	r.Powers = make([]Superpower, 0)

	return r
}

func DeserializeSuperhero(r io.Reader) (Superhero, error) {
	t := NewSuperhero()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSuperheroFromSchema(r io.Reader, schema string) (Superhero, error) {
	t := NewSuperhero()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSuperhero(r Superhero, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Affiliation_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Life, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Energy, w)
	if err != nil {
		return err
	}
	err = writeArraySuperpower(r.Powers, w)
	if err != nil {
		return err
	}
	return err
}

func (r Superhero) Serialize(w io.Writer) error {
	return writeSuperhero(r, w)
}

func (r Superhero) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"affiliation_id\",\"type\":\"int\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"life\",\"type\":\"float\"},{\"name\":\"energy\",\"type\":\"float\"},{\"name\":\"powers\",\"type\":{\"items\":{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"damage\",\"type\":\"float\"},{\"name\":\"energy\",\"type\":\"float\"},{\"name\":\"passive\",\"type\":\"boolean\"}],\"name\":\"Superpower\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"com.model.Superhero\",\"type\":\"record\"}"
}

func (r Superhero) SchemaName() string {
	return "com.model.Superhero"
}

func (_ Superhero) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Superhero) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Superhero) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Superhero) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Superhero) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Superhero) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Superhero) SetString(v string)   { panic("Unsupported operation") }
func (_ Superhero) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Superhero) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.Int{Target: &r.Affiliation_id}

		return w

	case 2:
		w := types.String{Target: &r.Name}

		return w

	case 3:
		w := types.Float{Target: &r.Life}

		return w

	case 4:
		w := types.Float{Target: &r.Energy}

		return w

	case 5:
		r.Powers = make([]Superpower, 0)

		w := ArraySuperpowerWrapper{Target: &r.Powers}

		return w

	}
	panic("Unknown field index")
}

func (r *Superhero) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Superhero) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Superhero) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Superhero) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Superhero) HintSize(int)                     { panic("Unsupported operation") }
func (_ Superhero) Finalize()                        {}

func (_ Superhero) AvroCRC64Fingerprint() []byte {
	return []byte(SuperheroAvroCRC64Fingerprint)
}

func (r Superhero) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["affiliation_id"], err = json.Marshal(r.Affiliation_id)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["life"], err = json.Marshal(r.Life)
	if err != nil {
		return nil, err
	}
	output["energy"], err = json.Marshal(r.Energy)
	if err != nil {
		return nil, err
	}
	output["powers"], err = json.Marshal(r.Powers)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Superhero) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["affiliation_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Affiliation_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for affiliation_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["life"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Life); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for life")
	}
	val = func() json.RawMessage {
		if v, ok := fields["energy"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Energy); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for energy")
	}
	val = func() json.RawMessage {
		if v, ok := fields["powers"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Powers); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for powers")
	}
	return nil
}
