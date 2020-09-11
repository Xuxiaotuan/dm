// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package unit

import (
	"context"
	"testing"

	"github.com/pingcap/check"
	"github.com/pingcap/dm/pkg/terror"
	"github.com/pingcap/errors"
)

func TestSuite(t *testing.T) {
	check.TestingT(t)
}

var _ = check.Suite(&testUnitSuite{})

type testUnitSuite struct{}

func (t *testUnitSuite) TestIsCtxCanceledProcessErr(c *check.C) {
	err := NewProcessError(context.Canceled)
	c.Assert(IsCtxCanceledProcessErr(err), check.IsTrue)

	err = NewProcessError(errors.New("123"))
	c.Assert(IsCtxCanceledProcessErr(err), check.IsFalse)

	terr := terror.ErrDBBadConn
	err = NewProcessError(terror.ErrDBBadConn)
	c.Assert(err.GetErrCode(), check.Equals, int32(terr.Code()))
	c.Assert(err.GetErrClass(), check.Equals, terr.Class().String())
	c.Assert(err.GetErrLevel(), check.Equals, terr.Level().String())
	c.Assert(err.GetMessage(), check.Equals, terr.Message())
	c.Assert(err.GetRawCause(), check.Equals, "")
}
