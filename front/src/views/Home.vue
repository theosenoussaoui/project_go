<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "home",
  data: () => ({
    votes: null
  }),
  components: {
    HelloWorld
  },
  beforeMount() {
    this.fetchVotes();
  },
  methods: {
    fetchVotes: async function() {
      const response    = await fetch("http://localhost:8080/votes");
      this.votes        = response;
    }
  }
};
</script>
