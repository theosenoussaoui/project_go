<template>
  <div class="login">
    <h1>This is a login page</h1>
    <el-row>
      <el-col :span="12" :offset="6">
        <el-card class="box-card">
          <FormWrapper @submit="login">
            <FormField htmlTag="input" type="text" name="email" label="Email"></FormField>
            <FormField htmlTag="input" type="text" name="password" label="Mot de passe"></FormField>
            <el-button type="primary" @click.prevent="login">Valider</el-button>
          </FormWrapper>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import FormWrapper from "../components/Form/FormWrapper.vue";
import FormField from "../components/Form/FormField.vue";
import { auth } from "../services/auth.js";

export default {
  name: "app",
  components: {
    FormWrapper,
    FormField
  },
  methods: {
    login: function() {
    //   alert("ok");
      auth.login("ok", "ook");
    }
  }
};
</script>
