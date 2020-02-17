// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     defaults.avsc
 */

package avro

import (
	"github.com/actgardner/gogen-avro/vm/types"
)

type UnionNullInt struct {
	Null      *types.NullVal
	Int       int32
	UnionType UnionNullIntTypeEnum
}

type UnionNullIntTypeEnum int

const (
	UnionNullIntTypeEnumNull UnionNullIntTypeEnum = 0
	UnionNullIntTypeEnumInt  UnionNullIntTypeEnum = 1
)

func NewUnionNullInt() *UnionNullInt {
	return &UnionNullInt{}
}

func (_ *UnionNullInt) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullInt) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullInt) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullInt) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullInt) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullInt) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullInt) SetLong(v int64) {
	r.UnionType = (UnionNullIntTypeEnum)(v)
}
func (r *UnionNullInt) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
		break
	case 1:
		return (*types.Int)(&r.Int)
		break

	}
	panic("Unknown field index")
}
func (_ *UnionNullInt) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullInt) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullInt) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullInt) Finalize()                        {}
