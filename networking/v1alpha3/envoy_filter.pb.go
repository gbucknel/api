// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/envoy_filter.proto

package v1alpha3

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/struct"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EnvoyFilters_ListenerMatch_ListenerType int32

const (
	// All listeners
	EnvoyFilters_ListenerMatch_ANY EnvoyFilters_ListenerMatch_ListenerType = 0
	// Inbound listener in sidecar
	EnvoyFilters_ListenerMatch_SIDECAR_INBOUND EnvoyFilters_ListenerMatch_ListenerType = 1
	// Outbound listener in sidecar
	EnvoyFilters_ListenerMatch_SIDECAR_OUTBOUND EnvoyFilters_ListenerMatch_ListenerType = 2
	// Gateway listener
	EnvoyFilters_ListenerMatch_GATEWAY EnvoyFilters_ListenerMatch_ListenerType = 3
)

var EnvoyFilters_ListenerMatch_ListenerType_name = map[int32]string{
	0: "ANY",
	1: "SIDECAR_INBOUND",
	2: "SIDECAR_OUTBOUND",
	3: "GATEWAY",
}
var EnvoyFilters_ListenerMatch_ListenerType_value = map[string]int32{
	"ANY":              0,
	"SIDECAR_INBOUND":  1,
	"SIDECAR_OUTBOUND": 2,
	"GATEWAY":          3,
}

func (x EnvoyFilters_ListenerMatch_ListenerType) String() string {
	return proto.EnumName(EnvoyFilters_ListenerMatch_ListenerType_name, int32(x))
}
func (EnvoyFilters_ListenerMatch_ListenerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEnvoyFilter, []int{0, 1, 0}
}

// Index/position in the filter chain.
type EnvoyFilters_InsertPosition_Index int32

const (
	// Insert first
	EnvoyFilters_InsertPosition_FIRST EnvoyFilters_InsertPosition_Index = 0
	// Insert last
	EnvoyFilters_InsertPosition_LAST EnvoyFilters_InsertPosition_Index = 1
	// Insert before the named filter.
	EnvoyFilters_InsertPosition_BEFORE EnvoyFilters_InsertPosition_Index = 2
	// Insert after the named filter.
	EnvoyFilters_InsertPosition_AFTER EnvoyFilters_InsertPosition_Index = 3
)

var EnvoyFilters_InsertPosition_Index_name = map[int32]string{
	0: "FIRST",
	1: "LAST",
	2: "BEFORE",
	3: "AFTER",
}
var EnvoyFilters_InsertPosition_Index_value = map[string]int32{
	"FIRST":  0,
	"LAST":   1,
	"BEFORE": 2,
	"AFTER":  3,
}

func (x EnvoyFilters_InsertPosition_Index) String() string {
	return proto.EnumName(EnvoyFilters_InsertPosition_Index_name, int32(x))
}
func (EnvoyFilters_InsertPosition_Index) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEnvoyFilter, []int{0, 2, 0}
}

type EnvoyFilters_Filter_FilterType int32

const (
	// placeholder
	EnvoyFilters_Filter_INVALID EnvoyFilters_Filter_FilterType = 0
	// Http filter
	EnvoyFilters_Filter_HTTP EnvoyFilters_Filter_FilterType = 1
	// Network filter
	EnvoyFilters_Filter_NETWORK EnvoyFilters_Filter_FilterType = 2
)

var EnvoyFilters_Filter_FilterType_name = map[int32]string{
	0: "INVALID",
	1: "HTTP",
	2: "NETWORK",
}
var EnvoyFilters_Filter_FilterType_value = map[string]int32{
	"INVALID": 0,
	"HTTP":    1,
	"NETWORK": 2,
}

func (x EnvoyFilters_Filter_FilterType) String() string {
	return proto.EnumName(EnvoyFilters_Filter_FilterType_name, int32(x))
}
func (EnvoyFilters_Filter_FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorEnvoyFilter, []int{0, 3, 0}
}

