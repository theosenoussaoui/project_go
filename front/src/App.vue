<template>
  <div>
    <el-row>
      <Navbar></Navbar>
    </el-row>
    <el-container>
      <CardList></CardList>
    </el-container>
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import CardList from "./components/CardList.vue";
export default {
  name: "app",
  components: {
    Navbar,
    CardList
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
body {
  margin: 0;
  padding: 0;
}

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

* {
  font-family: "system-ui" !important;
}
</style>
