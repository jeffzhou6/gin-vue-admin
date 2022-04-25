package album

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/album"
	albumReq "github.com/flipped-aurora/gin-vue-admin/server/model/album/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AlbumPhotoApi struct {
}

var albumPhotoService = service.ServiceGroupApp.AlbumServiceGroup.AlbumPhotoService

// CreateAlbumPhoto 创建AlbumPhoto
// @Tags AlbumPhoto
// @Summary 创建AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body album.AlbumPhoto true "创建AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /albumPhoto/createAlbumPhoto [post]
func (albumPhotoApi *AlbumPhotoApi) CreateAlbumPhoto(c *gin.Context) {
	var albumPhoto album.AlbumPhoto
	_ = c.ShouldBindJSON(&albumPhoto)
	albumPhoto.AdminId = utils.GetUserID(c)
	if err := albumPhotoService.CreateAlbumPhoto(albumPhoto); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAlbumPhoto 删除AlbumPhoto
// @Tags AlbumPhoto
// @Summary 删除AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body album.AlbumPhoto true "删除AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /albumPhoto/deleteAlbumPhoto [delete]
func (albumPhotoApi *AlbumPhotoApi) DeleteAlbumPhoto(c *gin.Context) {
	var albumPhoto album.AlbumPhoto
	_ = c.ShouldBindJSON(&albumPhoto)
	if err := albumPhotoService.DeleteAlbumPhoto(albumPhoto); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAlbumPhotoByIds 批量删除AlbumPhoto
// @Tags AlbumPhoto
// @Summary 批量删除AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /albumPhoto/deleteAlbumPhotoByIds [delete]
func (albumPhotoApi *AlbumPhotoApi) DeleteAlbumPhotoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := albumPhotoService.DeleteAlbumPhotoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAlbumPhoto 更新AlbumPhoto
// @Tags AlbumPhoto
// @Summary 更新AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body album.AlbumPhoto true "更新AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /albumPhoto/updateAlbumPhoto [put]
func (albumPhotoApi *AlbumPhotoApi) UpdateAlbumPhoto(c *gin.Context) {
	var albumPhoto album.AlbumPhoto
	_ = c.ShouldBindJSON(&albumPhoto)
	if err := albumPhotoService.UpdateAlbumPhoto(albumPhoto); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAlbumPhoto 用id查询AlbumPhoto
// @Tags AlbumPhoto
// @Summary 用id查询AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query album.AlbumPhoto true "用id查询AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /albumPhoto/findAlbumPhoto [get]
func (albumPhotoApi *AlbumPhotoApi) FindAlbumPhoto(c *gin.Context) {
	var albumPhoto album.AlbumPhoto
	_ = c.ShouldBindQuery(&albumPhoto)
	if err, realbumPhoto := albumPhotoService.GetAlbumPhoto(albumPhoto.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"realbumPhoto": realbumPhoto}, c)
	}
}

// GetAlbumPhotoList 分页获取AlbumPhoto列表
// @Tags AlbumPhoto
// @Summary 分页获取AlbumPhoto列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query albumReq.AlbumPhotoSearch true "分页获取AlbumPhoto列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /albumPhoto/getAlbumPhotoList [get]
func (albumPhotoApi *AlbumPhotoApi) GetAlbumPhotoList(c *gin.Context) {
	var pageInfo albumReq.AlbumPhotoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := albumPhotoService.GetAlbumPhotoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
