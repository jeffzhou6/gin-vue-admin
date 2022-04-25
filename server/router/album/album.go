package album

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlbumRouter struct {
}

// InitAlbumRouter 初始化 Album 路由信息
func (s *AlbumRouter) InitAlbumRouter(Router *gin.RouterGroup) {
	albumRouter := Router.Group("album").Use(middleware.OperationRecord())
	albumRouterWithoutRecord := Router.Group("album")
	var albumApi = v1.ApiGroupApp.AlbumApiGroup.AlbumApi
	{
		albumRouter.POST("createAlbum", albumApi.CreateAlbum)   // 新建Album
		albumRouter.DELETE("deleteAlbum", albumApi.DeleteAlbum) // 删除Album
		albumRouter.DELETE("deleteAlbumByIds", albumApi.DeleteAlbumByIds) // 批量删除Album
		albumRouter.PUT("updateAlbum", albumApi.UpdateAlbum)    // 更新Album
	}
	{
		albumRouterWithoutRecord.GET("findAlbum", albumApi.FindAlbum)        // 根据ID获取Album
		albumRouterWithoutRecord.GET("getAlbumList", albumApi.GetAlbumList)  // 获取Album列表
	}
}
