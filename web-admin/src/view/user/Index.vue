<template>
  <div class="card-list">
    <el-card>
      <el-input style="width:440px" @clear="searchUser" clearable v-model="searchForm.name"
                placeholder="请输入用户姓名" class="input-with-select">
        <template #append>
          <el-button icon="Search" @click="searchUser"/>
        </template>
      </el-input>
      <el-table :data="users" border style="width: 100%;margin-top:20px">
        <el-table-column label="ID-头像" min-width="160">
          <template #default="scope">
            <div style="display: flex; align-items: center">
              <el-image :src="scope.row.Avatar"
                        fit="cover" class="small-pic"
              />
              <span style="margin-left: 10px">{{ scope.row.Id }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="UserName" label="账号" min-width="150"/>
        <el-table-column prop="Name" label="昵称" min-width="200"/>
        <el-table-column prop="Email" label="邮件" min-width="150"/>
        <el-table-column prop="Signature" label="签名" min-width="200"/>
        <el-table-column prop="state" label="状态" min-width="150"/>
        <el-table-column prop="Used" label="已使用容量" min-width="180">
          <template #default="scope">
            {{ scope.row.usedStr }} / {{ scope.row.capacityStr }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" type="primary"
                       plain @click="buttonClick(0, scope.row)">恢复
            </el-button>
            <el-button type="danger" size="small"
                       @click="buttonClick(1, scope.row)">封禁
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页 -->
      <el-pagination style="margin-top:20px" :current-page="searchForm.current" :page-size="searchForm.size"
                     :page-sizes="[10, 20, 30, 40]" layout="->,total, sizes, prev, pager, next, jumper"
                     :total="total"
                     @size-change="handleSizeChange" @current-change="handleCurrentChange"/>
    </el-card>
  </div>

  <el-dialog v-model="dialogVisible[0]" title="确认恢复">
    <el-form label-position="top">
      <div>确认恢复用户使用吗？</div>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible[0] = false">关闭</el-button>
                  <el-button @click="setStatus(userOk, 0,false)" type="warning">确定</el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="dialogVisible[1]" title="封禁选项">
    <el-form label-position="top">
      <el-radio-group v-model="radio.value" style="display: flex; flex-direction: column; align-items: baseline">
        <el-radio :value="1" size="large" border style="margin-bottom: 20px">
          头像：
          <el-image :src="selectedUser.Avatar" fit="cover"
                    style="width: 30px; height: 30px; border-radius: 5px;"/>
        </el-radio>
        <el-radio :value="2" size="large" border
                  style="margin-bottom: 20px">
          账号：{{ selectedUser.UserName }}
        </el-radio>
        <el-radio :value="3" size="large" border
                  style="margin-bottom: 20px">
          昵称：{{ selectedUser.Name }}
        </el-radio>
        <el-radio :value="4" size="large" border
                  style="margin-bottom: 20px">
          签名：{{ selectedUser.Signature }}
        </el-radio>
      </el-radio-group>
      <!--            <el-form-item label="封禁理由">-->
      <!--                <el-input v-model="reason"/>-->
      <!--            </el-form-item>-->
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible[1] = false">关闭</el-button>
                  <el-button @click="setStatus(radio, 1,true)" type="warning">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="js" setup>
import userApi from "../../api/user";
import {onMounted, reactive, ref} from "vue";
import {ElMessage, ElMessageBox} from 'element-plus';
import {useRouter} from 'vue-router'
import {formatSize, formatState} from "@/utils/util.js";
import {userMap, userOk} from "@/utils/constant.js";
import {codeOk, promptError, promptSuccess} from "@/utils/http/base.js";

const router = useRouter();

// 用户数据
let users = ref([]),
    total = ref(0),
    radio = ref(0),
    selectedUser

// 搜索条件
const searchForm = reactive({
      current: 1,
      size: 10,
      name: ''
    }),
    dialogVisible = reactive([false, false])

// 获取用户列表
const getUserList = async () => {
  const res = await userApi.getUserList({'page': 0, 'size': 100})
  users.value = res.data.data
  users.value.forEach(user => {
    user.usedStr = formatSize(user.Used)
    user.capacityStr = formatSize(user.Capacity)
    user.state = userMap[user.Status]
  })
  // total.value = res.data.data.total;
}

function buttonClick(option, user) {
  selectedUser = user
  dialogVisible[option] = true
}

async function setStatus( radio, idx,status) {
  const statusValue = Number(status);
  if (radio) {
    if (radio === 0) {
      promptError('请先选择')
      return
    }
  }
  const resp = await userApi.setStatus(selectedUser.Id, statusValue)
  if (resp.data.code === codeOk) {
    await getUserList()
    promptSuccess(`设置成功`)
    console.log(idx)
    dialogVisible[idx] = false
  } else {
    promptError(`设置失败，${resp.data.msg}`)
  }
}

const handleSizeChange = (size) => {
  searchForm.size = size;
  getUserList();
}
const handleCurrentChange = (current) => {
  searchForm.current = current;
  getUserList();
}
const searchUser = () => {
  searchForm.current = 1;
  getUserList();
}
// 删除用户
const deleteUser = (id) => {
  ElMessageBox.confirm(
      '确定要删除该用户信息吗?',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(async () => {
    const res = await userApi.delUser({id: id});
    if (res.data.success) {
      ElMessage.success("删除成功")
     await getUserList();
    } else {
      ElMessage.error("删除失败")
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '取消删除',
    })
  })
}

onMounted(() => {
  getUserList();
})

</script>

<style scoped>
.small-pic {
  width: 50px;
  height: 50px;
  border-radius: 5px;
}
</style>