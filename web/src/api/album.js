import service from '@/utils/request'

// @Tags Album
// @Summary 创建Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Album true "创建Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /album/createAlbum [post]
export const createAlbum = (data) => {
  return service({
    url: '/album/createAlbum',
    method: 'post',
    data
  })
}

// @Tags Album
// @Summary 删除Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Album true "删除Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /album/deleteAlbum [delete]
export const deleteAlbum = (data) => {
  return service({
    url: '/album/deleteAlbum',
    method: 'delete',
    data
  })
}

// @Tags Album
// @Summary 删除Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /album/deleteAlbum [delete]
export const deleteAlbumByIds = (data) => {
  return service({
    url: '/album/deleteAlbumByIds',
    method: 'delete',
    data
  })
}

// @Tags Album
// @Summary 更新Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Album true "更新Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /album/updateAlbum [put]
export const updateAlbum = (data) => {
  return service({
    url: '/album/updateAlbum',
    method: 'put',
    data
  })
}

// @Tags Album
// @Summary 用id查询Album
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Album true "用id查询Album"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /album/findAlbum [get]
export const findAlbum = (params) => {
  return service({
    url: '/album/findAlbum',
    method: 'get',
    params
  })
}

// @Tags Album
// @Summary 分页获取Album列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Album列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /album/getAlbumList [get]
export const getAlbumList = (params) => {
  return service({
    url: '/album/getAlbumList',
    method: 'get',
    params
  })
}
