package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AdSearch struct{
    ad.Ad
    request.PageInfo
}
