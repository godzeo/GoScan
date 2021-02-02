// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ScanPort is the golang structure for table scan_port.
type ScanPort struct {
    Id            int         `orm:"id"             json:"id"`             //   
    CusName       string      `orm:"cus_name"       json:"cus_name"`       //   
    Host          string      `orm:"host"           json:"host"`           //   
    Port          int         `orm:"port"           json:"port"`           //   
    ServiceName   string      `orm:"service_name"   json:"service_name"`   //   
    VendorProduct string      `orm:"vendor_product" json:"vendor_product"` //   
    Version       string      `orm:"version"        json:"version"`        //   
    Flag          bool        `orm:"flag"           json:"flag"`           //   
    NsqFlag       bool        `orm:"nsq_flag"       json:"nsq_flag"`       //   
    HttpFlag      bool        `orm:"http_flag"      json:"http_flag"`      //   
    Url           string      `orm:"url"            json:"url"`            //   
    Code          int         `orm:"code"           json:"code"`           //   
    Title         string      `orm:"title"          json:"title"`          //   
    CreateAt      *gtime.Time `orm:"create_at"      json:"create_at"`      //   
}