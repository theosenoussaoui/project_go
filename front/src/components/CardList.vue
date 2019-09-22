<template>
  <div>
    <el-row>
      <Card></Card>
    </el-row>
    <el-row>
      <Card></Card>
    </el-row>
    <el-row>
      <Card></Card>
    </el-row>
  </div>
</template>

<script>
import Card from "./Card.vue";
export default {
  name: "app",
  components: {
    Card
  },
  data() {
    return {
      form: {
        name: "",
        region: "",
        date1: "",
        date2: "",
        delivery: false,
        type: [],
        resource: "",
        desc: ""
      }
    };
  },
  methods: {
    onSubmit() {
      console.log("submit!");
    }
  }
};
</script>

<style>
</style>
