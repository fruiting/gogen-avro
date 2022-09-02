package avro

import (
	"io"
	"testing"

	"github.com/fruiting/gogen-avro/v10/container"
	"github.com/fruiting/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTripExactBytes(t,
		func() container.AvroRecord { return &NestedTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeNestedTestRecord(r)
			return &record, err
		})
}
