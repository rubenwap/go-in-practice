

// Code generated - temporary package for codecgen - DO NOT EDIT.

package user

import (
	codec1978 "github.com/ugorji/go/codec"
	"os"
	"reflect"
	"bytes"
	"strings"
	"go/format"
)

func codecGenBoolPtr(b bool) *bool {
	return &b
}

func CodecGenTempWrite2944() {
	os.Remove("user_generated.go")
	fout, err := os.Create("user_generated.go")
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	var bJO, bS2A, bOE *bool
	_, _, _ = bJO, bS2A, bOE
	
	
	

	var typs []reflect.Type
	var typ reflect.Type
	var numfields int

	var t0 User
typ = reflect.TypeOf(t0)
	typs = append(typs, typ)
	if typ.Kind() == reflect.Struct { numfields += typ.NumField() } else { numfields += 1 }


	// println("initializing user_generated.go, buf size: 153*16",
	// 	153*16, "num fields: ", numfields)
	var out = bytes.NewBuffer(make([]byte, 0, numfields*1024)) // 153*16
	codec1978.Gen(out,
		"", "user", "2944", false,
		bJO, bS2A, bOE,
		codec1978.NewTypeInfos(strings.Split("codec,json", ",")),
		 typs...)

	bout, err := format.Source(out.Bytes())
	// println("... lengths: before formatting: ", len(out.Bytes()), ", after formatting", len(bout))
	if err != nil {
		fout.Write(out.Bytes())
		panic(err)
	}
	fout.Write(bout)
}

