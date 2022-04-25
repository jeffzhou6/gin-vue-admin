package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/album"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AlbumPhotoSearch struct {
	album.AlbumPhoto
	request.PageInfo
	keyword string `json: "keyword" form: "keyword" `
}
