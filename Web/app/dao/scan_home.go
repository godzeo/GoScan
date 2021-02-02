// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"Web/app/dao/internal"
)

// scanHomeDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type scanHomeDao struct {
	*internal.ScanHomeDao
}

var (
	// ScanHome is globally public accessible object for table scan_home operations.
	ScanHome = &scanHomeDao{
		internal.ScanHome,
	}
)

// Fill with you ideas below.