<template>

  <div>
    <textarea v-model="memo.content"></textarea>
    <input @click="onSubmit(memo)" type="button" value="添加">
    <li v-for="item in memos.list">
      <input v-bind:value="item.content" ><a @click="deleteMemo(item.id)" methods="delete" >删除</a>
    </li>
  </div>
</template>
<script>
import memoApi from "../../api/modules/memo";

export default {
  data() {
    return {memos: {list: [], total: 0}, memo: {content: ""}}
  }
  , props: {}, methods: {
    onSubmit: function (memo) {
      memoApi.addMemo(memo).then((v) => {
        this.memos.list.unshift(v.data)
        return v
      })
    },
    deleteMemo: function (id) {
      memoApi.deleteMemo(id).then((v) => {
              this.memos.list.forEach(function (item,index,arr){
                  if (item.id ==id) {
                      arr.splice(index,1);
                  }
              });

        return v
      })
    }
  },
  setup(props) {
    let memos = reactive({list: "", total: ""})
    onMounted(() => {
      memoApi.listMemo().then((v) => {
        memos.list = v.data.list
        memos.total = v.data.total
      })
    })
    return {memos, onMounted}
  },
}
</script>


<style scoped>

</style>