// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/container"
	"github.com/actgardner/gogen-avro/v9/vm"
)

func NewArrayTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewArrayTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ArrayTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewArrayTestRecordReader(r io.Reader) (*ArrayTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewArrayTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ArrayTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ArrayTestRecordReader) Read() (ArrayTestRecord, error) {
	t := NewArrayTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}