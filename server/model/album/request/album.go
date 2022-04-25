package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/album"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AlbumSearch struct {
	album.Album
	request.PageInfo
}
