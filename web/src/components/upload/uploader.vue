<template>
    <el-upload
      ref="uploader"
      class="avatar-uploader"
      :action="`${path}/fileUploadAndDownload/upload`"
      :headers="{ 'x-token': userStore.token }"
      :on-success="handleImageSuccess"
      :on-remove="handleImageRemove"
      :before-upload="beforeImageUpload"
      list-type="picture-card" 
      :limit="limit"
      :auto-upload="autoUpload"
      :multiple="multiple"
      :file-list="fileList"
      >
      <el-icon><Plus /></el-icon>
      <template #tip v-if="tips.length > 0">
        <div class="el-upload__tip text-red">
          {{ tips }}
        </div>
      </template>
    </el-upload>
</template>

<script setup>
import ImageCompress from '@/utils/image'
import { computed, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_API)
const emit = defineEmits(['on-success', 'on-remove', 'update:url'])

const uploader = ref(null)

const fileList = computed(() => {
  if(!props.url){
    return []
  }
  return [{
    name: '',
    url: props.url
  }]
})

const props = defineProps({
  url: {
      type: String,
      default: ''
  },
  limit: {
    type: Number,
    default: 1
  },
  autoUpload: {
    type: Boolean,
    default: true,
  },
  multiple: {
    type: Boolean,
    default: false,
  },
  fileSize: {
    type: Number,
    default: 2048 // 2M 超出后执行压缩
  },
  maxWH: {
    type: Number,
    default: 1920 // 图片长宽上限
  },
  tips: {
      type: String,
      default: "" 
  },
})

const beforeImageUpload = (file) => {
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  if (!isJPG && !isPng) {
    ElMessage.error('上传头像图片只能是 jpg或png 格式!')
    return false
  }

  const isRightSize = file.size / 1024 < props.fileSize
  if (!isRightSize) {
    // 压缩
    const compress = new ImageCompress(file, props.fileSize, props.maxWH)
    return compress.compress()
  }
  return isRightSize
}

const handleImageRemove = (res) => {
    emit('on-remove', res)
    updateUrl('')
}


const handleImageSuccess = (res) => {
  const { data } = res
  if (data.file) {
    updateUrl(data.file.url)
    emit('on-success', data.file.url)
  }
}

const updateUrl = (url) => {
    emit('update:url', url)
}

defineExpose({
  submit: () => {
    uploader.value.submit()
  },
})


</script>

<script>

export default {
  name: 'Uploader',
  methods: {

  }
}
</script>

<style lang="scss" scoped>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>