// `EnvoyFilter` describes Envoy proxy-specific filters that can be used to
// customize the Envoy proxy configuration generated by Istio networking
// subsystem (Pilot). This feature must be used with care, as incorrect
// configurations could potentially destabilize the entire mesh.
//
// NOTE: Since this is break glass configuration, there will not be any
// backward compatibility across different Istio releases. In other words,
// this configuration is subject to change based on internal implementation
// of Istio networking subsystem.
//
// The following example for Kubernetes enables Envoy's Lua filter for all
// inbound calls arriving at port 18080 of the reviews service pod with
// labels "app: reviews".
//
//     apiVersion: networking.istio.io/v1alpha3
//     kind: EnvoyFilter
//     metadata:
//       name: reviews-lua
//     spec:
//       selector:
//         app: reviews
//       filters:
//       - listenerMatch:
//           port: 18080
//           listenerType: SIDECAR_INBOUND #will match with the listener for the podIP:18080
//         filterName: envoy.lua
//         filterType: HTTP
//         filterConfig:
//           inlineCode: |
//             ... lua code ...
//
type EnvoyFilters struct {
	// One or more labels that indicate a specific set of pods/VMs whose
	// proxies should be configured to use these additional filters.  The
	// scope of label search is platform dependent. On Kubernetes, for
	// example, the scope includes pods running in all reachable
	// namespaces. Omitting the selector applies the filter to all proxies in
	// the mesh.
	Selector map[string]string `protobuf:"bytes,1,rep,name=selector" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// REQUIRED: Envoy network filters/http filters to be added to matching
	// listeners.  When adding network filters to http connections, care
	// should be taken to ensure that the filter is added before
	// envoy.http_connection_manager.
	Filters []*EnvoyFilters_Filter `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
}

func (m *EnvoyFilters) Reset()                    { *m = EnvoyFilters{} }
func (m *EnvoyFilters) String() string            { return proto.CompactTextString(m) }
func (*EnvoyFilters) ProtoMessage()               {}
func (*EnvoyFilters) Descriptor() ([]byte, []int) { return fileDescriptorEnvoyFilter, []int{0} }

func (m *EnvoyFilters) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *EnvoyFilters) GetFilters() []*EnvoyFilters_Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// Select a listener to add the filter to based on the match conditions.
// All conditions specified in the ListenerMatch must be met for the filter
// to be applied to a listener.
type EnvoyFilters_ListenerMatch struct {
	// Port associated with the listener. If not specified, matches all listeners.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Inbound vs outbound sidecar listener or gateway listener. If not specified,
	// matches all listeners.
	ListenerType EnvoyFilters_ListenerMatch_ListenerType `protobuf:"varint,2,opt,name=listener_type,json=listenerType,proto3,enum=istio.networking.v1alpha3.EnvoyFilters_ListenerMatch_ListenerType" json:"listener_type,omitempty"`
	// Selects a class of listeners for the same protocol. If not
	// specified, applies to listeners on all protocols. Use the protocol
	// selection to select all HTTP listeners (includes HTTP2/gRPC/HTTPS
	// where Envoy terminates TLS) or all TCP listeners (includes HTTPS
	// passthrough using SNI). Acceptable values are TCP/HTTP/MONGO.
	ListenerProtocol string `protobuf:"bytes,3,opt,name=listener_protocol,json=listenerProtocol,proto3" json:"listener_protocol,omitempty"`
	// One or more IP addresses to which the listener is bound. If
	// specified, should match atleast one address in the list.
	Address []string `protobuf:"bytes,4,rep,name=address" json:"address,omitempty"`
}

func (m *EnvoyFilters_ListenerMatch) Reset()         { *m = EnvoyFilters_ListenerMatch{} }
func (m *EnvoyFilters_ListenerMatch) String() string { return proto.CompactTextString(m) }
func (*EnvoyFilters_ListenerMatch) ProtoMessage()    {}
func (*EnvoyFilters_ListenerMatch) Descriptor() ([]byte, []int) {
	return fileDescriptorEnvoyFilter, []int{0, 1}
}

func (m *EnvoyFilters_ListenerMatch) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EnvoyFilters_ListenerMatch) GetListenerType() EnvoyFilters_ListenerMatch_ListenerType {
	if m != nil {
		return m.ListenerType
	}
	return EnvoyFilters_ListenerMatch_ANY
}

