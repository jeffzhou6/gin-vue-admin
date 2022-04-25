<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="top" label-width="80px">
        <el-form-item label="相册名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="关联广告:">
          <el-select
            v-model.number="formData.adId"
            filterable
            remote
            reserve-keyword
            placeholder="请输入关键词"
            :remote-method="searchAdList"
            :loading="loading">
            <el-option
              v-for="item in adList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="最大上传数:">
          <el-input v-model.number="formData.mUN" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="相片模式:">
          <el-select v-model="formData.upType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in up_typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="相册密码有效期:">
          <el-date-picker v-model="formData.expiredAt" type="datetime" placeholder="选择日期" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="拍照咨询标题:">
          <el-input v-model="formData.contactMe" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="拍照咨询图片:">
          
        </el-form-item>
        <el-form-item label="是否显示拍摄咨询:">
          <el-switch v-model="formData.isContactMe" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="是否公开相册:">
            <el-switch v-model="formData.isPublic" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="相册封面:">
          <!-- <upload-common
              v-show="!formData.cover"
              v-model:imageCommon="imageCommon"
              upload-btn-title="封面上传"
              class="upload-btn"
              @on-success="uploadCoverSuccess"
          />
          <el-image :src="formData.cover" v-show="formData.cover">
          </el-image> -->
        <uploader v-model:url="formData.cover" />
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
  name: 'Album'
}
</script>

<script setup>
import Uploader from '@/components/upload/uploader.vue'
import {
  createAlbum,
  updateAlbum,
  findAlbum
} from '@/api/album'

import {
  getAdListByKeyword,
} from '@/api/ad'


// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('create')
const switchOptions = ref([])
const up_typeOptions = ref([])
const formData = ref({
        adId: undefined,
        // code: '',
        cover: undefined,
        contactMe: '',
        contactUrl: '',
        expiredAt: new Date(),
        isContactMe: undefined,
        // isDelete: undefined,
        isPublic: undefined,
        mUN: 5,
        name: '',
        qrcodeUrl: '',
        // uid: 0,
        upType: 1,
        // views: 0,
        })

// 广告列表
const adList = ref([])
const loading = ref(false)

const imageUrl = ref('')
const imageCommon = ref('')

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAlbum({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.realbum
        type.value = 'update'
        formData.value.adId || (formData.value.adId = undefined)
      }
    } else {
      type.value = 'create'
    }
    switchOptions.value = await getDictFunc('switch')
    up_typeOptions.value = await getDictFunc('up_type')
    console.log(route.meta);
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createAlbum(formData.value)
          break
        case 'update':
          res = await updateAlbum(formData.value)
          break
        default:
          res = await createAlbum(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        back()
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
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

const uploadCoverSuccess = (res) => {
  formData.value.cover = res
}
</script>

<style>
</style>
