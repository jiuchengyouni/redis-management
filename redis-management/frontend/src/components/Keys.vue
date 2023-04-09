<template>
  <main>
    <el-dialog
        v-model="keyDialogVisible"
        title="创建键"
        width="50%"
    >
      <el-form :model="keyForm" label-width="120px">
        <el-form-item label="键的名称">
          <el-input placeholder="请输入键名" v-model="keyForm.key" />
        </el-form-item>
        <el-form-item label="键的类型">
          <el-select v-model="keyForm.type" placeholder="请选择键的类型">
            <el-option value="string" label="string"></el-option>
            <el-option value="hash" label="hash"></el-option>
            <el-option value="list" label="list"></el-option>
            <el-option value="set" label="set"></el-option>
            <el-option value="zset" label="zset"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="createKey">创建</el-button>
          <el-button @click="keyDialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-form :inline="true" :model="form" class="demo-form-inline">
      <el-form-item label="">
        <el-input v-model="form.keyword" placeholder="输入键的信息" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="getKeyList">查询</el-button>
      </el-form-item>
    </el-form>
    <el-button @click="keyDialogVisible = true" style="width: 60%; margin-bottom: 12px;">创建键</el-button>
    <div v-for="item in keys" @click="selectKeyKey(item)">
      <div v-if="item===selectKey" >
        <div class="item key-select-item">
          <div>{{item}}</div>
          <el-popconfirm title="确认删除?" @confirm="deleteKey(item)">
            <template #reference>
              <el-button text type="danger" @click.stop>删除</el-button>
            </template>
          </el-popconfirm>
        </div>
      </div>
      <div v-else>
        <div class="item key-item">
          <div>{{item}}</div>
          <el-popconfirm title="确认删除?" @confirm="deleteKey(item)">
            <template #reference>
              <el-button text type="danger" @click.stop>删除</el-button>
            </template>
          </el-popconfirm>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup>
import {defineProps, ref} from "vue";
import { reactive } from 'vue'
import {CreateValue, DeleteKeyValue, KeyList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
import {watch} from "vue";
let selectKey = ref()
let props=defineProps(['keyDb','keyConnId'])
let emits=defineEmits(['emit-select-key'])
let keys=ref()
let keyDialogVisible=ref(false)
let keyForm=ref({})
console.log(props)
const form = reactive({
  keyword: '',
})
watch(props,()=>{
  getKeyList()
})

function getKeyList() {
  let data={
    connection_id:props.keyConnId,
    db:props.keyDb,
    keyword:form.keyword,
  }
  KeyList(data).then(res=>{
    if (res.code!==200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
      return
    }
    keys.value=res.data
  })
}

function selectKeyKey(key) {
  selectKey.value = key
  emits('emit-select-key', key)
}

function deleteKey(key) {
  DeleteKeyValue({connection_id: props.keyConnId, db: props.keyDb, key: key}).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    ElNotification({
      title:res.msg,
      type: "success",
    })
    getKeyList()
  })
}

function createKey(){
  CreateValue({connection_id:props.keyConnId,db:props.keyDb,key:keyForm.value.key,type:keyForm.value.type}).then(res=>{
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    ElNotification({
      title:"新增成功",
      type: "success",
    })
    keyDialogVisible.value=false
    getKeyList()
  })
}
</script>

<style scoped>
.key-item{
  color: #303133;
  background-color: #79bbff;
  padding: 7px 14px;
}
.key-select-item{
  color: #606266;
  background-color: #cfd9df;
  padding: 7px 14px;
}
</style>