func (m *EnvoyFilters_ListenerMatch) GetListenerProtocol() string {
	if m != nil {
		return m.ListenerProtocol
	}
	return ""
}

func (m *EnvoyFilters_ListenerMatch) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

// Indicates the relative index in the filter chain where the filter should be inserted.
type EnvoyFilters_InsertPosition struct {
	// Position of this filter in the filter chain.
	Index EnvoyFilters_InsertPosition_Index `protobuf:"varint,1,opt,name=index,proto3,enum=istio.networking.v1alpha3.EnvoyFilters_InsertPosition_Index" json:"index,omitempty"`
	// If BEFORE or AFTER position is specified, specify the name of the
	// filter relative to which this filter should be inserted.
	RelativeTo string `protobuf:"bytes,2,opt,name=relative_to,json=relativeTo,proto3" json:"relative_to,omitempty"`
}

func (m *EnvoyFilters_InsertPosition) Reset()         { *m = EnvoyFilters_InsertPosition{} }
func (m *EnvoyFilters_InsertPosition) String() string { return proto.CompactTextString(m) }
func (*EnvoyFilters_InsertPosition) ProtoMessage()    {}
func (*EnvoyFilters_InsertPosition) Descriptor() ([]byte, []int) {
	return fileDescriptorEnvoyFilter, []int{0, 2}
}

func (m *EnvoyFilters_InsertPosition) GetIndex() EnvoyFilters_InsertPosition_Index {
	if m != nil {
		return m.Index
	}
	return EnvoyFilters_InsertPosition_FIRST
}

func (m *EnvoyFilters_InsertPosition) GetRelativeTo() string {
	if m != nil {
		return m.RelativeTo
	}
	return ""
}

// Envoy filters to be added to a network or http filter chain.
type EnvoyFilters_Filter struct {
	// Filter will be added to the listner only if the match conditions are true.
	// If not specified, the filters will be applied to all listeners.
	ListenerMatch *EnvoyFilters_ListenerMatch `protobuf:"bytes,1,opt,name=listener_match,json=listenerMatch" json:"listener_match,omitempty"`
	// Insert position in the filter chain. Defaults to FIRST
	InsertPosition *EnvoyFilters_InsertPosition `protobuf:"bytes,2,opt,name=insert_position,json=insertPosition" json:"insert_position,omitempty"`
	// REQUIRED: The type of filter to instantiate.
	FilterType EnvoyFilters_Filter_FilterType `protobuf:"varint,3,opt,name=filter_type,json=filterType,proto3,enum=istio.networking.v1alpha3.EnvoyFilters_Filter_FilterType" json:"filter_type,omitempty"`
	// REQUIRED: The name of the filter to instantiate. The name must match a supported
	// filter _compiled into_ Envoy.
	FilterName string `protobuf:"bytes,4,opt,name=filter_name,json=filterName,proto3" json:"filter_name,omitempty"`
	// REQUIRED: Filter specific configuration which depends on the filter being
	// instantiated.
	FilterConfig *google_protobuf2.Struct `protobuf:"bytes,5,opt,name=filter_config,json=filterConfig" json:"filter_config,omitempty"`
}

func (m *EnvoyFilters_Filter) Reset()         { *m = EnvoyFilters_Filter{} }
func (m *EnvoyFilters_Filter) String() string { return proto.CompactTextString(m) }
func (*EnvoyFilters_Filter) ProtoMessage()    {}
func (*EnvoyFilters_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptorEnvoyFilter, []int{0, 3}
}

func (m *EnvoyFilters_Filter) GetListenerMatch() *EnvoyFilters_ListenerMatch {
	if m != nil {
		return m.ListenerMatch
	}
	return nil
}

func (m *EnvoyFilters_Filter) GetInsertPosition() *EnvoyFilters_InsertPosition {
	if m != nil {
		return m.InsertPosition
	}
	return nil
}

func (m *EnvoyFilters_Filter) GetFilterType() EnvoyFilters_Filter_FilterType {
	if m != nil {
		return m.FilterType
	}
	return EnvoyFilters_Filter_INVALID
}

