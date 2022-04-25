// 自动生成模板AlbumPhoto
package album

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AlbumPhoto 结构体
// 如果含有time.Time 请自行import time包
type AlbumPhoto struct {
	global.GVA_MODEL
	Aid            uint   `json:"aid" form:"aid" gorm:"column:aid;default:0;comment:相册ID;size:10;"`
	Uid            uint   `json:"uid" form:"uid" gorm:"column:uid;default:0;comment:用户ID;size:10;"`
	AdminId        uint   `json:"adminId" form:"adminId" gorm:"column:admin_id;default:0;comment:管理员ID;size:10;"`
	Url            string `json:"url" form:"url" gorm:"column:url;default:'';comment:照片;size:100;"`
	FormatUrl      string `json:"formatUrl" form:"formatUrl" gorm:"column:format_url;default:'';comment:水印照片;size:100;"`
	ThumbUrl       string `json:"thumbUrl" form:"thumbUrl" gorm:"column:thumb_url;default:'';comment:照片缩略图;size:100;"`
	Name           string `json:"name" form:"name" gorm:"column:name;default:'';comment:照片名称;size:100;"`
	IsPrinted      *bool  `json:"isPrinted" form:"isPrinted" gorm:"column:is_printed;default:0;comment:是否已打印;"`
	IsCover        uint   `json:"isCover" form:"isCover" gorm:"column:is_cover;default:0;comment:是否为封面;"`
	Remark         string `json:"remark" form:"remark" gorm:"column:remark;default:'';comment:照片备注;size:30;"`
	IsDelete       uint   `json:"isDelete" form:"isDelete" gorm:"column:is_delete;default:0;comment:是否已删除;"`
	FormatThumbUrl string `json:"formatThumbUrl" form:"formatThumbUrl" `
	AlbumName      string `json:"albumName"`
}

// TableName AlbumPhoto 表名
func (AlbumPhoto) TableName() string {
	return "jyjy_album_photo"
}
