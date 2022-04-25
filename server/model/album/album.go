// 自动生成模板Album
package album

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Album 结构体
// 如果含有time.Time 请自行import time包
type Album struct {
	global.GVA_MODEL
	AdId        *int       `json:"adId" form:"adId" gorm:"column:ad_id;default:0;comment:广告ID;size:10;"`
	Code        string     `json:"code" form:"code" gorm:"column:code;comment:相册密码;size:10;"`
	Cover       string     `json:"cover" form:"cover" gorm:"column:cover;comment:封面图;size:120;"`
	ContactMe   string     `json:"contactMe" form:"contactMe" gorm:"column:contact_me;comment:拍照咨询;size:30;"`
	ContactUrl  string     `json:"contactUrl" form:"contactUrl" gorm:"column:contact_url;comment:拍照咨询图片;size:100;"`
	ExpiredAt   *time.Time `json:"expiredAt" form:"expiredAt" gorm:"column:expired_at;comment:密码有效期;"`
	IsContactMe *int       `json:"isContactMe" form:"isContactMe" gorm:"column:is_contact_me;default:0;comment:是否显示拍摄咨询;"`
	IsDelete    *int       `json:"isDelete" form:"isDelete" gorm:"column:is_delete;default:0;comment:是否已删除;"`
	IsPublic    *int       `json:"isPublic" form:"isPublic" gorm:"column:is_public;default:0;comment:是否公开相册;"`
	MUN         *int       `json:"mUN" form:"mUN" gorm:"column:m_u_n;comment:最大上传数;size:10;"`
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:相册名称;size:100;"`
	QrcodeUrl   string     `json:"qrcodeUrl" form:"qrcodeUrl" gorm:"column:qrcode_url;comment:小程序码;size:150;"`
	Uid         *int       `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;default:0;size:10;"`
	UpType      *int       `json:"upType" form:"upType" gorm:"column:up_type;default:1;comment:相片模式;"`
	Views       *int       `json:"views" form:"views" gorm:"column:views;default:0;comment:浏览量;size:10;"`
	PhotoCount  uint       `json:"photoCount" form:"photoCount"`
}

// TableName Album 表名
func (Album) TableName() string {
	return "jyjy_album"
}