func (m *EnvoyFilters_Filter) GetFilterName() string {
	if m != nil {
		return m.FilterName
	}
	return ""
}

func (m *EnvoyFilters_Filter) GetFilterConfig() *google_protobuf2.Struct {
	if m != nil {
		return m.FilterConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*EnvoyFilters)(nil), "istio.networking.v1alpha3.EnvoyFilters")
	proto.RegisterType((*EnvoyFilters_ListenerMatch)(nil), "istio.networking.v1alpha3.EnvoyFilters.ListenerMatch")
	proto.RegisterType((*EnvoyFilters_InsertPosition)(nil), "istio.networking.v1alpha3.EnvoyFilters.InsertPosition")
	proto.RegisterType((*EnvoyFilters_Filter)(nil), "istio.networking.v1alpha3.EnvoyFilters.Filter")
	proto.RegisterEnum("istio.networking.v1alpha3.EnvoyFilters_ListenerMatch_ListenerType", EnvoyFilters_ListenerMatch_ListenerType_name, EnvoyFilters_ListenerMatch_ListenerType_value)
	proto.RegisterEnum("istio.networking.v1alpha3.EnvoyFilters_InsertPosition_Index", EnvoyFilters_InsertPosition_Index_name, EnvoyFilters_InsertPosition_Index_value)
	proto.RegisterEnum("istio.networking.v1alpha3.EnvoyFilters_Filter_FilterType", EnvoyFilters_Filter_FilterType_name, EnvoyFilters_Filter_FilterType_value)
}
func (m *EnvoyFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnvoyFilters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Selector) > 0 {
		for k, _ := range m.Selector {
			dAtA[i] = 0xa
			i++
			v := m.Selector[k]
			mapSize := 1 + len(k) + sovEnvoyFilter(uint64(len(k))) + 1 + len(v) + sovEnvoyFilter(uint64(len(v)))
			i = encodeVarintEnvoyFilter(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintEnvoyFilter(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintEnvoyFilter(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Filters) > 0 {
		for _, msg := range m.Filters {
			dAtA[i] = 0x12
			i++
			i = encodeVarintEnvoyFilter(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EnvoyFilters_ListenerMatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnvoyFilters_ListenerMatch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Port != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(m.Port))
	}
	if m.ListenerType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(m.ListenerType))
	}
	if len(m.ListenerProtocol) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(len(m.ListenerProtocol)))
		i += copy(dAtA[i:], m.ListenerProtocol)
	}
	if len(m.Address) > 0 {
		for _, s := range m.Address {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *EnvoyFilters_InsertPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnvoyFilters_InsertPosition) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(m.Index))
	}
	if len(m.RelativeTo) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(len(m.RelativeTo)))
		i += copy(dAtA[i:], m.RelativeTo)
	}
	return i, nil
}

func (m *EnvoyFilters_Filter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnvoyFilters_Filter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ListenerMatch != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(m.ListenerMatch.Size()))
		n1, err := m.ListenerMatch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.InsertPosition != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(m.InsertPosition.Size()))
		n2, err := m.InsertPosition.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.FilterType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(m.FilterType))
	}
	if len(m.FilterName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(len(m.FilterName)))
		i += copy(dAtA[i:], m.FilterName)
	}
	if m.FilterConfig != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEnvoyFilter(dAtA, i, uint64(m.FilterConfig.Size()))
		n3, err := m.FilterConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintEnvoyFilter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnvoyFilters) Size() (n int) {
	var l int
	_ = l
	if len(m.Selector) > 0 {
		for k, v := range m.Selector {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovEnvoyFilter(uint64(len(k))) + 1 + len(v) + sovEnvoyFilter(uint64(len(v)))
			n += mapEntrySize + 1 + sovEnvoyFilter(uint64(mapEntrySize))
		}
	}
	if len(m.Filters) > 0 {
		for _, e := range m.Filters {
			l = e.Size()
			n += 1 + l + sovEnvoyFilter(uint64(l))
		}
	}
	return n
}

