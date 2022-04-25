<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="相册ID:">
          <el-input v-model.number="formData.aid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户ID:">
          <el-input v-model.number="formData.uid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="管理员ID:">
          <el-input v-model.number="formData.adminId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="照片:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="formatUrl字段:">
          <el-input v-model="formData.formatUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="照片缩略图:">
          <el-input v-model="formData.thumbUrl" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="照片名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已打印:">
          <el-switch v-model="formData.isPrinted" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="是否为封面:">
          <el-select v-model="formData.isCover" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in switchOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="照片备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否已删除:">
          <el-select v-model="formData.isDelete" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in switchOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AlbumPhoto'
}
</script>

<script setup>
import {
  createAlbumPhoto,
  updateAlbumPhoto,
  findAlbumPhoto
} from '@/api/albumPhoto'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const switchOptions = ref([])
const formData = ref({
        aid: 0,
        uid: 0,
        adminId: 0,
        url: '',
        formatUrl: '',
        thumbUrl: '',
        name: '',
        isPrinted: false,
        isCover: undefined,
        remark: '',
        isDelete: undefined,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAlbumPhoto({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.realbumPhoto
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    switchOptions.value = await getDictFunc('switch')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createAlbumPhoto(formData.value)
          break
        case 'update':
          res = await updateAlbumPhoto(formData.value)
          break
        default:
          res = await createAlbumPhoto(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
