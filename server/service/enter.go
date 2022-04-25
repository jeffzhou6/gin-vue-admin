package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/service/album"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	AlbumServiceGroup   album.ServiceGroup
	AdServiceGroup      ad.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
