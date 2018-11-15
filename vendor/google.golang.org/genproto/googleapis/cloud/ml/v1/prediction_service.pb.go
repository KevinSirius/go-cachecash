// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/ml/v1/prediction_service.proto

package ml

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request for predictions to be issued against a trained model.
//
// The body of the request is a single JSON object with a single top-level
// field:
//
// <dl>
//   <dt>instances</dt>
//   <dd>A JSON array containing values representing the instances to use for
//       prediction.</dd>
// </dl>
//
// The structure of each element of the instances list is determined by your
// model's input definition. Instances can include named inputs or can contain
// only unlabeled values.
//
// Not all data includes named inputs. Some instances will be simple
// JSON values (boolean, number, or string). However, instances are often lists
// of simple values, or complex nested lists. Here are some examples of request
// bodies:
//
// CSV data with each row encoded as a string value:
// <pre>
// {"instances": ["1.0,true,\\"x\\"", "-2.0,false,\\"y\\""]}
// </pre>
// Plain text:
// <pre>
// {"instances": ["the quick brown fox", "la bruja le dio"]}
// </pre>
// Sentences encoded as lists of words (vectors of strings):
// <pre>
// {
//   "instances": [
//     ["the","quick","brown"],
//     ["la","bruja","le"],
//     ...
//   ]
// }
// </pre>
// Floating point scalar values:
// <pre>
// {"instances": [0.0, 1.1, 2.2]}
// </pre>
// Vectors of integers:
// <pre>
// {
//   "instances": [
//     [0, 1, 2],
//     [3, 4, 5],
//     ...
//   ]
// }
// </pre>
// Tensors (in this case, two-dimensional tensors):
// <pre>
// {
//   "instances": [
//     [
//       [0, 1, 2],
//       [3, 4, 5]
//     ],
//     ...
//   ]
// }
// </pre>
// Images can be represented different ways. In this encoding scheme the first
// two dimensions represent the rows and columns of the image, and the third
// contains lists (vectors) of the R, G, and B values for each pixel.
// <pre>
// {
//   "instances": [
//     [
//       [
//         [138, 30, 66],
//         [130, 20, 56],
//         ...
//       ],
//       [
//         [126, 38, 61],
//         [122, 24, 57],
//         ...
//       ],
//       ...
//     ],
//     ...
//   ]
// }
// </pre>
// JSON strings must be encoded as UTF-8. To send binary data, you must
// base64-encode the data and mark it as binary. To mark a JSON string
// as binary, replace it with a JSON object with a single attribute named `b64`:
// <pre>{"b64": "..."} </pre>
// For example:
//
// Two Serialized tf.Examples (fake data, for illustrative purposes only):
// <pre>
// {"instances": [{"b64": "X5ad6u"}, {"b64": "IA9j4nx"}]}
// </pre>
// Two JPEG image byte strings (fake data, for illustrative purposes only):
// <pre>
// {"instances": [{"b64": "ASa8asdf"}, {"b64": "JLK7ljk3"}]}
// </pre>
// If your data includes named references, format each instance as a JSON object
// with the named references as the keys:
//
// JSON input data to be preprocessed:
// <pre>
// {
//   "instances": [
//     {
//       "a": 1.0,
//       "b": true,
//       "c": "x"
//     },
//     {
//       "a": -2.0,
//       "b": false,
//       "c": "y"
//     }
//   ]
// }
// </pre>
// Some models have an underlying TensorFlow graph that accepts multiple input
// tensors. In this case, you should use the names of JSON name/value pairs to
// identify the input tensors, as shown in the following exmaples:
//
// For a graph with input tensor aliases "tag" (string) and "image"
// (base64-encoded string):
// <pre>
// {
//   "instances": [
//     {
//       "tag": "beach",
//       "image": {"b64": "ASa8asdf"}
//     },
//     {
//       "tag": "car",
//       "image": {"b64": "JLK7ljk3"}
//     }
//   ]
// }
// </pre>
// For a graph with input tensor aliases "tag" (string) and "image"
// (3-dimensional array of 8-bit ints):
// <pre>
// {
//   "instances": [
//     {
//       "tag": "beach",
//       "image": [
//         [
//           [138, 30, 66],
//           [130, 20, 56],
//           ...
//         ],
//         [
//           [126, 38, 61],
//           [122, 24, 57],
//           ...
//         ],
//         ...
//       ]
//     },
//     {
//       "tag": "car",
//       "image": [
//         [
//           [255, 0, 102],
//           [255, 0, 97],
//           ...
//         ],
//         [
//           [254, 1, 101],
//           [254, 2, 93],
//           ...
//         ],
//         ...
//       ]
//     },
//     ...
//   ]
// }
// </pre>
// If the call is successful, the response body will contain one prediction
// entry per instance in the request body. If prediction fails for any
// instance, the response body will contain no predictions and will contian
// a single error entry instead.
type PredictRequest struct {
	// Required. The resource name of a model or a version.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//
	// Required. The prediction request body.
	HttpBody             *httpbody.HttpBody `protobuf:"bytes,2,opt,name=http_body,json=httpBody,proto3" json:"http_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PredictRequest) Reset()         { *m = PredictRequest{} }
func (m *PredictRequest) String() string { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()    {}
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca7339bffd66a462, []int{0}
}

func (m *PredictRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictRequest.Unmarshal(m, b)
}
func (m *PredictRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictRequest.Marshal(b, m, deterministic)
}
func (m *PredictRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictRequest.Merge(m, src)
}
func (m *PredictRequest) XXX_Size() int {
	return xxx_messageInfo_PredictRequest.Size(m)
}
func (m *PredictRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PredictRequest proto.InternalMessageInfo

func (m *PredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictRequest) GetHttpBody() *httpbody.HttpBody {
	if m != nil {
		return m.HttpBody
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "google.cloud.ml.v1.PredictRequest")
}

func init() {
	proto.RegisterFile("google/cloud/ml/v1/prediction_service.proto", fileDescriptor_ca7339bffd66a462)
}

var fileDescriptor_ca7339bffd66a462 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xa6, 0xe3, 0xc7, 0x4f, 0x17, 0xc1, 0x43, 0x10, 0x9d, 0x45, 0x64, 0xd4, 0xcb, 0x9c, 0x90,
	0xd0, 0xe9, 0x69, 0xe2, 0x65, 0x27, 0x6f, 0x96, 0x79, 0x10, 0xbc, 0x8c, 0xac, 0x0d, 0x59, 0x24,
	0xcd, 0x1b, 0xdb, 0xac, 0x30, 0xc4, 0x8b, 0x37, 0xcf, 0x7e, 0x34, 0xbf, 0x82, 0x1f, 0x44, 0xd2,
	0x04, 0x99, 0xd4, 0xdb, 0x4b, 0xde, 0xe7, 0x79, 0x9f, 0x3f, 0x41, 0x17, 0x02, 0x40, 0x28, 0x4e,
	0x73, 0x05, 0xeb, 0x82, 0x96, 0x8a, 0x36, 0x29, 0x35, 0x15, 0x2f, 0x64, 0x6e, 0x25, 0xe8, 0x45,
	0xcd, 0xab, 0x46, 0xe6, 0x9c, 0x98, 0x0a, 0x2c, 0x60, 0xec, 0xc1, 0xa4, 0x05, 0x93, 0x52, 0x91,
	0x26, 0x8d, 0x4f, 0xc2, 0x01, 0x66, 0x24, 0x65, 0x5a, 0x83, 0x65, 0x8e, 0x58, 0x7b, 0x46, 0x7c,
	0xbc, 0xb5, 0x5d, 0x59, 0x6b, 0x96, 0x50, 0x6c, 0xfc, 0x2a, 0x79, 0x40, 0xfb, 0x99, 0x17, 0x9a,
	0xf3, 0xe7, 0x35, 0xaf, 0x2d, 0xc6, 0xe8, 0x9f, 0x66, 0x25, 0x1f, 0x44, 0xc3, 0x68, 0xd4, 0x9f,
	0xb7, 0x33, 0x4e, 0x51, 0xdf, 0xf1, 0x16, 0x8e, 0x38, 0xe8, 0x0d, 0xa3, 0xd1, 0xde, 0xe4, 0x80,
	0x04, 0x1b, 0xcc, 0x48, 0x72, 0x6b, 0xad, 0x99, 0x41, 0xb1, 0x99, 0xef, 0xae, 0xc2, 0x34, 0x79,
	0x8f, 0xd0, 0xd1, 0x9d, 0x56, 0x52, 0xf3, 0xec, 0x27, 0xc8, 0xbd, 0xcf, 0x81, 0x35, 0xda, 0x09,
	0x8f, 0x38, 0x21, 0xdd, 0x34, 0xe4, 0xb7, 0xa3, 0xf8, 0x4f, 0xa9, 0xe4, 0xfc, 0xed, 0xf3, 0xeb,
	0xa3, 0x77, 0x96, 0x9c, 0xba, 0xb2, 0x5e, 0x9c, 0xcd, 0x1b, 0x53, 0xc1, 0x13, 0xcf, 0x6d, 0x4d,
	0xc7, 0xe3, 0xd7, 0x69, 0xe8, 0x6f, 0x1a, 0x8d, 0x67, 0x0a, 0xc5, 0x39, 0x94, 0x1d, 0x25, 0x77,
	0xae, 0x49, 0x67, 0x87, 0x1d, 0x83, 0x99, 0xab, 0x26, 0x8b, 0x1e, 0xaf, 0x02, 0x43, 0x80, 0x62,
	0x5a, 0x10, 0xa8, 0x04, 0x15, 0x5c, 0xb7, 0xc5, 0x51, 0xbf, 0x62, 0x46, 0xd6, 0xdb, 0xbf, 0x76,
	0x5d, 0xaa, 0xe5, 0xff, 0x16, 0x70, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x81, 0x8e, 0x25, 0xca,
	0xd5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnlinePredictionServiceClient is the client API for OnlinePredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnlinePredictionServiceClient interface {
	// Performs prediction on the data in the request.
	//
	// **** REMOVE FROM GENERATED DOCUMENTATION
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type onlinePredictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewOnlinePredictionServiceClient(cc *grpc.ClientConn) OnlinePredictionServiceClient {
	return &onlinePredictionServiceClient{cc}
}

func (c *onlinePredictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.ml.v1.OnlinePredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlinePredictionServiceServer is the server API for OnlinePredictionService service.
type OnlinePredictionServiceServer interface {
	// Performs prediction on the data in the request.
	//
	// **** REMOVE FROM GENERATED DOCUMENTATION
	Predict(context.Context, *PredictRequest) (*httpbody.HttpBody, error)
}

func RegisterOnlinePredictionServiceServer(s *grpc.Server, srv OnlinePredictionServiceServer) {
	s.RegisterService(&_OnlinePredictionService_serviceDesc, srv)
}

func _OnlinePredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlinePredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1.OnlinePredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlinePredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnlinePredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.ml.v1.OnlinePredictionService",
	HandlerType: (*OnlinePredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _OnlinePredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/ml/v1/prediction_service.proto",
}
