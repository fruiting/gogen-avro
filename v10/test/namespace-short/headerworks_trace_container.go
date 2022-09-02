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

func NewHeaderworksTraceWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewHeaderworksTrace()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type HeaderworksTraceReader struct {
	r io.Reader
	p *vm.Program
}

func NewHeaderworksTraceReader(r io.Reader) (*HeaderworksTraceReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewHeaderworksTrace()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &HeaderworksTraceReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r HeaderworksTraceReader) Read() (HeaderworksTrace, error) {
	t := NewHeaderworksTrace()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
