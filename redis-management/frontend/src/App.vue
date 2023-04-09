<script setup>
import ConnectionList from "./components/ConnectionList.vue"
import ConnectionManage from "./components/ConnectionManage.vue";
import Keys from "./components/Keys.vue";
import {reactive, ref} from "vue";
import {ConnectionEdit, CreateValue, DbList, GetKeyValue, KeyList, SendEmail} from "../wailsjs/go/main/App.js";
import KeyValue from "./components/KeyValue.vue";
import { useDark, useToggle } from '@vueuse/core'
import {ElNotification} from "element-plus";
import Email from "./components/Email.vue";
let updateFlag=ref(true)
let keyValue=ref()
let keyList=ref()
let keyConnId=ref()
let keyDb=ref()
let keyKey = ref()
const theme1=ref(true)
const isDark = useDark()
const toggleDark = useToggle(isDark)
function updateConnection(){
  updateFlag.value=!updateFlag.value
}

function addKeyValue(){
  CreateValue({connection_id:"4fe66383-a1d9-441a-b546-0f273321df72",db:0,key:"name2",type:"string"}).then(res=>{
    console.log(res)
  })
}

function selectDB(db,connection_id){
  keyDb.value=db
  keyConnId.value=connection_id
}
GetKeyValue({connection_id:"4fe66383-a1d9-441a-b546-0f273321df72",db:0,key:"name"}).then(res=>{
  keyList.value=res
})
function selectKey(key) {
  keyKey.value = key
}


</script>


<template>
  <el-row>
    <el-col :span="6" style="height: 100vh;padding:15px;" >
      <el-switch
          v-model="theme1"
          @click="toggleDark()"
          class="mb-2"
          active-text="白天模式"
          inactive-text="深夜模式"
      />
      <div style="margin-bottom:12px" >
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="updateConnection"/>
        <!--        <el-button @click="addKeyValue()">测试新增</el-button>-->
      </div>
      <ConnectionList @emit-select-db="selectDB" :update="updateFlag"/>
      <div align="flex-end">
        <Email/>
      </div>
    </el-col>
    <el-col :span="8" style="padding:  15px">
      <Keys :keyDb="keyDb" :keyConnId="keyConnId" @emit-select-key="selectKey" />
    </el-col>
    <el-col :span="8" style="margin-left: 40px">
      <!--      dbList==>{{keyList}}-->
      <KeyValue :keyDb="keyDb" :keyConnId="keyConnId" :keyKey="keyKey"/>
    </el-col>
  </el-row>
</template>

<style>
</style>
