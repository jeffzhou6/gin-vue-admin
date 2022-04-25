import service from '@/utils/request'

// @Tags AlbumPhoto
// @Summary 创建AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AlbumPhoto true "创建AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /albumPhoto/createAlbumPhoto [post]
export const createAlbumPhoto = (data) => {
  return service({
    url: '/albumPhoto/createAlbumPhoto',
    method: 'post',
    data
  })
}

// @Tags AlbumPhoto
// @Summary 删除AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AlbumPhoto true "删除AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /albumPhoto/deleteAlbumPhoto [delete]
export const deleteAlbumPhoto = (data) => {
  return service({
    url: '/albumPhoto/deleteAlbumPhoto',
    method: 'delete',
    data
  })
}

// @Tags AlbumPhoto
// @Summary 删除AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /albumPhoto/deleteAlbumPhoto [delete]
export const deleteAlbumPhotoByIds = (data) => {
  return service({
    url: '/albumPhoto/deleteAlbumPhotoByIds',
    method: 'delete',
    data
  })
}

// @Tags AlbumPhoto
// @Summary 更新AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AlbumPhoto true "更新AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /albumPhoto/updateAlbumPhoto [put]
export const updateAlbumPhoto = (data) => {
  return service({
    url: '/albumPhoto/updateAlbumPhoto',
    method: 'put',
    data
  })
}

// @Tags AlbumPhoto
// @Summary 用id查询AlbumPhoto
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AlbumPhoto true "用id查询AlbumPhoto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /albumPhoto/findAlbumPhoto [get]
export const findAlbumPhoto = (params) => {
  return service({
    url: '/albumPhoto/findAlbumPhoto',
    method: 'get',
    params
  })
}

// @Tags AlbumPhoto
// @Summary 分页获取AlbumPhoto列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AlbumPhoto列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /albumPhoto/getAlbumPhotoList [get]
export const getAlbumPhotoList = (params) => {
  return service({
    url: '/albumPhoto/getAlbumPhotoList',
    method: 'get',
    params
  })
}
