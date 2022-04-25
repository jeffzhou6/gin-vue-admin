// 自动生成模板Ad
package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Ad 结构体
// 如果含有time.Time 请自行import time包
type Ad struct {
      global.GVA_MODEL
      IsDelete  *bool `json:"isDelete" form:"isDelete" gorm:"column:is_delete;comment:是否已删除;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:ad name;size:50;"`
}


// TableName Ad 表名
func (Ad) TableName() string {
  return "jyjy_ad"
}

