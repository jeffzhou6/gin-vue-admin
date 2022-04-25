package album

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/album"
	albumReq "github.com/flipped-aurora/gin-vue-admin/server/model/album/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

type AlbumPhotoService struct {
}

// CreateAlbumPhoto 创建AlbumPhoto记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumPhotoService *AlbumPhotoService) CreateAlbumPhoto(albumPhoto album.AlbumPhoto) (err error) {
	err = global.GVA_DB.Create(&albumPhoto).Error
	return err
}

// DeleteAlbumPhoto 删除AlbumPhoto记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumPhotoService *AlbumPhotoService) DeleteAlbumPhoto(albumPhoto album.AlbumPhoto) (err error) {
	err = global.GVA_DB.Delete(&albumPhoto).Error
	return err
}

// DeleteAlbumPhotoByIds 批量删除AlbumPhoto记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumPhotoService *AlbumPhotoService) DeleteAlbumPhotoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]album.AlbumPhoto{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAlbumPhoto 更新AlbumPhoto记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumPhotoService *AlbumPhotoService) UpdateAlbumPhoto(albumPhoto album.AlbumPhoto) (err error) {
	err = global.GVA_DB.Save(&albumPhoto).Error
	return err
}

// GetAlbumPhoto 根据id获取AlbumPhoto记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumPhotoService *AlbumPhotoService) GetAlbumPhoto(id uint) (err error, albumPhoto album.AlbumPhoto) {
	err = global.GVA_DB.Where("id = ?", id).First(&albumPhoto).Error
	return
}

// GetAlbumPhotoInfoList 分页获取AlbumPhoto记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumPhotoService *AlbumPhotoService) GetAlbumPhotoInfoList(info albumReq.AlbumPhotoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&album.AlbumPhoto{})
	var albumPhotos []album.AlbumPhoto
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Aid != 0 {
		db = db.Where("aid = ?", info.Aid)
	}
	if info.Keyword != "" {
		db = db.Where("id like ? or name like ?", info.Keyword+"%", "%"+info.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&albumPhotos).Error

	oss := upload.NewOss()
	for k, v := range albumPhotos {
		v.ThumbUrl = oss.GetImageThumbUrl(v.Url)
		v.FormatThumbUrl = oss.GetImageThumbUrl(v.FormatUrl)
		err, albumName := ServiceGroupApp.AlbumService.GetAlbumName(v.Aid)
		if err != nil {
			continue
		}
		v.AlbumName = albumName
		albumPhotos[k] = v
	}

	return err, albumPhotos, total
}

// GetAlbumPhotoCount 获取AlbumPhoto数量
// Author [piexlmax](https://github.com/piexlmax)
func (albumPhotoService *AlbumPhotoService) GetAlbumPhotoCount(aid uint) uint {
	// 创建db
	db := global.GVA_DB.Model(&album.AlbumPhoto{})
	db = db.Where("aid = ?", aid)
	var count int64
	db.Count(&count)
	return uint(count)
}
