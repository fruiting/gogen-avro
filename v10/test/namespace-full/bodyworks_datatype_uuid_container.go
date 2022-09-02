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

func NewBodyworksDatatypeUUIDWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewBodyworksDatatypeUUID()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type BodyworksDatatypeUUIDReader struct {
	r io.Reader
	p *vm.Program
}

func NewBodyworksDatatypeUUIDReader(r io.Reader) (*BodyworksDatatypeUUIDReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewBodyworksDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &BodyworksDatatypeUUIDReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r BodyworksDatatypeUUIDReader) Read() (BodyworksDatatypeUUID, error) {
	t := NewBodyworksDatatypeUUID()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
