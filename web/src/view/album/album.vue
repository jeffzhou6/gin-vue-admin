<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="相册名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="goToCreateAlbum">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="60" />
        <el-table-column align="left" label="相册名称" width="120">
            <template #default="scope"><el-link type="primary" @click="goToAlbumPhoto(scope.row)">{{ scope.row.name }}</el-link></template>
        </el-table-column>
        <el-table-column align="left" label="封面图" width="120" >
            <template #default="scope"><el-image :src="scope.row.cover" /></template>
        </el-table-column>
        <el-table-column align="left" label="小程序码" prop="qrcodeUrl" width="120" >
            <template #default="scope"><el-image :src="scope.row.qrcodeUrl" /></template>
        </el-table-column>
        <el-table-column align="left" label="相册密码" prop="code" width="120" />
        <el-table-column align="left" label="密码有效期" width="160">
            <template #default="scope">{{ formatDate(scope.row.expiredAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="公开相册" prop="isPublic" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.isPublic) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户上传限制" prop="mUN" width="120" />
        <el-table-column align="left" label="相片模式" prop="upType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.upType,up_typeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="照片总量" prop="photoCount" width="120" />
        <el-table-column align="left" label="浏览数量" prop="views" width="120" />
        <el-table-column align="left" label="更新时间" width="160">
            <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" width="160">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="goToUpdateAlbum(scope.row)">编辑</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Album'
}
</script>

<script setup>
import {
  createAlbum,
  deleteAlbum,
  deleteAlbumByIds,
  updateAlbum,
  findAlbum,
  getAlbumList
} from '@/api/album'

import {
  getAdListByKeyword,
} from '@/api/ad'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import router from '@/router/index'

// 自动化生成的字典（可能为空）以及字段
const switchOptions = ref([])
const up_typeOptions = ref([])

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 广告列表
const adList = ref([])
const loading = ref(false)

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAlbumList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    switchOptions.value = await getDictFunc('switch')
    up_typeOptions.value = await getDictFunc('up_type')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteAlbumFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteAlbumByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 删除行
const deleteAlbumFunc = async (row) => {
    const res = await deleteAlbum({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


const goToCreateAlbum = () => {
  router.push({name: 'album-create'})
}

const goToUpdateAlbum = (item) => {
  router.push({name: 'album-update', query: {id: item.ID}})
}

const goToAlbumPhoto = (item) => {
  router.push({name: 'album-photo-list', query: {id: item.ID}})
}

const searchAdList = async(query) => {
  if( query.trim() === ''){
    adList.value = []
    return
  }
  loading.value = true
  const res = await getAdListByKeyword({keyword: query.trim()})
  adList.value = res.data.map(item => {
    return { value: item.ID, label:  item.name };
  });
  loading.value = false
}

</script>

<style>
</style>
