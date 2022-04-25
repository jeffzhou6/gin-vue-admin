package album

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlbumPhotoRouter struct {
}

// InitAlbumPhotoRouter 初始化 AlbumPhoto 路由信息
func (s *AlbumPhotoRouter) InitAlbumPhotoRouter(Router *gin.RouterGroup) {
	albumPhotoRouter := Router.Group("albumPhoto").Use(middleware.OperationRecord())
	albumPhotoRouterWithoutRecord := Router.Group("albumPhoto")
	var albumPhotoApi = v1.ApiGroupApp.AlbumApiGroup.AlbumPhotoApi
	{
		albumPhotoRouter.POST("createAlbumPhoto", albumPhotoApi.CreateAlbumPhoto)   // 新建AlbumPhoto
		albumPhotoRouter.DELETE("deleteAlbumPhoto", albumPhotoApi.DeleteAlbumPhoto) // 删除AlbumPhoto
		albumPhotoRouter.DELETE("deleteAlbumPhotoByIds", albumPhotoApi.DeleteAlbumPhotoByIds) // 批量删除AlbumPhoto
		albumPhotoRouter.PUT("updateAlbumPhoto", albumPhotoApi.UpdateAlbumPhoto)    // 更新AlbumPhoto
	}
	{
		albumPhotoRouterWithoutRecord.GET("findAlbumPhoto", albumPhotoApi.FindAlbumPhoto)        // 根据ID获取AlbumPhoto
		albumPhotoRouterWithoutRecord.GET("getAlbumPhotoList", albumPhotoApi.GetAlbumPhotoList)  // 获取AlbumPhoto列表
	}
}
