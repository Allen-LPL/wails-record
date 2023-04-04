<template>
  <main>
    <div class="record-list">
      <div style="margin-bottom: 12px">
        <el-row :gutter="20">
          <el-col :span="11">
            <el-form :model="form">
              <el-form-item class="timeSearStyle" label="时间段" :label-width="formLabelWidth">
                <el-date-picker
                    class="dateStyle"
                    v-model="search"
                    type="daterange"
                    format="YYYY-MM-DD"
                    value-format="YYYY-MM-DD"
                    align="right"
                    unlink-panels
                    range-separator="至"
                    start-placeholder="开始日期"
                    end-placeholder="结束日期"
                    :default-value="[new Date(new Date().getTime() - 7 * 24 * 60 * 60 * 1000), new Date()]"
                    :picker-options="pickerOptions">
                </el-date-picker>
              </el-form-item>
            </el-form>
          </el-col>

          <el-col :span="2" :style="{ textAlign: 'right' }">
            <div  ><!--v-slot="footer" -->
              <el-button type="success" @click="onSearch">查询</el-button>
            </div>
          </el-col>
          <el-col :span="2" :style="{ textAlign: 'right' }">
            <Create title="添加" btn-type="primary" @emit-connection-list="getDataList"/>
          </el-col>
        </el-row>
      </div>
      <el-table :data="list" tyle="width: 100%" :cell-style="{ textAlign: 'center' }" :header-cell-style="{ 'text-align': 'center' }">
        <el-table-column prop="contesta,contestb" label="比赛" width="180">
          <!-- eslint-disable-next-line-->
          <template v-slot="scope">
            <span>{{scope.row.contesta}}</span>
            <span>&nbsp;VS&nbsp;</span>
            <span>{{scope.row.contestb}}</span>
          </template>
        </el-table-column>

        <el-table-column prop="date" label="日期" width="110"></el-table-column>

        <el-table-column prop="[typea,typeb]" label="让球" width="110">
          <template v-slot="scope">
            <div class="alignment-container">
              <el-space>
                <el-card>
                  <template #header> {{scope.row.typea}} </template>
                  {{scope.row.typeb}}
                </el-card>
              </el-space>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="wina,winb" label="胜利数" width="110">
          <template v-slot="scope">
            <div class="alignment-container">
              <el-space>
                <el-card>
                  <template #header> {{scope.row.wina}} </template>
                  {{scope.row.winb}}
                </el-card>
              </el-space>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="drawa,drawb" label="平局数" width="110">
          <template v-slot="scope">
            <div class="alignment-container">
              <el-space>
                <el-card>
                  <template #header> {{scope.row.drawa}} </template>
                  {{scope.row.drawb}}
                </el-card>
              </el-space>
            </div>
          </template>

        </el-table-column>
        <el-table-column prop="losea,loseb" label="失败数" width="110">
          <template v-slot="scope">
            <div class="alignment-container">
              <el-space>
                <el-card>
                  <template #header> {{scope.row.losea}} </template>
                  {{scope.row.loseb}}
                </el-card>
              </el-space>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作" >
          <template v-slot="scope">
            <el-row :gutter="15">
              <el-col :span="3">
                <Create title="修改" btn-type="primary" :data="scope.row" @emit-connection-list="getDataList"/>
              </el-col>
              <el-col :span="10" :style="{ textAlign: 'left', margin: '0 0 0 4vh' }">
                <el-popconfirm title="确认删除?" @confirm="recordDelete(scope.row.date, scope.row.contesta, scope.row.contestb)">
                  <template #reference>
                    <el-button type="danger" @click.stop>删除</el-button>
                  </template>
                </el-popconfirm>
              </el-col>
            </el-row>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div>
      <el-pagination
          :total="total"
          @current-change="handlePageChange"
          layout="prev, pager, next"
          :pageSize="5"
      >
      </el-pagination>
    </div>
  </main>
</template>

<script setup lang="ts">
import {ref, watch} from "vue";
import Create from "./Create.vue";
import {DeleteData, GetDataByDate, GetDataList, UpdateData} from "../../wailsjs/go/main/App.js";
import { ElTable, ElTableColumn, ElNotification } from "element-plus";
import 'dayjs/locale/zh-cn'
let list = ref([]);

// 定义总条数
let total = ref(0);

// 定义当前页码
const page = ref(1);

// 定义每页显示条数
const size = ref(5);

let props = defineProps(['flush'])
watch(props, (newFlush)=>{
  getDataList()
})

// 定义搜索关键字
const search = ref()
// 定义搜索方法
const onSearch = async () => {
  try {
    console.log("search", search.value, page.value, size.value, total.value)
    const res = await GetDataByDate(search.value, page.value, size.value) // 调用后端搜索数据方法
    list.value = res.data
    total.value = res.total
  } catch (err) {
    ElNotification({
      title:"搜索失败",
      type: "error",
    })
    console.error(err)
  }
}
let form = ref({})
const formLabelWidth = '120px'
const pickerOptions = {
  disabledDate(time) {
    // 获取当前日期
    const today = new Date()
    // 获取一周前的日期
    const oneWeekAgo = new Date(today.getTime() - 7 * 24 * 60 * 60 * 1000)
    // 返回true表示禁用，返回false表示可选
    return time.getTime() < oneWeekAgo.getTime() || time.getTime() > today.getTime()
  },
  shortcuts: [{
    text: '最近一周',
    onClick(picker) {
      const end = new Date();
      const start = new Date();
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
      picker.$emit('pick', [start, end]);
    }
  }, {
    text: '最近一个月',
    onClick(picker) {
      const end = new Date();
      const start = new Date();
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
      picker.$emit('pick', [start, end]);
    }
  }, {
    text: '最近三个月',
    onClick(picker) {
      const end = new Date();
      const start = new Date();
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
      picker.$emit('pick', [start, end]);
    }
  }]
}

// 删除连接
function recordDelete(date, contesta, contestb) {
  DeleteData(date, contesta, contestb ).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
    }
    ElNotification({
      title:res.msg,
      type: "success",
    })
    getDataList()
  })
}

// 处理页码改变
const handlePageChange = (val: number) => {
  page.value = val; // 更新当前页码
  getDataList(); // 重新获取列表数据
};


async function getDataList() {
  try {
    // 判断是否是日期搜索，如果是则调用搜索方法，如果不是则调用获取列表方法
    if (search.value) {
      await onSearch()
      return
    }
    // 使用await等待GetDataList函数的promise返回结果
    let res = await GetDataList(page.value, size.value);
    if (res.code !== 200) {
      ElNotification({ title: res.msg, type: "error" });
    }
    list.value = res.data;
    total.value = res.total;
    console.log(101010, res, res.total, total.value);
  } catch (error) {
    // 处理可能的错误
    console.error(error);
  }
  console.log(12312312312, page.value, size.value, total.value)
}

getDataList()

</script>

<style scoped>

</style>