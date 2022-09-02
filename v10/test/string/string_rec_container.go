// Code generated by github.com/fruiting/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/fruiting/gogen-avro/v10/compiler"
	"github.com/fruiting/gogen-avro/v10/container"
	"github.com/fruiting/gogen-avro/v10/vm"
)

func NewStringRecWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewStringRec()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type StringRecReader struct {
	r io.Reader
	p *vm.Program
}

func NewStringRecReader(r io.Reader) (*StringRecReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewStringRec()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &StringRecReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r StringRecReader) Read() (StringRec, error) {
	t := NewStringRec()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
