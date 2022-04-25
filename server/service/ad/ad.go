package ad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	adReq "github.com/flipped-aurora/gin-vue-admin/server/model/ad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AdService struct {
}

// CreateAd 创建Ad记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) CreateAd(ad ad.Ad) (err error) {
	err = global.GVA_DB.Create(&ad).Error
	return err
}

// DeleteAd 删除Ad记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) DeleteAd(ad ad.Ad) (err error) {
	err = global.GVA_DB.Delete(&ad).Error
	return err
}

// DeleteAdByIds 批量删除Ad记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) DeleteAdByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ad.Ad{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAd 更新Ad记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) UpdateAd(ad ad.Ad) (err error) {
	err = global.GVA_DB.Save(&ad).Error
	return err
}

// GetAd 根据id获取Ad记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) GetAd(id uint) (err error, ad ad.Ad) {
	err = global.GVA_DB.Where("id = ?", id).First(&ad).Error
	return
}

// GetAdInfoList 分页获取Ad记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) GetAdInfoList(info adReq.AdSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ad.Ad{})
	var ads []ad.Ad
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&ads).Error
	return err, ads, total
}

// GetAdInfoList 通过keyword获取Ad记录
// Author [piexlmax](https://github.com/piexlmax)
func (adService *AdService) GetAdInfoListByKeyword(keyword string) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&ad.Ad{})
	var ads []ad.Ad
	// 如果有条件搜索 下方会自动创建搜索语句
	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	err = db.Find(&ads).Error
	return err, ads
}
