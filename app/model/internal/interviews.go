// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// Interviews is the golang structure for table interviews.
type Interviews struct {
	Id         int    `orm:"id,primary"  json:"id"`         //
	Name       string `orm:"name"        json:"name"`       //
	CreateTime int64  `orm:"create_time" json:"createTime"` //
}