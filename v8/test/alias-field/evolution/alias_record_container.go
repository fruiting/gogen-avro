// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/container"
	"github.com/actgardner/gogen-avro/v8/vm"
)

func NewAliasRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewAliasRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type AliasRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewAliasRecordReader(r io.Reader) (*AliasRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewAliasRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &AliasRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r AliasRecordReader) Read() (AliasRecord, error) {
	t := NewAliasRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}