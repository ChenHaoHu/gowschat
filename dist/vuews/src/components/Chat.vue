<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>聊天信息 {{uid}} / {{name}} / {{gid}} </span>
        <el-row :gutter="20" type="flex" justify="center">
          <el-col :span="6">
            <el-input v-model="uid" clearable>
            </el-input>
          </el-col>
          <el-col :span="6">
            <el-input v-model="name" clearable>
            </el-input>
          </el-col>
          <el-col :span="6">
            <el-input v-model="gid" clearable>
            </el-input>
          </el-col>
          <el-col :span="3">
            <el-button type="primary" @click="init">Link</el-button>
            </el-input>
          </el-col>
        </el-row>
      </div>
      <div v-for="o in chatrecord" class="text item" style="text-align: left;font-size: 20px;font-weight:600;">
        <div v-if="o.from != uid && o.to != uid && o.from != 'NOTICE'">
          <el-alert :title='o.from+"发送信息给"+o.to' type="success">{{o.msg}}
          </el-alert>
        </div>
        <div v-if="o.from=='NOTICE'">
          <el-alert :title='"发送信息给"+o.to' type="error">{{o.msg}}
          </el-alert>
        </div>
        <div v-if="o.from==uid">
          <el-alert :title='"发送信息给"+o.to' type="info">{{o.msg}}
          </el-alert>
        </div>
        <div v-if="o.to == uid && o.from != 'NOTICE'">
          <el-alert :title='"接收来自"+o.from+"的信息"' type="warning">{{o.msg}}
          </el-alert>
        </div>
      </div>
    </el-card>
    <br>
    <div>
      <el-row :gutter="20" type="flex" justify="center">
        <el-col :span="6">
          <el-input placeholder="张三？" v-model="selectChat" clearable>
          </el-input>
        </el-col>
        <el-col :span="12">
          <el-input placeholder="请输入内容" v-model="chatmsg" clearable>
          </el-input>
        </el-col>
        <el-col :span="3">
          <el-button type="primary" @click="send">发送</el-button>
          </el-input>
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script>
export default {
  name: 'chat',
  data() {
    var uid = "112";
    var name = "hcy";
    var gid = "2121";
    return {
      socket: "",
      uid: uid,
      name: name,
      gid: gid,
      chatlist: [{
        value: '选项1',
        label: '黄金糕'
      }, {
        value: '选项2',
        label: '双皮奶'
      }],
      selectChat: "",
      chatmsg: "",
      chatrecord: []
    }
  },
  computed: {
    wsuri: function() {
      return "ws://127.0.0.1:8081/ws/" + this.uid + "," + this.name + "," + this.gid;
    }
  },
  methods: {
    init: function(e) {
      if (typeof(WebSocket) === "undefined") {
        alert("您的浏览器不支持socket")
      } else {
        if (this.socket != "") {
          this.socket.close()
        }
        // 实例化socket
        this.socket = new WebSocket(this.wsuri)
        // 监听socket连接
        this.socket.onopen = this.open
        // 监听socket错误信息
        this.socket.onerror = this.error
        this.socket.onclose = this.close
        // 监听socket消息
        this.socket.onmessage = this.getMessage
      }
    },
    open: function() {
      console.log("socket连接成功")
      const h = this.$createElement;
      this.$notify({
        title: '连接成功',
        message: h('i', { style: 'color: teal' }, this.wsuri + 'socket连接成功')
      });
    },
    error: function() {
      const h = this.$createElement;
      this.$notify({
        title: '连接失败',
        message: h('i', { style: 'color: teal' }, this.wsuri + '连接失败')
      });
    },

    getMessage: function(msg) {
      console.log(msg.data)
      var data = JSON.parse(msg.data);
      this.chatrecord.push({
        from: data.Uid,
        to: data.ToUid,
        msg: data.Msg
      })


    },
    send: function() {

      this.chatrecord.push({
        from: this.uid,
        to: this.selectChat,
        msg: this.chatmsg
      })

      var data = {
        touid: this.selectChat,
        msg: this.chatmsg,
        sendtype: 0
      }


      this.socket.send(JSON.stringify(data))
    },
    close: function() {
      console.log("socket已经关闭")
      const h = this.$createElement;
      this.$notify({
        title: '连接关闭',
        message: h('i', { style: 'color: teal' }, this.wsuri + '连接关闭')
      });
    }
  },
  destroyed() {
    // 销毁监听
    this.socket.onclose = this.close
  }
}

</script>
<style scoped>
</style>
