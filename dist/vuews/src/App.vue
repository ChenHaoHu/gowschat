<template>
  <div id="app">
    <el-row :gutter="20" type="flex" justify="center">
      <el-col :span="9">
        <chat :uid="100" :gid="100" :name="111" />
      </el-col>
      <el-col :span="9">
        <chat  :uid="101" :gid="101" :name="222" />
      </el-col>
      <el-col :span="9">
        <chat  :uid="102" :gid="100" :name="333"/>
      </el-col>
    </el-row>
    <br>
    <br>
    <el-button type="danger" @click="dialogFormVisible = true">发送全局通知</el-button>
    <el-dialog title="全局通知" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="发送内容">
          <el-input v-model="form.msg" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="发送类型">
          <el-select v-model="form.type" placeholder="请选择活动区域">
            <el-option label="应用内全局通知" value="3"></el-option>
            <el-option label="群内全局通知" value="2"></el-option>
            <el-option label="群内单点通知" value="4"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Uid" v-if="form.type==4">
          <el-input v-model="form.uid"></el-input>
        </el-form-item>
        <el-form-item label="Gid" v-if="form.type!=3">
          <el-input v-model="form.gid"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="sendNotice">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import chat from './components/chat'
import axios from 'axios'

export default {
  name: 'App',
  data() {
    return {
      dialogFormVisible: false,
      form: {
        type: "2"
      }

    }
  },
  components: {
    chat
  },
  methods: {
    sendNotice() {
      this.dialogFormVisible = false

      let params = new FormData();

      params.append('sendtype', parseInt(this.form.type));
      params.append('msg', this.form.msg);
      params.append('gid', this.form.gid);
      params.append('uid', this.form.uid);

      axios.post('/chat/notice', params).then(function(response) {

          this.$message({
            showClose: true,
            message: '发送成功',
            type: 'success'
          });

        })

        .catch(function(error) {
          this.$message({
            showClose: true,
            message: '发送出错',
            type: 'error'
          });

        })
    }
  }
}

</script>
<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

</style>
