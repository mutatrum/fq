package decoders

import (
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/scalar"
)

func decodeNullFn(sms ...scalar.Mapper) (DecodeFn, error) {
	// null is written as zero bytes.
	return func(name string, d *decode.D) interface{} {
		d.FieldValueNil(name, sms...)
		return nil
	}, nil
}
