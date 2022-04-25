package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/router/album"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Album   album.RouterGroup
	Ad      ad.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
