// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/container"
	"github.com/actgardner/gogen-avro/v8/vm"
)

func NewNameConflictTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewNameConflictTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type NameConflictTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewNameConflictTestRecordReader(r io.Reader) (*NameConflictTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewNameConflictTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &NameConflictTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r NameConflictTestRecordReader) Read() (NameConflictTestRecord, error) {
	t := NewNameConflictTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}