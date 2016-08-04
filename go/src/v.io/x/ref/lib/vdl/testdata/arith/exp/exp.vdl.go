// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: exp

// Package exp is used to test that embedding interfaces works across packages.
// The arith.Calculator vdl interface embeds the Exp interface.
package exp

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Interface definitions

// ExpClientMethods is the client interface
// containing Exp methods.
type ExpClientMethods interface {
	Exp(_ *context.T, x float64, _ ...rpc.CallOpt) (float64, error)
}

// ExpClientStub adds universal methods to ExpClientMethods.
type ExpClientStub interface {
	ExpClientMethods
	rpc.UniversalServiceMethods
}

// ExpClient returns a client stub for Exp.
func ExpClient(name string) ExpClientStub {
	return implExpClientStub{name}
}

type implExpClientStub struct {
	name string
}

func (c implExpClientStub) Exp(ctx *context.T, i0 float64, opts ...rpc.CallOpt) (o0 float64, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Exp", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

// ExpServerMethods is the interface a server writer
// implements for Exp.
type ExpServerMethods interface {
	Exp(_ *context.T, _ rpc.ServerCall, x float64) (float64, error)
}

// ExpServerStubMethods is the server interface containing
// Exp methods, as expected by rpc.Server.
// There is no difference between this interface and ExpServerMethods
// since there are no streaming methods.
type ExpServerStubMethods ExpServerMethods

// ExpServerStub adds universal methods to ExpServerStubMethods.
type ExpServerStub interface {
	ExpServerStubMethods
	// Describe the Exp interfaces.
	Describe__() []rpc.InterfaceDesc
}

// ExpServer returns a server stub for Exp.
// It converts an implementation of ExpServerMethods into
// an object that may be used by rpc.Server.
func ExpServer(impl ExpServerMethods) ExpServerStub {
	stub := implExpServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implExpServerStub struct {
	impl ExpServerMethods
	gs   *rpc.GlobState
}

func (s implExpServerStub) Exp(ctx *context.T, call rpc.ServerCall, i0 float64) (float64, error) {
	return s.impl.Exp(ctx, call, i0)
}

func (s implExpServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implExpServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{ExpDesc}
}

// ExpDesc describes the Exp interface.
var ExpDesc rpc.InterfaceDesc = descExp

// descExp hides the desc to keep godoc clean.
var descExp = rpc.InterfaceDesc{
	Name:    "Exp",
	PkgPath: "v.io/x/ref/lib/vdl/testdata/arith/exp",
	Methods: []rpc.MethodDesc{
		{
			Name: "Exp",
			InArgs: []rpc.ArgDesc{
				{"x", ``}, // float64
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // float64
			},
		},
	},
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	return struct{}{}
}