<template>
  <form @submit="submitForm">
    <slot></slot>
    <el-button type="primary" @click.prevent="submitForm">Valider</el-button>
  </form>
</template>

<script>
export default {
  name: "FormWrapper",
  data: function() {
    return {
      values: {}
    };
  },
  provide: function() {
    return {
      setValue: this.setValue
    };
  },
  methods: {
    submitForm: function() {
      this.$emit("onSubmit", this.values);
      console.log(Object.keys(this.values).length);
      alert("form submitted");
    },
    setValue: function(name, value) {
      this.values[name] = value;
    }
  }
};
</script>

<style scoped></style>
