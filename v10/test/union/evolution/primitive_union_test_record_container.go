// Code generated by github.com/fruiting/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"io"

	"github.com/fruiting/gogen-avro/v10/compiler"
	"github.com/fruiting/gogen-avro/v10/container"
	"github.com/fruiting/gogen-avro/v10/vm"
)

func NewPrimitiveUnionTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewPrimitiveUnionTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type PrimitiveUnionTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewPrimitiveUnionTestRecordReader(r io.Reader) (*PrimitiveUnionTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewPrimitiveUnionTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &PrimitiveUnionTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r PrimitiveUnionTestRecordReader) Read() (PrimitiveUnionTestRecord, error) {
	t := NewPrimitiveUnionTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}