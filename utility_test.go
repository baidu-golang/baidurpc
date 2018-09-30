// Go support for Protocol Buffers RPC which compatiable with https://github.com/Baidu-ecom/Jprotobuf-rpc-socket
//
// Copyright 2002-2007 the original author or authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package pbrpc_test

import (
	"testing"
	"time"

	pbrpc "github.com/baidu-golang/pbrpc"
)

func TestTimetookInSeconds(t *testing.T) {

	now := time.Now().UnixNano()

	time.Sleep(time.Second)

	cost := pbrpc.TimetookInSeconds(now)

	if cost < 1 {
		t.Error("time took is not acceptable.")
	}
}
