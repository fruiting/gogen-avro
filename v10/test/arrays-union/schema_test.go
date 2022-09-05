package avro

import (
	"io"
	"testing"

	"github.com/fruiting/gogen-avro/v10/container"
	"github.com/fruiting/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTrip(t, func() container.AvroRecord { return &ArrayTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeArrayTestRecord(r)
			return &record, err
		})
}