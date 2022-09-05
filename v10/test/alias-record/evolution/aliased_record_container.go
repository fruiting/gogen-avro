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

func NewAliasedRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewAliasedRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type AliasedRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewAliasedRecordReader(r io.Reader) (*AliasedRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewAliasedRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &AliasedRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r AliasedRecordReader) Read() (AliasedRecord, error) {
	t := NewAliasedRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}