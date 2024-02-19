// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package keyshare

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Params_2_list)(nil)

type _Params_2_list struct {
	list *[]string
}

func (x *_Params_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Params_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_Params_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Params_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Params_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Params at list field TrustedAddresses as it is not of Message kind"))
}

func (x *_Params_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Params_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_Params_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Params                               protoreflect.MessageDescriptor
	fd_Params_key_expiry                    protoreflect.FieldDescriptor
	fd_Params_trusted_addresses             protoreflect.FieldDescriptor
	fd_Params_slash_fraction_no_keyshare    protoreflect.FieldDescriptor
	fd_Params_slash_fraction_wrong_keyshare protoreflect.FieldDescriptor
	fd_Params_minimum_bonded                protoreflect.FieldDescriptor
	fd_Params_max_idled_block               protoreflect.FieldDescriptor
)

func init() {
	file_fairyring_keyshare_params_proto_init()
	md_Params = File_fairyring_keyshare_params_proto.Messages().ByName("Params")
	fd_Params_key_expiry = md_Params.Fields().ByName("key_expiry")
	fd_Params_trusted_addresses = md_Params.Fields().ByName("trusted_addresses")
	fd_Params_slash_fraction_no_keyshare = md_Params.Fields().ByName("slash_fraction_no_keyshare")
	fd_Params_slash_fraction_wrong_keyshare = md_Params.Fields().ByName("slash_fraction_wrong_keyshare")
	fd_Params_minimum_bonded = md_Params.Fields().ByName("minimum_bonded")
	fd_Params_max_idled_block = md_Params.Fields().ByName("max_idled_block")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_fairyring_keyshare_params_proto_msgTypes[0]
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
	if x.KeyExpiry != uint64(0) {
		value := protoreflect.ValueOfUint64(x.KeyExpiry)
		if !f(fd_Params_key_expiry, value) {
			return
		}
	}
	if len(x.TrustedAddresses) != 0 {
		value := protoreflect.ValueOfList(&_Params_2_list{list: &x.TrustedAddresses})
		if !f(fd_Params_trusted_addresses, value) {
			return
		}
	}
	if len(x.SlashFractionNoKeyshare) != 0 {
		value := protoreflect.ValueOfBytes(x.SlashFractionNoKeyshare)
		if !f(fd_Params_slash_fraction_no_keyshare, value) {
			return
		}
	}
	if len(x.SlashFractionWrongKeyshare) != 0 {
		value := protoreflect.ValueOfBytes(x.SlashFractionWrongKeyshare)
		if !f(fd_Params_slash_fraction_wrong_keyshare, value) {
			return
		}
	}
	if x.MinimumBonded != uint64(0) {
		value := protoreflect.ValueOfUint64(x.MinimumBonded)
		if !f(fd_Params_minimum_bonded, value) {
			return
		}
	}
	if x.MaxIdledBlock != uint64(0) {
		value := protoreflect.ValueOfUint64(x.MaxIdledBlock)
		if !f(fd_Params_max_idled_block, value) {
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
	case "fairyring.keyshare.Params.key_expiry":
		return x.KeyExpiry != uint64(0)
	case "fairyring.keyshare.Params.trusted_addresses":
		return len(x.TrustedAddresses) != 0
	case "fairyring.keyshare.Params.slash_fraction_no_keyshare":
		return len(x.SlashFractionNoKeyshare) != 0
	case "fairyring.keyshare.Params.slash_fraction_wrong_keyshare":
		return len(x.SlashFractionWrongKeyshare) != 0
	case "fairyring.keyshare.Params.minimum_bonded":
		return x.MinimumBonded != uint64(0)
	case "fairyring.keyshare.Params.max_idled_block":
		return x.MaxIdledBlock != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.Params"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.Params does not contain field %s", fd.FullName()))
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
	case "fairyring.keyshare.Params.key_expiry":
		x.KeyExpiry = uint64(0)
	case "fairyring.keyshare.Params.trusted_addresses":
		x.TrustedAddresses = nil
	case "fairyring.keyshare.Params.slash_fraction_no_keyshare":
		x.SlashFractionNoKeyshare = nil
	case "fairyring.keyshare.Params.slash_fraction_wrong_keyshare":
		x.SlashFractionWrongKeyshare = nil
	case "fairyring.keyshare.Params.minimum_bonded":
		x.MinimumBonded = uint64(0)
	case "fairyring.keyshare.Params.max_idled_block":
		x.MaxIdledBlock = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.Params"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.Params does not contain field %s", fd.FullName()))
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
	case "fairyring.keyshare.Params.key_expiry":
		value := x.KeyExpiry
		return protoreflect.ValueOfUint64(value)
	case "fairyring.keyshare.Params.trusted_addresses":
		if len(x.TrustedAddresses) == 0 {
			return protoreflect.ValueOfList(&_Params_2_list{})
		}
		listValue := &_Params_2_list{list: &x.TrustedAddresses}
		return protoreflect.ValueOfList(listValue)
	case "fairyring.keyshare.Params.slash_fraction_no_keyshare":
		value := x.SlashFractionNoKeyshare
		return protoreflect.ValueOfBytes(value)
	case "fairyring.keyshare.Params.slash_fraction_wrong_keyshare":
		value := x.SlashFractionWrongKeyshare
		return protoreflect.ValueOfBytes(value)
	case "fairyring.keyshare.Params.minimum_bonded":
		value := x.MinimumBonded
		return protoreflect.ValueOfUint64(value)
	case "fairyring.keyshare.Params.max_idled_block":
		value := x.MaxIdledBlock
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.Params"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.Params does not contain field %s", descriptor.FullName()))
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
	case "fairyring.keyshare.Params.key_expiry":
		x.KeyExpiry = value.Uint()
	case "fairyring.keyshare.Params.trusted_addresses":
		lv := value.List()
		clv := lv.(*_Params_2_list)
		x.TrustedAddresses = *clv.list
	case "fairyring.keyshare.Params.slash_fraction_no_keyshare":
		x.SlashFractionNoKeyshare = value.Bytes()
	case "fairyring.keyshare.Params.slash_fraction_wrong_keyshare":
		x.SlashFractionWrongKeyshare = value.Bytes()
	case "fairyring.keyshare.Params.minimum_bonded":
		x.MinimumBonded = value.Uint()
	case "fairyring.keyshare.Params.max_idled_block":
		x.MaxIdledBlock = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.Params"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.Params does not contain field %s", fd.FullName()))
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
	case "fairyring.keyshare.Params.trusted_addresses":
		if x.TrustedAddresses == nil {
			x.TrustedAddresses = []string{}
		}
		value := &_Params_2_list{list: &x.TrustedAddresses}
		return protoreflect.ValueOfList(value)
	case "fairyring.keyshare.Params.key_expiry":
		panic(fmt.Errorf("field key_expiry of message fairyring.keyshare.Params is not mutable"))
	case "fairyring.keyshare.Params.slash_fraction_no_keyshare":
		panic(fmt.Errorf("field slash_fraction_no_keyshare of message fairyring.keyshare.Params is not mutable"))
	case "fairyring.keyshare.Params.slash_fraction_wrong_keyshare":
		panic(fmt.Errorf("field slash_fraction_wrong_keyshare of message fairyring.keyshare.Params is not mutable"))
	case "fairyring.keyshare.Params.minimum_bonded":
		panic(fmt.Errorf("field minimum_bonded of message fairyring.keyshare.Params is not mutable"))
	case "fairyring.keyshare.Params.max_idled_block":
		panic(fmt.Errorf("field max_idled_block of message fairyring.keyshare.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.Params"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "fairyring.keyshare.Params.key_expiry":
		return protoreflect.ValueOfUint64(uint64(0))
	case "fairyring.keyshare.Params.trusted_addresses":
		list := []string{}
		return protoreflect.ValueOfList(&_Params_2_list{list: &list})
	case "fairyring.keyshare.Params.slash_fraction_no_keyshare":
		return protoreflect.ValueOfBytes(nil)
	case "fairyring.keyshare.Params.slash_fraction_wrong_keyshare":
		return protoreflect.ValueOfBytes(nil)
	case "fairyring.keyshare.Params.minimum_bonded":
		return protoreflect.ValueOfUint64(uint64(0))
	case "fairyring.keyshare.Params.max_idled_block":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: fairyring.keyshare.Params"))
		}
		panic(fmt.Errorf("message fairyring.keyshare.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in fairyring.keyshare.Params", d.FullName()))
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
		if x.KeyExpiry != 0 {
			n += 1 + runtime.Sov(uint64(x.KeyExpiry))
		}
		if len(x.TrustedAddresses) > 0 {
			for _, s := range x.TrustedAddresses {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.SlashFractionNoKeyshare)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SlashFractionWrongKeyshare)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MinimumBonded != 0 {
			n += 1 + runtime.Sov(uint64(x.MinimumBonded))
		}
		if x.MaxIdledBlock != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxIdledBlock))
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
		if x.MaxIdledBlock != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxIdledBlock))
			i--
			dAtA[i] = 0x30
		}
		if x.MinimumBonded != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MinimumBonded))
			i--
			dAtA[i] = 0x28
		}
		if len(x.SlashFractionWrongKeyshare) > 0 {
			i -= len(x.SlashFractionWrongKeyshare)
			copy(dAtA[i:], x.SlashFractionWrongKeyshare)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SlashFractionWrongKeyshare)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.SlashFractionNoKeyshare) > 0 {
			i -= len(x.SlashFractionNoKeyshare)
			copy(dAtA[i:], x.SlashFractionNoKeyshare)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SlashFractionNoKeyshare)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TrustedAddresses) > 0 {
			for iNdEx := len(x.TrustedAddresses) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.TrustedAddresses[iNdEx])
				copy(dAtA[i:], x.TrustedAddresses[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TrustedAddresses[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.KeyExpiry != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.KeyExpiry))
			i--
			dAtA[i] = 0x8
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
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field KeyExpiry", wireType)
				}
				x.KeyExpiry = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.KeyExpiry |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TrustedAddresses", wireType)
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
				x.TrustedAddresses = append(x.TrustedAddresses, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SlashFractionNoKeyshare", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SlashFractionNoKeyshare = append(x.SlashFractionNoKeyshare[:0], dAtA[iNdEx:postIndex]...)
				if x.SlashFractionNoKeyshare == nil {
					x.SlashFractionNoKeyshare = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SlashFractionWrongKeyshare", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SlashFractionWrongKeyshare = append(x.SlashFractionWrongKeyshare[:0], dAtA[iNdEx:postIndex]...)
				if x.SlashFractionWrongKeyshare == nil {
					x.SlashFractionWrongKeyshare = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinimumBonded", wireType)
				}
				x.MinimumBonded = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MinimumBonded |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxIdledBlock", wireType)
				}
				x.MaxIdledBlock = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxIdledBlock |= uint64(b&0x7F) << shift
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
// source: fairyring/keyshare/params.proto

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

	KeyExpiry                  uint64   `protobuf:"varint,1,opt,name=key_expiry,json=keyExpiry,proto3" json:"key_expiry,omitempty"`
	TrustedAddresses           []string `protobuf:"bytes,2,rep,name=trusted_addresses,json=trustedAddresses,proto3" json:"trusted_addresses,omitempty"`
	SlashFractionNoKeyshare    []byte   `protobuf:"bytes,3,opt,name=slash_fraction_no_keyshare,json=slashFractionNoKeyshare,proto3" json:"slash_fraction_no_keyshare,omitempty"`
	SlashFractionWrongKeyshare []byte   `protobuf:"bytes,4,opt,name=slash_fraction_wrong_keyshare,json=slashFractionWrongKeyshare,proto3" json:"slash_fraction_wrong_keyshare,omitempty"`
	MinimumBonded              uint64   `protobuf:"varint,5,opt,name=minimum_bonded,json=minimumBonded,proto3" json:"minimum_bonded,omitempty"`
	MaxIdledBlock              uint64   `protobuf:"varint,6,opt,name=max_idled_block,json=maxIdledBlock,proto3" json:"max_idled_block,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fairyring_keyshare_params_proto_msgTypes[0]
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
	return file_fairyring_keyshare_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetKeyExpiry() uint64 {
	if x != nil {
		return x.KeyExpiry
	}
	return 0
}

func (x *Params) GetTrustedAddresses() []string {
	if x != nil {
		return x.TrustedAddresses
	}
	return nil
}

func (x *Params) GetSlashFractionNoKeyshare() []byte {
	if x != nil {
		return x.SlashFractionNoKeyshare
	}
	return nil
}

func (x *Params) GetSlashFractionWrongKeyshare() []byte {
	if x != nil {
		return x.SlashFractionWrongKeyshare
	}
	return nil
}

func (x *Params) GetMinimumBonded() uint64 {
	if x != nil {
		return x.MinimumBonded
	}
	return 0
}

func (x *Params) GetMaxIdledBlock() uint64 {
	if x != nil {
		return x.MaxIdledBlock
	}
	return 0
}

var File_fairyring_keyshare_params_proto protoreflect.FileDescriptor

var file_fairyring_keyshare_params_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x65, 0x79,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03, 0x0a, 0x06,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x6b, 0x0a, 0x1a, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x66, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x17, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12,
	0x71, 0x0a, 0x1d, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x77, 0x72, 0x6f, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x1a, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x72, 0x6f, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x62, 0x6f,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x69,
	0x6d, 0x75, 0x6d, 0x42, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78,
	0x5f, 0x69, 0x64, 0x6c, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x42, 0xb3, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x42, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x65,
	0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x46, 0x4b, 0x58, 0xaa, 0x02, 0x12, 0x46,
	0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72,
	0x65, 0xca, 0x02, 0x12, 0x46, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x4b, 0x65,
	0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0xe2, 0x02, 0x1e, 0x46, 0x61, 0x69, 0x72, 0x79, 0x72, 0x69,
	0x6e, 0x67, 0x5c, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x46, 0x61, 0x69, 0x72, 0x79, 0x72,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x4b, 0x65, 0x79, 0x73, 0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fairyring_keyshare_params_proto_rawDescOnce sync.Once
	file_fairyring_keyshare_params_proto_rawDescData = file_fairyring_keyshare_params_proto_rawDesc
)

func file_fairyring_keyshare_params_proto_rawDescGZIP() []byte {
	file_fairyring_keyshare_params_proto_rawDescOnce.Do(func() {
		file_fairyring_keyshare_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_fairyring_keyshare_params_proto_rawDescData)
	})
	return file_fairyring_keyshare_params_proto_rawDescData
}

var file_fairyring_keyshare_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fairyring_keyshare_params_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: fairyring.keyshare.Params
}
var file_fairyring_keyshare_params_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fairyring_keyshare_params_proto_init() }
func file_fairyring_keyshare_params_proto_init() {
	if File_fairyring_keyshare_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fairyring_keyshare_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_fairyring_keyshare_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fairyring_keyshare_params_proto_goTypes,
		DependencyIndexes: file_fairyring_keyshare_params_proto_depIdxs,
		MessageInfos:      file_fairyring_keyshare_params_proto_msgTypes,
	}.Build()
	File_fairyring_keyshare_params_proto = out.File
	file_fairyring_keyshare_params_proto_rawDesc = nil
	file_fairyring_keyshare_params_proto_goTypes = nil
	file_fairyring_keyshare_params_proto_depIdxs = nil
}
