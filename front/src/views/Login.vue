<template>
  <div class="login">
    <h1>This is a login page</h1>
    <el-row>
      <el-col :span="12" :offset="6">
        <el-card class="box-card">
          <FormWrapper action="">
            <FormField htmlTag="input" type="text" name="email" label="Email"></FormField>
            <FormField htmlTag="input" type="text" name="password" label="Mot de passe"></FormField>
          </FormWrapper>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import FormWrapper from "../components/Form/FormWrapper.vue";
import FormField from "../components/Form/FormField.vue";
export default {
  name: "app",
  components: {
    FormWrapper,
    FormField
  }
};
</script>
