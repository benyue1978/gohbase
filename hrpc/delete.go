// Copyright (C) 2015  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"context"
	"github.com/yuesong1978/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

// DeleteTable represents a DeleteTable HBase call
type DeleteTable struct {
	base
	namespace []byte
}

// NewDeleteTable creates a new DeleteTable request that will delete the
// given table in HBase. For use by the admin client.
func NewDeleteTable(ctx context.Context, namespace []byte, table []byte) *DeleteTable {
	return &DeleteTable{
		base: base{
			table:    table,
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		namespace: namespace,
	}
}

// Name returns the name of this RPC call.
func (dt *DeleteTable) Name() string {
	return "DeleteTable"
}

// Description returns the description of this RPC call.
func (dt *DeleteTable) Description() string {
	return dt.Name()
}

// ToProto converts the RPC into a protobuf message
func (dt *DeleteTable) ToProto() proto.Message {
	return &pb.DeleteTableRequest{
		TableName: &pb.TableName{
			Namespace: dt.namespace,
			Qualifier: dt.table,
		},
	}
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (dt *DeleteTable) NewResponse() proto.Message {
	return &pb.DeleteTableResponse{}
}
