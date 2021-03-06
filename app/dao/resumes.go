// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"wizz-home-page/app/dao/internal"
)

// resumesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type resumesDao struct {
	*internal.ResumesDao
}

var (
	// Resumes is globally public accessible object for table resumes operations.
	Resumes = &resumesDao{
		internal.Resumes,
	}
)

// Fill with you ideas below.
