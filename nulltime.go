// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2013 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

// NullTime represents a time.Time that may be NULL.
//
// Deprecated: This is an alias to the sql.NullTime type,
// which you should use directly instead.
type NullTime = sql.NullTime
