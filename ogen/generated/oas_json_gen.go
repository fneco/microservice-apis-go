// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode encodes CancelOrderNotFound as json.
func (s *CancelOrderNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes CancelOrderNotFound from json.
func (s *CancelOrderNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CancelOrderNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = CancelOrderNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CancelOrderNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CancelOrderNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes CancelOrderUnprocessableEntity as json.
func (s *CancelOrderUnprocessableEntity) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes CancelOrderUnprocessableEntity from json.
func (s *CancelOrderUnprocessableEntity) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CancelOrderUnprocessableEntity to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = CancelOrderUnprocessableEntity(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CancelOrderUnprocessableEntity) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CancelOrderUnprocessableEntity) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateOrderSchema) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateOrderSchema) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("order")
		e.ArrStart()
		for _, elem := range s.Order {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfCreateOrderSchema = [1]string{
	0: "order",
}

// Decode decodes CreateOrderSchema from json.
func (s *CreateOrderSchema) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateOrderSchema to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "order":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Order = make([]OrderItemSchema, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem OrderItemSchema
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Order = append(s.Order, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"order\"")
			}
		default:
			return errors.Errorf("unexpected field %q", k)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateOrderSchema")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfCreateOrderSchema) {
					name = jsonFieldsNameOfCreateOrderSchema[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CreateOrderSchema) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateOrderSchema) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes DeleteOrderNotFound as json.
func (s *DeleteOrderNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes DeleteOrderNotFound from json.
func (s *DeleteOrderNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DeleteOrderNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = DeleteOrderNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DeleteOrderNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DeleteOrderNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes DeleteOrderUnprocessableEntity as json.
func (s *DeleteOrderUnprocessableEntity) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes DeleteOrderUnprocessableEntity from json.
func (s *DeleteOrderUnprocessableEntity) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DeleteOrderUnprocessableEntity to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = DeleteOrderUnprocessableEntity(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DeleteOrderUnprocessableEntity) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DeleteOrderUnprocessableEntity) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("detail")
		s.Detail.Encode(e)
	}
}

var jsonFieldsNameOfError = [1]string{
	0: "detail",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "detail":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Detail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"detail\"")
			}
		default:
			return errors.Errorf("unexpected field %q", k)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfError) {
					name = jsonFieldsNameOfError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ErrorDetail as json.
func (s ErrorDetail) Encode(e *jx.Encoder) {
	switch s.Type {
	case StringErrorDetail:
		e.Str(s.String)
	case StringArrayErrorDetail:
		e.ArrStart()
		for _, elem := range s.StringArray {
			e.Str(elem)
		}
		e.ArrEnd()
	}
}

// Decode decodes ErrorDetail from json.
func (s *ErrorDetail) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorDetail to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Array:
		s.StringArray = make([]string, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem string
			v, err := d.Str()
			elem = string(v)
			if err != nil {
				return err
			}
			s.StringArray = append(s.StringArray, elem)
			return nil
		}); err != nil {
			return err
		}
		s.Type = StringArrayErrorDetail
	case jx.String:
		v, err := d.Str()
		s.String = string(v)
		if err != nil {
			return err
		}
		s.Type = StringErrorDetail
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ErrorDetail) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorDetail) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetOrderNotFound as json.
func (s *GetOrderNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetOrderNotFound from json.
func (s *GetOrderNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetOrderNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetOrderNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetOrderNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetOrderNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GetOrderSchema) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GetOrderSchema) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		json.EncodeUUID(e, s.ID)
	}
	{
		e.FieldStart("created")
		json.EncodeDateTime(e, s.Created)
	}
	{
		e.FieldStart("status")
		s.Status.Encode(e)
	}
	{
		e.FieldStart("order")
		e.ArrStart()
		for _, elem := range s.Order {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

var jsonFieldsNameOfGetOrderSchema = [4]string{
	0: "id",
	1: "created",
	2: "status",
	3: "order",
}

// Decode decodes GetOrderSchema from json.
func (s *GetOrderSchema) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetOrderSchema to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := json.DecodeUUID(d)
				s.ID = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "created":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.Created = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created\"")
			}
		case "status":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "order":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				s.Order = make([]OrderItemSchema, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem OrderItemSchema
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Order = append(s.Order, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"order\"")
			}
		default:
			return errors.Errorf("unexpected field %q", k)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GetOrderSchema")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfGetOrderSchema) {
					name = jsonFieldsNameOfGetOrderSchema[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetOrderSchema) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetOrderSchema) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetOrderSchemaStatus as json.
