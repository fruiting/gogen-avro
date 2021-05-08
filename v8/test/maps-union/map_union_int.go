// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
	"io"
)

func writeMapUnionInt(r map[string]*UnionInt, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeUnionInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapUnionIntWrapper struct {
	Target *map[string]*UnionInt
	keys   []string
	values []*UnionInt
}

func (_ *MapUnionIntWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapUnionIntWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapUnionIntWrapper) NullField(_ int) {
	r.values[len(r.values)-1] = nil
}

func (r *MapUnionIntWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapUnionIntWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v *UnionInt
	v = NewUnionInt()

	r.values = append(r.values, v)
	return r.values[len(r.values)-1]
}

func (_ *MapUnionIntWrapper) AppendArray() types.Field { panic("Unsupported operation") }