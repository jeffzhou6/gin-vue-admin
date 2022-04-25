package album

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/album"
	albumReq "github.com/flipped-aurora/gin-vue-admin/server/model/album/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

type AlbumService struct {
}

// CreateAlbum 创建Album记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumService *AlbumService) CreateAlbum(album album.Album) (err error) {
	album.Code = utils.RandStr(10)
	if album.Cover != "" {
		oss := upload.NewOss()
		album.Cover = oss.GetImageThumbUrl(album.Cover)
	}
	err = global.GVA_DB.Create(&album).Error
	return err
}

// DeleteAlbum 删除Album记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumService *AlbumService) DeleteAlbum(album album.Album) (err error) {
	err = global.GVA_DB.Delete(&album).Error
	return err
}

// DeleteAlbumByIds 批量删除Album记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumService *AlbumService) DeleteAlbumByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]album.Album{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAlbum 更新Album记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumService *AlbumService) UpdateAlbum(album album.Album) (err error) {
	if album.Cover != "" {
		oss := upload.NewOss()
		album.Cover = oss.GetImageThumbUrl(album.Cover)
	}
	err = global.GVA_DB.Save(&album).Error
	return err
}

// GetAlbum 根据id获取Album记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumService *AlbumService) GetAlbum(id uint) (err error, album album.Album) {
	err = global.GVA_DB.Where("id = ?", id).First(&album).Error
	return
}

// GetAlbumInfoList 分页获取Album记录
// Author [piexlmax](https://github.com/piexlmax)
func (albumService *AlbumService) GetAlbumInfoList(info albumReq.AlbumSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&album.Album{})
	var albums []album.Album
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Limit(limit).Offset(offset).Find(&albums).Error

	var albumPhotoService = AlbumPhotoService{}
	for k, v := range albums {
		v.PhotoCount = albumPhotoService.GetAlbumPhotoCount(v.ID)
		albums[k] = v
	}
	return err, albums, total
}

// GetAlbumInfoList 批量获取相册名称字典
// Author [piexlmax](https://github.com/piexlmax)
func (*AlbumService) GetAlbumName(id uint) (error, string) {
	// 创建db
	db := global.GVA_DB.Model(&album.Album{})

	var album album.Album
	err := db.Where("id = ?", id).Select("name").First(&album).Error
	if err != nil {
		return err, ""
	}
	return nil, album.Name
}
