// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package blobgrantv1

import (
	_ "cosmossdk.io/api/amino"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Params                              protoreflect.MessageDescriptor
	fd_Params_gas_per_liquidity            protoreflect.FieldDescriptor
	fd_Params_expiry_duration              protoreflect.FieldDescriptor
	fd_Params_grant_token_refill_threshold protoreflect.FieldDescriptor
	fd_Params_block_height_duration        protoreflect.FieldDescriptor
)

func init() {
	file_sunrise_blobgrant_v1_params_proto_init()
	md_Params = File_sunrise_blobgrant_v1_params_proto.Messages().ByName("Params")
	fd_Params_gas_per_liquidity = md_Params.Fields().ByName("gas_per_liquidity")
	fd_Params_expiry_duration = md_Params.Fields().ByName("expiry_duration")
	fd_Params_grant_token_refill_threshold = md_Params.Fields().ByName("grant_token_refill_threshold")
	fd_Params_block_height_duration = md_Params.Fields().ByName("block_height_duration")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_sunrise_blobgrant_v1_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GasPerLiquidity != "" {
		value := protoreflect.ValueOfString(x.GasPerLiquidity)
		if !f(fd_Params_gas_per_liquidity, value) {
			return
		}
	}
	if x.ExpiryDuration != nil {
		value := protoreflect.ValueOfMessage(x.ExpiryDuration.ProtoReflect())
		if !f(fd_Params_expiry_duration, value) {
			return
		}
	}
	if x.GrantTokenRefillThreshold != "" {
		value := protoreflect.ValueOfString(x.GrantTokenRefillThreshold)
		if !f(fd_Params_grant_token_refill_threshold, value) {
			return
		}
	}
	if x.BlockHeightDuration != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BlockHeightDuration)
		if !f(fd_Params_block_height_duration, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "sunrise.blobgrant.v1.Params.gas_per_liquidity":
		return x.GasPerLiquidity != ""
	case "sunrise.blobgrant.v1.Params.expiry_duration":
		return x.ExpiryDuration != nil
	case "sunrise.blobgrant.v1.Params.grant_token_refill_threshold":
		return x.GrantTokenRefillThreshold != ""
	case "sunrise.blobgrant.v1.Params.block_height_duration":
		return x.BlockHeightDuration != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.blobgrant.v1.Params"))
		}
		panic(fmt.Errorf("message sunrise.blobgrant.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "sunrise.blobgrant.v1.Params.gas_per_liquidity":
		x.GasPerLiquidity = ""
	case "sunrise.blobgrant.v1.Params.expiry_duration":
		x.ExpiryDuration = nil
	case "sunrise.blobgrant.v1.Params.grant_token_refill_threshold":
		x.GrantTokenRefillThreshold = ""
	case "sunrise.blobgrant.v1.Params.block_height_duration":
		x.BlockHeightDuration = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.blobgrant.v1.Params"))
		}
		panic(fmt.Errorf("message sunrise.blobgrant.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "sunrise.blobgrant.v1.Params.gas_per_liquidity":
		value := x.GasPerLiquidity
		return protoreflect.ValueOfString(value)
	case "sunrise.blobgrant.v1.Params.expiry_duration":
		value := x.ExpiryDuration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "sunrise.blobgrant.v1.Params.grant_token_refill_threshold":
		value := x.GrantTokenRefillThreshold
		return protoreflect.ValueOfString(value)
	case "sunrise.blobgrant.v1.Params.block_height_duration":
		value := x.BlockHeightDuration
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.blobgrant.v1.Params"))
		}
		panic(fmt.Errorf("message sunrise.blobgrant.v1.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "sunrise.blobgrant.v1.Params.gas_per_liquidity":
		x.GasPerLiquidity = value.Interface().(string)
	case "sunrise.blobgrant.v1.Params.expiry_duration":
		x.ExpiryDuration = value.Message().Interface().(*durationpb.Duration)
	case "sunrise.blobgrant.v1.Params.grant_token_refill_threshold":
		x.GrantTokenRefillThreshold = value.Interface().(string)
	case "sunrise.blobgrant.v1.Params.block_height_duration":
		x.BlockHeightDuration = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.blobgrant.v1.Params"))
		}
		panic(fmt.Errorf("message sunrise.blobgrant.v1.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.blobgrant.v1.Params.expiry_duration":
		if x.ExpiryDuration == nil {
			x.ExpiryDuration = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.ExpiryDuration.ProtoReflect())
	case "sunrise.blobgrant.v1.Params.gas_per_liquidity":
		panic(fmt.Errorf("field gas_per_liquidity of message sunrise.blobgrant.v1.Params is not mutable"))
	case "sunrise.blobgrant.v1.Params.grant_token_refill_threshold":
		panic(fmt.Errorf("field grant_token_refill_threshold of message sunrise.blobgrant.v1.Params is not mutable"))
	case "sunrise.blobgrant.v1.Params.block_height_duration":
		panic(fmt.Errorf("field block_height_duration of message sunrise.blobgrant.v1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.blobgrant.v1.Params"))
		}
		panic(fmt.Errorf("message sunrise.blobgrant.v1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.blobgrant.v1.Params.gas_per_liquidity":
		return protoreflect.ValueOfString("")
	case "sunrise.blobgrant.v1.Params.expiry_duration":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "sunrise.blobgrant.v1.Params.grant_token_refill_threshold":
		return protoreflect.ValueOfString("")
	case "sunrise.blobgrant.v1.Params.block_height_duration":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.blobgrant.v1.Params"))
		}
		panic(fmt.Errorf("message sunrise.blobgrant.v1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in sunrise.blobgrant.v1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.GasPerLiquidity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ExpiryDuration != nil {
			l = options.Size(x.ExpiryDuration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.GrantTokenRefillThreshold)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BlockHeightDuration != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockHeightDuration))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.BlockHeightDuration != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockHeightDuration))
			i--
			dAtA[i] = 0x20
		}
		if len(x.GrantTokenRefillThreshold) > 0 {
			i -= len(x.GrantTokenRefillThreshold)
			copy(dAtA[i:], x.GrantTokenRefillThreshold)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.GrantTokenRefillThreshold)))
			i--
			dAtA[i] = 0x1a
		}
		if x.ExpiryDuration != nil {
			encoded, err := options.Marshal(x.ExpiryDuration)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.GasPerLiquidity) > 0 {
			i -= len(x.GasPerLiquidity)
			copy(dAtA[i:], x.GasPerLiquidity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.GasPerLiquidity)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasPerLiquidity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.GasPerLiquidity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ExpiryDuration", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ExpiryDuration == nil {
					x.ExpiryDuration = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ExpiryDuration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GrantTokenRefillThreshold", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.GrantTokenRefillThreshold = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockHeightDuration", wireType)
				}
				x.BlockHeightDuration = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockHeightDuration |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sunrise/blobgrant/v1/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the parameters for the module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasPerLiquidity           string               `protobuf:"bytes,1,opt,name=gas_per_liquidity,json=gasPerLiquidity,proto3" json:"gas_per_liquidity,omitempty"`
	ExpiryDuration            *durationpb.Duration `protobuf:"bytes,2,opt,name=expiry_duration,json=expiryDuration,proto3" json:"expiry_duration,omitempty"`
	GrantTokenRefillThreshold string               `protobuf:"bytes,3,opt,name=grant_token_refill_threshold,json=grantTokenRefillThreshold,proto3" json:"grant_token_refill_threshold,omitempty"`
	BlockHeightDuration       uint64               `protobuf:"varint,4,opt,name=block_height_duration,json=blockHeightDuration,proto3" json:"block_height_duration,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sunrise_blobgrant_v1_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_sunrise_blobgrant_v1_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetGasPerLiquidity() string {
	if x != nil {
		return x.GasPerLiquidity
	}
	return ""
}

func (x *Params) GetExpiryDuration() *durationpb.Duration {
	if x != nil {
		return x.ExpiryDuration
	}
	return nil
}

func (x *Params) GetGrantTokenRefillThreshold() string {
	if x != nil {
		return x.GrantTokenRefillThreshold
	}
	return ""
}

func (x *Params) GetBlockHeightDuration() uint64 {
	if x != nil {
		return x.BlockHeightDuration
	}
	return 0
}

var File_sunrise_blobgrant_v1_params_proto protoreflect.FileDescriptor

var file_sunrise_blobgrant_v1_params_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x62, 0x6c, 0x6f,
	0x62, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f,
	0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x03,
	0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x68, 0x0a, 0x11, 0x67, 0x61, 0x73, 0x5f,
	0x70, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x1b, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x4c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d, 0x10, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x44, 0x65, 0x63, 0xa8, 0xe7, 0xb0, 0x2a,
	0x01, 0x52, 0x0f, 0x67, 0x61, 0x73, 0x50, 0x65, 0x72, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x12, 0x51, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf, 0x1f, 0x01,
	0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x1c, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xc8, 0xde, 0x1f,
	0x00, 0xda, 0xde, 0x1f, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69,
	0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x19, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x26, 0xe8, 0xa0,
	0x1f, 0x01, 0x8a, 0xe7, 0xb0, 0x2a, 0x1d, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x78,
	0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x42, 0xcc, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x6e,
	0x72, 0x69, 0x73, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x6c, 0x6f, 0x62, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x58, 0xaa, 0x02, 0x14, 0x53, 0x75, 0x6e, 0x72,
	0x69, 0x73, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x14, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x5c, 0x42, 0x6c, 0x6f, 0x62, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x20, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73,
	0x65, 0x5c, 0x42, 0x6c, 0x6f, 0x62, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x53, 0x75, 0x6e,
	0x72, 0x69, 0x73, 0x65, 0x3a, 0x3a, 0x42, 0x6c, 0x6f, 0x62, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sunrise_blobgrant_v1_params_proto_rawDescOnce sync.Once
	file_sunrise_blobgrant_v1_params_proto_rawDescData = file_sunrise_blobgrant_v1_params_proto_rawDesc
)

func file_sunrise_blobgrant_v1_params_proto_rawDescGZIP() []byte {
	file_sunrise_blobgrant_v1_params_proto_rawDescOnce.Do(func() {
		file_sunrise_blobgrant_v1_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_sunrise_blobgrant_v1_params_proto_rawDescData)
	})
	return file_sunrise_blobgrant_v1_params_proto_rawDescData
}

var file_sunrise_blobgrant_v1_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sunrise_blobgrant_v1_params_proto_goTypes = []interface{}{
	(*Params)(nil),              // 0: sunrise.blobgrant.v1.Params
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_sunrise_blobgrant_v1_params_proto_depIdxs = []int32{
	1, // 0: sunrise.blobgrant.v1.Params.expiry_duration:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sunrise_blobgrant_v1_params_proto_init() }
func file_sunrise_blobgrant_v1_params_proto_init() {
	if File_sunrise_blobgrant_v1_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sunrise_blobgrant_v1_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sunrise_blobgrant_v1_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sunrise_blobgrant_v1_params_proto_goTypes,
		DependencyIndexes: file_sunrise_blobgrant_v1_params_proto_depIdxs,
		MessageInfos:      file_sunrise_blobgrant_v1_params_proto_msgTypes,
	}.Build()
	File_sunrise_blobgrant_v1_params_proto = out.File
	file_sunrise_blobgrant_v1_params_proto_rawDesc = nil
	file_sunrise_blobgrant_v1_params_proto_goTypes = nil
	file_sunrise_blobgrant_v1_params_proto_depIdxs = nil
}
