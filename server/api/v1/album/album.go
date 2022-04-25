package album

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/album"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    albumReq "github.com/flipped-aurora/gin-vue-admin/server/model/album/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AlbumApi struct {
}

var albumService = service.ServiceGroupApp.AlbumServiceGroup.AlbumService


// CreateAlbum 创建Album
// @Tags Album
// @Summary 创建Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body album.Album true "创建Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /album/createAlbum [post]
func (albumApi *AlbumApi) CreateAlbum(c *gin.Context) {
	var album album.Album
	_ = c.ShouldBindJSON(&album)
	if err := albumService.CreateAlbum(album); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAlbum 删除Album
// @Tags Album
// @Summary 删除Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body album.Album true "删除Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /album/deleteAlbum [delete]
func (albumApi *AlbumApi) DeleteAlbum(c *gin.Context) {
	var album album.Album
	_ = c.ShouldBindJSON(&album)
	if err := albumService.DeleteAlbum(album); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAlbumByIds 批量删除Album
// @Tags Album
// @Summary 批量删除Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /album/deleteAlbumByIds [delete]
func (albumApi *AlbumApi) DeleteAlbumByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := albumService.DeleteAlbumByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAlbum 更新Album
// @Tags Album
// @Summary 更新Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body album.Album true "更新Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /album/updateAlbum [put]
func (albumApi *AlbumApi) UpdateAlbum(c *gin.Context) {
	var album album.Album
	_ = c.ShouldBindJSON(&album)
	if err := albumService.UpdateAlbum(album); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAlbum 用id查询Album
// @Tags Album
// @Summary 用id查询Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query album.Album true "用id查询Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /album/findAlbum [get]
func (albumApi *AlbumApi) FindAlbum(c *gin.Context) {
	var album album.Album
	_ = c.ShouldBindQuery(&album)
	if err, realbum := albumService.GetAlbum(album.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"realbum": realbum}, c)
	}
}

// GetAlbumList 分页获取Album列表
// @Tags Album
// @Summary 分页获取Album列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query albumReq.AlbumSearch true "分页获取Album列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /album/getAlbumList [get]
func (albumApi *AlbumApi) GetAlbumList(c *gin.Context) {
	var pageInfo albumReq.AlbumSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := albumService.GetAlbumInfoList(pageInfo); err != nil {
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
