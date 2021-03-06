// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// membersDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type membersDao struct {
	*internal.MembersDao
}

var (
	// Members is globally public accessible object for table members operations.
	Members = &membersDao{
		internal.Members,
	}
)

// Fill with you ideas below.
