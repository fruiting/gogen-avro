// Code generated by github.com/alanctgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     nested.avsc
 */
package avro

import (
	"io"
)

type NestedTestRecord struct {
	NumberField *NumberRecord
	OtherField  *NestedRecord
}

func DeserializeNestedTestRecord(r io.Reader) (*NestedTestRecord, error) {
	return readNestedTestRecord(r)
}

func (r *NestedTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"NumberField\",\"type\":{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"DoubleField\",\"type\":\"double\"}],\"name\":\"NumberRecord\",\"type\":\"record\"}},{\"name\":\"OtherField\",\"type\":{\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"NestedRecord\",\"type\":\"record\"}}],\"name\":\"NestedTestRecord\",\"type\":\"record\"}"
}

func (r *NestedTestRecord) Serialize(w io.Writer) error {
	return writeNestedTestRecord(r, w)
}
