/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv4{}).Type1()): VnicSpecIpv4Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv6{}).Type1()): VnicSpecIpv6Codec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv4{}).Type1()): VnicSpecIpv4Codec{},
		jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv6{}).Type1()): VnicSpecIpv6Codec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type VnicSpecIpv4Codec struct {
}

func (VnicSpecIpv4Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*VnicSpecIpv4)(ptr) == nil
}

func (VnicSpecIpv4Codec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*VnicSpecIpv4)(ptr)
	var objs []VnicSpecIpv4
	if obj != nil {
		objs = []VnicSpecIpv4{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv4{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (VnicSpecIpv4Codec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*VnicSpecIpv4)(ptr) = VnicSpecIpv4{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []VnicSpecIpv4

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv4{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*VnicSpecIpv4)(ptr) = objs[0]
			} else {
				*(*VnicSpecIpv4)(ptr) = VnicSpecIpv4{}
			}
		} else {
			*(*VnicSpecIpv4)(ptr) = VnicSpecIpv4{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj VnicSpecIpv4

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv4{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*VnicSpecIpv4)(ptr) = obj
		} else {
			*(*VnicSpecIpv4)(ptr) = VnicSpecIpv4{}
		}
	default:
		iter.ReportError("decode VnicSpecIpv4", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type VnicSpecIpv6Codec struct {
}

func (VnicSpecIpv6Codec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*VnicSpecIpv6)(ptr) == nil
}

func (VnicSpecIpv6Codec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*VnicSpecIpv6)(ptr)
	var objs []VnicSpecIpv6
	if obj != nil {
		objs = []VnicSpecIpv6{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv6{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (VnicSpecIpv6Codec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*VnicSpecIpv6)(ptr) = VnicSpecIpv6{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []VnicSpecIpv6

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv6{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*VnicSpecIpv6)(ptr) = objs[0]
			} else {
				*(*VnicSpecIpv6)(ptr) = VnicSpecIpv6{}
			}
		} else {
			*(*VnicSpecIpv6)(ptr) = VnicSpecIpv6{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj VnicSpecIpv6

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(VnicSpecIpv6{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*VnicSpecIpv6)(ptr) = obj
		} else {
			*(*VnicSpecIpv6)(ptr) = VnicSpecIpv6{}
		}
	default:
		iter.ReportError("decode VnicSpecIpv6", "unexpected JSON type")
	}
}