func (m *EnvoyFilters_ListenerMatch) Size() (n int) {
	var l int
	_ = l
	if m.Port != 0 {
		n += 1 + sovEnvoyFilter(uint64(m.Port))
	}
	if m.ListenerType != 0 {
		n += 1 + sovEnvoyFilter(uint64(m.ListenerType))
	}
	l = len(m.ListenerProtocol)
	if l > 0 {
		n += 1 + l + sovEnvoyFilter(uint64(l))
	}
	if len(m.Address) > 0 {
		for _, s := range m.Address {
			l = len(s)
			n += 1 + l + sovEnvoyFilter(uint64(l))
		}
	}
	return n
}

func (m *EnvoyFilters_InsertPosition) Size() (n int) {
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovEnvoyFilter(uint64(m.Index))
	}
	l = len(m.RelativeTo)
	if l > 0 {
		n += 1 + l + sovEnvoyFilter(uint64(l))
	}
	return n
}

func (m *EnvoyFilters_Filter) Size() (n int) {
	var l int
	_ = l
	if m.ListenerMatch != nil {
		l = m.ListenerMatch.Size()
		n += 1 + l + sovEnvoyFilter(uint64(l))
	}
	if m.InsertPosition != nil {
		l = m.InsertPosition.Size()
		n += 1 + l + sovEnvoyFilter(uint64(l))
	}
	if m.FilterType != 0 {
		n += 1 + sovEnvoyFilter(uint64(m.FilterType))
	}
	l = len(m.FilterName)
	if l > 0 {
		n += 1 + l + sovEnvoyFilter(uint64(l))
	}
	if m.FilterConfig != nil {
		l = m.FilterConfig.Size()
		n += 1 + l + sovEnvoyFilter(uint64(l))
	}
	return n
}

