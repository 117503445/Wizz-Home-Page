// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// passagesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type passagesDao struct {
	*internal.PassagesDao
}

var (
	// Passages is globally public accessible object for table passages operations.
	Passages = &passagesDao{
		internal.Passages,
	}
)

// Fill with you ideas below.
