
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="供应商id:" prop="supplierID">
          <el-input v-model="formData.supplierID" :clearable="true"  placeholder="请输入供应商id" />
       </el-form-item>
        <el-form-item label="supplierName:" prop="supplierName">
          <el-input v-model="formData.supplierName" :clearable="true"  placeholder="请输入supplierName" />
       </el-form-item>
        <el-form-item label="电话号码:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入电话号码" />
       </el-form-item>
        <el-form-item label="属于哪个用户:" prop="userID">
          <el-input v-model="formData.userID" :clearable="true"  placeholder="请输入属于哪个用户" />
       </el-form-item>
        <el-form-item label="状态 0合作中，1暂停合作，2黑名单:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="累计订单:" prop="cumulativeOrder">
          <el-input v-model.number="formData.cumulativeOrder" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <RichEdit v-model="formData.remark"/>
       </el-form-item>
        <el-form-item label="是否软删:" prop="isDelete">
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
  createSupplier,
  updateSupplier,
  findSupplier
} from '@/api/app/supplier'

defineOptions({
    name: 'SupplierForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            supplierID: '',
            supplierName: '',
            phone: '',
            userID: '',
            status: undefined,
            cumulativeOrder: undefined,
            remark: '',
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
      const res = await findSupplier({ ID: route.query.id })
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
               res = await createSupplier(formData.value)
               break
             case 'update':
               res = await updateSupplier(formData.value)
               break
             default:
               res = await createSupplier(formData.value)
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
