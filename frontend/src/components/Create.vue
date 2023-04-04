<template>
  <main>
    <el-button :type="btnType" @click="dialogVisible = true">{{title}}</el-button>
    <el-dialog
        v-model="dialogVisible"
        :title="title"
        width="60%"
        append-to-body="false"
        modal-append-to-body="false"
    >
      <el-form :model="form" label-width="120px">
        <el-form-item label="日期" prop="date">
          <el-date-picker
              v-if="data === undefined"
              v-model="form.date"
              type="date"
              placeholder="选择日期"
              style="width: 100%"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
          ></el-date-picker>
          <el-date-picker
              v-else
              v-model="form.date"
              type="date"
              placeholder="选择日期"
              style="width: 100%"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              readonly
          ></el-date-picker>
        </el-form-item>

        <el-row :gutter="2">
          <el-col :span="10">
            <el-form-item label="" prop="contesta">
              <el-input v-if="data === undefined" v-model="form.contesta" placeholder="请输入比赛A"></el-input>
              <el-input v-else v-model="form.contesta" placeholder="请输入比赛A" readonly></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="10">
            <el-form-item label="" prop="contestb">
              <el-input v-model="form.contestb" placeholder="请输入比赛B"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

<!--        让球-->
        <el-row :gutter="2">
          <el-col :span="10">
            <el-form-item label="让球" prop="typea">
              <el-input v-model="form.typea"  default="0" readonly></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="10">
            <el-form-item label="" prop="typeb">
              <el-input v-model="form.typeb"  default="1/2" readonly></el-input>
            </el-form-item>
          </el-col>
        </el-row>

<!--        胜利数-->
        <el-row :gutter="2">
          <el-col :span="10">
            <el-form-item label="胜利数" prop="wina">
              <el-input v-model="form.wina" placeholder="比赛A胜利数" ></el-input>
            </el-form-item>
          </el-col>
          <span style="margin: 1vh -8vh 0 6vh">VS</span>

          <el-col :span="10">
            <el-form-item label="" prop="winb">
              <el-input v-model="form.winb" placeholder="比赛B胜利数" ></el-input>
            </el-form-item>
          </el-col>
        </el-row>

<!--        平局数-->
        <el-row :gutter="2">
          <el-col :span="10">
            <el-form-item label="平局数" prop="drawa">
              <el-input v-model="form.drawa" placeholder="比赛A平局数"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="10">
            <el-form-item label="" prop="drawb">
              <el-input v-model="form.drawb" placeholder="比赛B平局数"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

<!--        失败数-->
        <el-row :gutter="2">
          <el-col :span="10">
            <el-form-item label="失败数" prop="losea">
              <el-input v-model="form.losea" placeholder="比赛A失败数"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="10">
            <el-form-item label="" prop="loseb">
              <el-input v-model="form.loseb" placeholder="比赛B失败数"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item>
          <el-button v-if="data === undefined" type="primary" @click="onSubmit">创建</el-button>
          <el-button v-else type="primary" @click="editRecord">保存</el-button>
          <el-button @click="dialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </main>
</template>

<script setup >
import {ref, reactive, watch} from "vue";
import { SaveData,UpdateData } from "../../wailsjs/go/main/App.js";
import { ElNotification } from "element-plus";

let props = defineProps(['title', 'btnType', 'data'])
const dialogVisible = ref(false)
const emits = defineEmits(['emit-connection-list'])
let form = ref({typea: "0", typeb: "1/2"})
if (props.data !== undefined) {
form.value = props.data
  form.value.typea = "0"
  form.value.typeb = "1/2"
}
// console.log("create", form.value)
// 侦听props.data上的变化
watch(() => props.data, (newData, oldData) => {
  // 将新的数据赋值给form
  form.value = newData
  form.value.typea = "0"
  form.value.typeb = "1/2"
})
// console.log("create -1 ", form.value)


function onSubmit() {
  SaveData({date: form.value.date, contesta: form.value.contesta, contestb: form.value.contestb,
    typea: form.value.typea, wina: form.value.wina, drawa: form.value.drawa, losea: form.value.losea,
    typeb: form.value.typeb, winb: form.value.winb, drawb: form.value.drawb, loseb: form.value.loseb}).then(res => {
    dialogVisible.value = false
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
    // 获取新的连接地址
    emits('emit-connection-list')
  })
}

function editRecord() {
  UpdateData({date: form.value.date, contesta: form.value.contesta, contestb: form.value.contestb,
    typea: form.value.typea, wina: form.value.wina, drawa: form.value.drawa, losea: form.value.losea,
    typeb: form.value.typeb, winb: form.value.winb, drawb: form.value.drawb, loseb: form.value.loseb}).then(res => {
    dialogVisible.value = false
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
    // 获取新的连接地址
    emits('emit-connection-list')
  })
}

</script>

<style scoped>

</style>