func (s GetOrderSchemaStatus) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes GetOrderSchemaStatus from json.
func (s *GetOrderSchemaStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetOrderSchemaStatus to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch GetOrderSchemaStatus(v) {
	case GetOrderSchemaStatusCreated:
		*s = GetOrderSchemaStatusCreated
	case GetOrderSchemaStatusPaid:
		*s = GetOrderSchemaStatusPaid
	case GetOrderSchemaStatusProgress:
		*s = GetOrderSchemaStatusProgress
	case GetOrderSchemaStatusCancelled:
		*s = GetOrderSchemaStatusCancelled
	case GetOrderSchemaStatusDispatched:
		*s = GetOrderSchemaStatusDispatched
	case GetOrderSchemaStatusDelivered:
		*s = GetOrderSchemaStatusDelivered
	default:
		*s = GetOrderSchemaStatus(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s GetOrderSchemaStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetOrderSchemaStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetOrderUnprocessableEntity as json.
func (s *GetOrderUnprocessableEntity) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetOrderUnprocessableEntity from json.
func (s *GetOrderUnprocessableEntity) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetOrderUnprocessableEntity to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetOrderUnprocessableEntity(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetOrderUnprocessableEntity) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetOrderUnprocessableEntity) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GetOrdersOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GetOrdersOK) encodeFields(e *jx.Encoder) {
	{
		if s.Orders != nil {
			e.FieldStart("orders")
			e.ArrStart()
			for _, elem := range s.Orders {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfGetOrdersOK = [1]string{
	0: "orders",
}

// Decode decodes GetOrdersOK from json.
func (s *GetOrdersOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetOrdersOK to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "orders":
			if err := func() error {
				s.Orders = make([]GetOrderSchema, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem GetOrderSchema
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Orders = append(s.Orders, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"orders\"")
			}
		default:
			return errors.Errorf("unexpected field %q", k)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GetOrdersOK")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetOrdersOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetOrdersOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int64 as json.
func (o OptInt64) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int64(int64(o.Value))
}

// Decode decodes int64 from json.
func (o *OptInt64) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt64 to nil")
	}
	o.Set = true
	v, err := d.Int64()
	if err != nil {
		return err
	}
	o.Value = int64(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt64) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt64) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *OrderItemSchema) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *OrderItemSchema) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("product")
		e.Str(s.Product)
	}
	{
		e.FieldStart("size")
		s.Size.Encode(e)
	}
	{
		if s.Quantity.Set {
			e.FieldStart("quantity")
			s.Quantity.Encode(e)
		}
	}
}

var jsonFieldsNameOfOrderItemSchema = [3]string{
	0: "product",
	1: "size",
	2: "quantity",
}

// Decode decodes OrderItemSchema from json.
func (s *OrderItemSchema) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode OrderItemSchema to nil")
	}
	var requiredBitSet [1]uint8
	s.setDefaults()

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "product":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Product = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"product\"")
			}
		case "size":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Size.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"size\"")
			}
		case "quantity":
			if err := func() error {
				s.Quantity.Reset()
				if err := s.Quantity.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"quantity\"")
			}
		default:
			return errors.Errorf("unexpected field %q", k)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode OrderItemSchema")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfOrderItemSchema) {
					name = jsonFieldsNameOfOrderItemSchema[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *OrderItemSchema) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OrderItemSchema) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes OrderItemSchemaSize as json.
func (s OrderItemSchemaSize) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes OrderItemSchemaSize from json.
func (s *OrderItemSchemaSize) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode OrderItemSchemaSize to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch OrderItemSchemaSize(v) {
	case OrderItemSchemaSizeSmall:
		*s = OrderItemSchemaSizeSmall
	case OrderItemSchemaSizeMedium:
		*s = OrderItemSchemaSizeMedium
	case OrderItemSchemaSizeBig:
		*s = OrderItemSchemaSizeBig
	default:
		*s = OrderItemSchemaSize(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OrderItemSchemaSize) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OrderItemSchemaSize) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PayOrderNotFound as json.
func (s *PayOrderNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes PayOrderNotFound from json.
func (s *PayOrderNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PayOrderNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = PayOrderNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PayOrderNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PayOrderNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes PayOrderUnprocessableEntity as json.
func (s *PayOrderUnprocessableEntity) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes PayOrderUnprocessableEntity from json.
func (s *PayOrderUnprocessableEntity) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode PayOrderUnprocessableEntity to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = PayOrderUnprocessableEntity(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *PayOrderUnprocessableEntity) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *PayOrderUnprocessableEntity) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UpdateOrderNotFound as json.
func (s *UpdateOrderNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes UpdateOrderNotFound from json.
func (s *UpdateOrderNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateOrderNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UpdateOrderNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateOrderNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateOrderNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UpdateOrderUnprocessableEntity as json.
func (s *UpdateOrderUnprocessableEntity) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes UpdateOrderUnprocessableEntity from json.
func (s *UpdateOrderUnprocessableEntity) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateOrderUnprocessableEntity to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = UpdateOrderUnprocessableEntity(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateOrderUnprocessableEntity) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateOrderUnprocessableEntity) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
