
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="productID:" prop="productID">
          <el-input v-model="formData.productID" :clearable="true"  placeholder="请输入productID" />
       </el-form-item>
        <el-form-item label="name:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入name" />
       </el-form-item>
        <el-form-item label="price:" prop="price">
          <el-input v-model.number="formData.price" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="supplierID:" prop="supplierID">
          <el-input v-model="formData.supplierID" :clearable="true"  placeholder="请输入supplierID" />
       </el-form-item>
        <el-form-item label="description:" prop="description">
          <RichEdit v-model="formData.description"/>
       </el-form-item>
        <el-form-item label="stock:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类别名称:" prop="categoryName">
          <el-input v-model="formData.categoryName" :clearable="true"  placeholder="请输入类别名称" />
       </el-form-item>
        <el-form-item label="categoryID:" prop="categoryID">
          <el-input v-model="formData.categoryID" :clearable="true"  placeholder="请输入categoryID" />
       </el-form-item>
        <el-form-item label="imgs:" prop="imgs">
           <SelectImage v-model="formData.imgs" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="isDelete:" prop="isDelete">
          <el-input v-model.number="formData.isDelete" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createProduct,
  updateProduct,
  findProduct
} from '@/api/app/product'

defineOptions({
    name: 'ProductForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            productID: '',
            name: '',
            price: undefined,
            supplierID: '',
            description: '',
            stock: undefined,
            categoryName: '',
            categoryID: '',
            imgs: [],
            isDelete: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProduct({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createProduct(formData.value)
               break
             case 'update':
               res = await updateProduct(formData.value)
               break
             default:
               res = await createProduct(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
