<template>
  <main>
    <el-form :span="5" :model="form">
      <el-form-item label="键">
        <el-input type="textarea" autosize v-model="form.key" disabled placeholder="" />
      </el-form-item>
      <el-row>
        <el-col :span="11">
          <el-form-item label="过期时间">
            <el-input v-model="form.ttl" placeholder="" />
          </el-form-item>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="11">
          <el-form-item label="数据类型">
            <el-input v-model="form.type" disabled placeholder="" />
          </el-form-item>
        </el-col>
      </el-row>
      <div v-if="form.type === 'string'">
        <el-form-item label="值">
          <el-input type="textarea" :autosize="{ minRows: 6 }" v-model="form.value" placeholder="" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="updateKey">保存</el-button>
        </el-form-item>
      </div>
    </el-form>
  </main>
</template>

<script setup>
import {ref, watch} from 'vue'
import {
  GetKeyValue, UpdateKeyValue,
} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
let props = defineProps(['keyDb', 'keyConnId', 'keyKey'])
let form = ref({})
let hashDialogVisible = ref(false)
let hashDialogTitle = ref("")
let hashForm = ref({})
let listDialogVisible = ref(false)
let listForm = ref({})
let setDialogVisible = ref(false)
let setForm = ref({})
let zsetDialogVisible = ref(false)
let zsetForm = ref({})

watch(()=>props.keyKey, () => {
  console.log(props)
  getTheValue()
})

function getTheValue() {
  GetKeyValue({connection_id:props.keyConnId, db: props.keyDb, key: props.keyKey}).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    form.value = res.data
    form.value.key = props.keyKey
  })
}

function updateKey(){
  UpdateKeyValue({connection_id:props.keyConnId, db: props.keyDb, key: props.keyKey,value: form.value.value, ttl: form.value.ttl}).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    ElNotification({
      title:"更新成功",
      type: "success",
    })
  })
}
</script>

<style scoped>

</style>