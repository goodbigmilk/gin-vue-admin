
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="purchaseID:" prop="purchaseID">
          <el-input v-model="formData.purchaseID" :clearable="true"  placeholder="请输入purchaseID" />
       </el-form-item>
        <el-form-item label="productID:" prop="productID">
          <el-input v-model="formData.productID" :clearable="true"  placeholder="请输入productID" />
       </el-form-item>
        <el-form-item label="supplierID:" prop="supplierID">
          <el-input v-model="formData.supplierID" :clearable="true"  placeholder="请输入supplierID" />
       </el-form-item>
        <el-form-item label="warehouseID:" prop="warehouseID">
          <el-input v-model="formData.warehouseID" :clearable="true"  placeholder="请输入warehouseID" />
       </el-form-item>
        <el-form-item label="warehouseName:" prop="warehouseName">
          <el-input v-model="formData.warehouseName" :clearable="true"  placeholder="请输入warehouseName" />
       </el-form-item>
        <el-form-item label="quantity:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="status:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="userID:" prop="userID">
          <el-input v-model.number="formData.userID" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="userName:" prop="userName">
          <el-input v-model="formData.userName" :clearable="true"  placeholder="请输入userName" />
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
  createPurchases,
  updatePurchases,
  findPurchases
} from '@/api/app/purchases'

defineOptions({
    name: 'PurchasesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            purchaseID: '',
            productID: '',
            supplierID: '',
            warehouseID: '',
            warehouseName: '',
            quantity: undefined,
            status: undefined,
            userID: undefined,
            userName: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPurchases({ ID: route.query.id })
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
               res = await createPurchases(formData.value)
               break
             case 'update':
               res = await updatePurchases(formData.value)
               break
             default:
               res = await createPurchases(formData.value)
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
