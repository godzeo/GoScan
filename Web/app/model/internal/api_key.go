// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ApiKey is the golang structure for table api_key.
type ApiKey struct {
    Id       int         `orm:"id"        json:"id"`        //   
    Key      string      `orm:"key"       json:"key"`       //   
    Value    string      `orm:"value"     json:"value"`     //   
    CreateAt *gtime.Time `orm:"create_at" json:"create_at"` //   
}