func sovEnvoyFilter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEnvoyFilter(x uint64) (n int) {
	return sovEnvoyFilter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnvoyFilters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvoyFilter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EnvoyFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnvoyFilters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEnvoyFilter
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEnvoyFilter
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthEnvoyFilter
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEnvoyFilter
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthEnvoyFilter
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipEnvoyFilter(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthEnvoyFilter
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Selector[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, &EnvoyFilters_Filter{})
			if err := m.Filters[len(m.Filters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvoyFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EnvoyFilters_ListenerMatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvoyFilter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListenerMatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenerMatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenerType", wireType)
			}
			m.ListenerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ListenerType |= (EnvoyFilters_ListenerMatch_ListenerType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenerProtocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenerProtocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvoyFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EnvoyFilters_InsertPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvoyFilter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InsertPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InsertPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (EnvoyFilters_InsertPosition_Index(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelativeTo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelativeTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvoyFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EnvoyFilters_Filter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvoyFilter
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Filter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Filter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenerMatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ListenerMatch == nil {
				m.ListenerMatch = &EnvoyFilters_ListenerMatch{}
			}
			if err := m.ListenerMatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertPosition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InsertPosition == nil {
				m.InsertPosition = &EnvoyFilters_InsertPosition{}
			}
			if err := m.InsertPosition.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterType", wireType)
			}
			m.FilterType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FilterType |= (EnvoyFilters_Filter_FilterType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FilterConfig == nil {
				m.FilterConfig = &google_protobuf2.Struct{}
			}
			if err := m.FilterConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvoyFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvoyFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEnvoyFilter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnvoyFilter
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnvoyFilter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEnvoyFilter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEnvoyFilter
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEnvoyFilter(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEnvoyFilter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnvoyFilter   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("networking/v1alpha3/envoy_filter.proto", fileDescriptorEnvoyFilter) }

var fileDescriptorEnvoyFilter = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0xed, 0xa4, 0x69, 0x26, 0x3f, 0xf5, 0xb7, 0x5f, 0x25, 0x4c, 0x84, 0x4a, 0x94, 0x0b,
	0x54, 0x09, 0xc9, 0x81, 0x54, 0x20, 0x7e, 0x7a, 0xe3, 0xb4, 0x0e, 0xb5, 0x28, 0x49, 0xd9, 0xb8,
	0x54, 0xad, 0x90, 0x22, 0x37, 0xdd, 0xa4, 0xab, 0xba, 0x5e, 0xcb, 0xde, 0x06, 0x72, 0xcd, 0xab,
	0xf0, 0x08, 0x5c, 0xf1, 0x04, 0x5c, 0xf2, 0x08, 0xa8, 0x4f, 0x82, 0xbc, 0x6b, 0xbb, 0x8d, 0x04,
	0x52, 0xcb, 0x95, 0x77, 0x26, 0x67, 0xce, 0xcc, 0x9c, 0x99, 0x0c, 0x3c, 0x0a, 0x08, 0xff, 0xc4,
	0xa2, 0x73, 0x1a, 0x4c, 0xdb, 0xb3, 0xa7, 0x9e, 0x1f, 0x9e, 0x79, 0x9b, 0x6d, 0x12, 0xcc, 0xd8,
	0x7c, 0x34, 0xa1, 0x3e, 0x27, 0x91, 0x19, 0x46, 0x8c, 0x33, 0x74, 0x9f, 0xc6, 0x9c, 0x32, 0xf3,
	0x1a, 0x6d, 0x66, 0xe8, 0xc6, 0x83, 0x29, 0x63, 0x53, 0x9f, 0xb4, 0x05, 0xf0, 0xe4, 0x72, 0xd2,
	0x8e, 0x79, 0x74, 0x39, 0xe6, 0x32, 0xb0, 0xf5, 0xa5, 0x0c, 0x55, 0x3b, 0xe1, 0xeb, 0x09, 0xba,
	0x18, 0xbd, 0x87, 0x95, 0x98, 0xf8, 0x64, 0xcc, 0x59, 0x64, 0x28, 0x4d, 0x6d, 0xa3, 0xd2, 0x79,
	0x66, 0xfe, 0x95, 0xdc, 0xbc, 0x19, 0x6a, 0x0e, 0xd3, 0x38, 0x3b, 0xe0, 0xd1, 0x1c, 0xe7, 0x34,
	0x68, 0x17, 0x4a, 0xb2, 0xd8, 0xd8, 0x50, 0x05, 0xa3, 0x79, 0x5b, 0x46, 0xf9, 0xc5, 0x59, 0x78,
	0xe3, 0x35, 0xd4, 0x16, 0x92, 0x20, 0x1d, 0xb4, 0x73, 0x32, 0x37, 0x94, 0xa6, 0xb2, 0x51, 0xc6,
	0xc9, 0x13, 0xad, 0x41, 0x71, 0xe6, 0xf9, 0x97, 0xc4, 0x50, 0x85, 0x4f, 0x1a, 0xaf, 0xd4, 0x17,
	0x4a, 0xe3, 0xab, 0x0a, 0xb5, 0x3d, 0x1a, 0x73, 0x12, 0x90, 0xe8, 0x9d, 0xc7, 0xc7, 0x67, 0x08,
	0x41, 0x21, 0x64, 0x11, 0x17, 0xe1, 0x35, 0x2c, 0xde, 0x68, 0x0a, 0x35, 0x3f, 0x05, 0x8d, 0xf8,
	0x3c, 0x94, 0x3c, 0xf5, 0x4e, 0xf7, 0xb6, 0x25, 0x2f, 0x64, 0xc8, 0x2d, 0x77, 0x1e, 0x12, 0x5c,
	0xf5, 0x6f, 0x58, 0xe8, 0x31, 0xfc, 0x97, 0x27, 0x12, 0xb3, 0x18, 0x33, 0xdf, 0xd0, 0x44, 0xd1,
	0x7a, 0xf6, 0xc3, 0x7e, 0xea, 0x47, 0x06, 0x94, 0xbc, 0xd3, 0xd3, 0x88, 0xc4, 0xb1, 0x51, 0x68,
	0x6a, 0x1b, 0x65, 0x9c, 0x99, 0xad, 0x01, 0x54, 0x6f, 0x26, 0x41, 0x25, 0xd0, 0xac, 0xfe, 0x91,
	0xbe, 0x84, 0xfe, 0x87, 0xd5, 0xa1, 0xb3, 0x63, 0x6f, 0x5b, 0x78, 0xe4, 0xf4, 0xbb, 0x83, 0x83,
	0xfe, 0x8e, 0xae, 0xa0, 0x35, 0xd0, 0x33, 0xe7, 0xe0, 0xc0, 0x95, 0x5e, 0x15, 0x55, 0xa0, 0xf4,
	0xc6, 0x72, 0xed, 0x43, 0xeb, 0x48, 0xd7, 0x1a, 0xdf, 0x15, 0xa8, 0x3b, 0x41, 0x4c, 0x22, 0xbe,
	0xcf, 0x62, 0xca, 0x29, 0x0b, 0x10, 0x86, 0x22, 0x0d, 0x4e, 0xc9, 0x67, 0x21, 0x54, 0xbd, 0xb3,
	0x75, 0x5b, 0x2d, 0x16, 0x69, 0x4c, 0x27, 0xe1, 0xc0, 0x92, 0x0a, 0x3d, 0x84, 0x4a, 0x44, 0x7c,
	0x8f, 0xd3, 0x19, 0x19, 0x71, 0x96, 0x4e, 0x0b, 0x32, 0x97, 0xcb, 0x5a, 0x9b, 0x50, 0x14, 0x01,
	0xa8, 0x0c, 0xc5, 0x9e, 0x83, 0x87, 0xae, 0xbe, 0x84, 0x56, 0xa0, 0xb0, 0x67, 0x0d, 0x5d, 0x5d,
	0x41, 0x00, 0xcb, 0x5d, 0xbb, 0x37, 0xc0, 0xb6, 0xae, 0x26, 0x00, 0xab, 0xe7, 0xda, 0x58, 0xd7,
	0x1a, 0xdf, 0x34, 0x58, 0x96, 0xd9, 0xd1, 0x47, 0xa8, 0xe7, 0xfa, 0x5e, 0x24, 0xc3, 0x10, 0xd5,
	0xdf, 0x61, 0x9d, 0x17, 0x26, 0x89, 0xf3, 0xad, 0x90, 0xab, 0x33, 0x82, 0x55, 0x2a, 0xba, 0x1b,
	0x85, 0x69, 0x7b, 0xa2, 0x85, 0x4a, 0xe7, 0xf9, 0xbf, 0x89, 0x83, 0xeb, 0x74, 0x51, 0xf3, 0x63,
	0xa8, 0xc8, 0xad, 0x97, 0x5b, 0xa8, 0x09, 0xe5, 0x5f, 0xde, 0xed, 0x8f, 0x93, 0x7e, 0xc4, 0xf2,
	0xc1, 0x24, 0x7f, 0x27, 0xda, 0xa7, 0xdc, 0x81, 0x77, 0x41, 0x8c, 0x82, 0xd4, 0x5e, 0xba, 0xfa,
	0xde, 0x05, 0x41, 0x5b, 0x50, 0x4b, 0x01, 0x63, 0x16, 0x4c, 0xe8, 0xd4, 0x28, 0x8a, 0xde, 0xee,
	0x99, 0xf2, 0x96, 0x98, 0xd9, 0x2d, 0x31, 0x87, 0xe2, 0x96, 0xe0, 0xaa, 0x44, 0x6f, 0x0b, 0x70,
	0xeb, 0x09, 0xc0, 0x75, 0xe2, 0x64, 0xb9, 0x9c, 0xfe, 0x07, 0x6b, 0xcf, 0xd9, 0x91, 0x03, 0xdc,
	0x75, 0xdd, 0x7d, 0x5d, 0x49, 0xdc, 0x7d, 0xdb, 0x3d, 0x1c, 0xe0, 0xb7, 0xba, 0xda, 0x35, 0x7f,
	0x5c, 0xad, 0x2b, 0x3f, 0xaf, 0xd6, 0x95, 0x5f, 0x57, 0xeb, 0xca, 0x71, 0x53, 0x36, 0x49, 0x59,
	0xdb, 0x0b, 0x69, 0xfb, 0x0f, 0x17, 0xf0, 0x64, 0x59, 0x14, 0xb0, 0xf9, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x90, 0x0c, 0x14, 0xeb, 0x1f, 0x05, 0x00, 0x00,
}
