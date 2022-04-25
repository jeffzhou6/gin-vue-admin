package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/album"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	AlbumApiGroup   album.ApiGroup
	AdApiGroup      ad.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
