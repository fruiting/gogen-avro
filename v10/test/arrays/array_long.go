// Code generated by github.com/fruiting/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/fruiting/gogen-avro/v10/vm"
	"github.com/fruiting/gogen-avro/v10/vm/types"
)

func writeArrayLong(r []int64, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteLong(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayLongWrapper struct {
	Target *[]int64
}

func (_ ArrayLongWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayLongWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayLongWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayLongWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayLongWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayLongWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayLongWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayLongWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayLongWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayLongWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayLongWrapper) Finalize()                        {}
func (_ ArrayLongWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayLongWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]int64, 0, s)
	}
}
func (r ArrayLongWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayLongWrapper) AppendArray() types.Field {
	var v int64

	*r.Target = append(*r.Target, v)
	return &types.Long{Target: &(*r.Target)[len(*r.Target)-1]}
}
