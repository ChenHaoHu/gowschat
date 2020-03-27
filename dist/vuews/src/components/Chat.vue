<template>
  <div>
    <div>
      <el-row :gutter="20" type="flex" justify="center">
        <el-form style="display: flex;justify-content: center;width: 100%">
          <el-col :span="5">
            <el-form-item label="UID">
              <el-input v-model="uid" clearable>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="Name">
              <el-input v-model="name" clearable>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="Gid">
              <el-input v-model="gid" clearable>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item label="LINK">
              <el-button type="primary" @click="link">Link</el-button>
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item label="CLOSE">
              <el-button type="primary" @click="out">OUT</el-button>
              </el-input>
            </el-form-item>
          </el-col>
        </el-form>
      </el-row>
    </div>
    <el-card class="box-card" style="height: 400px;overflow-y:auto;" >
      <div slot="header" class="clearfix">
        <span>聊天窗口 </span>
        <el-tag type="success" v-if="islink">连接</el-tag>
        <el-tag type="danger" v-else="islink">未连接</el-tag>
        <br />
        <span>uid: {{uid}} / name: {{name}} / gid: {{gid}} </span>
      </div>
      <div v-for="o in chatrecord" class="text item" style="text-align: left;font-size: 20px;font-weight:600;">
        <div v-if=" o.type=='this'  ">
          <el-alert :title='"发送信息给"+o.to' type="success"  :closable="false">{{o.msg}}
          </el-alert>
        </div>
        <div v-if="o.from=='NOTICE'&&  o.type=='that'">
          <el-alert :title='"NOTICE发送信息给"+o.to' type="error"  :closable="false" >{{o.msg}}
          </el-alert>
        </div>
  <!--       <div v-if="o.from==uid &&  o.type=='that'">
          <el-alert :title='"发送信息给"+o.to' type="info"  :closable="false">{{o.msg}}
          </el-alert>
        </div> -->
        <div v-if="o.from != 'NOTICE' &&  o.type=='that'">
          <el-alert :title='"接收来自"+o.from+"的信息"' type="warning"  :closable="false">{{o.msg}}
          </el-alert>
        </div>
      </div>
    </el-card>
    <br />
    <el-switch v-model="p2p" active-text="单聊" inactive-text="群发">
    </el-switch>
    <br /> <br />
    <div>
      <el-row :gutter="20" type="flex" justify="center">
        <el-col :span="6" v-if="p2p==1">
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
  props:["uid","gid","name"],
  data() {
    // var uid = "uid1";
    // var name = "name1";
    // var gid = "group1";
    return {
      p2p: 0,
      islink: false,
      socket: "",
      // uid: "uid",
      // name: "name",
      // gid: "gid",
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
      console.log("ws://"+location.hostname+":"+location.port+"/ws/" + this.uid + "," + this.name + "," + this.gid)
      return "ws://"+location.hostname+":"+location.port+"/ws/" + this.uid + "," + this.name + "," + this.gid;
    }
  },
  methods: {
    out(){
      this.socket.close()
    },
    link: function(e) {
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
      this.islink = true
      // const h = this.$createElement;
      this.$message({
        showClose: true,
        message: this.wsuri + 'socket连接成功',
        type: 'success'
      });
      // this.$notify({
      //   title: '连接成功',
      //   message: h('i', { style: 'color: teal' }, this.wsuri + 'socket连接成功')
      // });
    },
    error: function() {
      this.islink = false
      // const h = this.$createElement;
      // this.$notify({
      //   title: '连接失败',
      //   message: h('i', { style: 'color: teal' }, this.wsuri + '连接失败')
      // });
      this.$message({
        showClose: true,
        message: this.wsuri + '连接失败',
        type: 'error'
      });
    },

    getMessage: function(msg) {
      console.log(msg.data)
      var data = JSON.parse(msg.data);
      this.chatrecord.unshift({
        from: data.Uid,
        to: data.ToUid,
        msg: data.Msg,
        type:"that"
      })


    },
    send: function() {

    if(this.islink == true){
        this.chatrecord.unshift({
        from: this.uid,
        to: this.selectChat,
        msg: this.chatmsg,
        type:"this",
      })

      if (this.p2p == 1) {
        var data = {
          touid: this.selectChat,
          msg: this.chatmsg,
          sendtype: 0
        }
      } else {
        var data = {
          touid: this.selectChat,
          msg: this.chatmsg,
          sendtype: 1
        }
      }


      this.socket.send(JSON.stringify(data))
    }else{
      this.$message({
        showClose: true,
        message: '请先连接',
        type: 'error'
      });
    }
    },
    close: function() {
      this.islink = false
      console.log("socket已经关闭")
      // const h = this.$createElement;
      // this.$notify({
      //   title: '连接关闭',
      //   message: h('i', { style: 'color: teal' }, this.wsuri + '连接关闭')
      // });

      this.$message({
        showClose: true,
        message: this.wsuri + '连接关闭',
        type: 'error'
      });
    }
  },
  destroyed() {
    // 销毁监听
    this.islink = false
    this.socket.onclose = this.close
  }
}

</script>
<style scoped>
</style>
