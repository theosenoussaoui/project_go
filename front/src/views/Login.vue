<template>
    <div class="login">
        <h1>This is a login page</h1>
        <el-row>
            <el-col :span="12" :offset="6">
                <el-card class="box-card">
                    <FormWrapper @submit.prevent="handleSubmit">
                        <FormField
                            htmlTag="input"
                            type="text"
                            name="email"
                            label="Email"
                            v-model="email"
                        ></FormField>
                        <FormField
                            htmlTag="input"
                            type="text"
                            name="password"
                            label="Mot de passe"
                            v-model="password"
                        ></FormField>
                        <el-button type="primary" :disabled="status.loggingIn">Valider</el-button>
                    </FormWrapper>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import FormWrapper from "../components/Form/FormWrapper.vue";
import FormField from "../components/Form/FormField.vue";
import { mapState, mapActions } from "vuex";

export default {
    components: {
        FormWrapper,
        FormField
    },
    data() {
        return {
            username: "",
            password: "",
            submitted: false
        };
    },
    computed: {
        ...mapState("account", ["status"])
    },
    created() {
        // reset login status
        this.logout();
    },
    methods: {
        ...mapActions("account", ["login", "logout"]),
        handleSubmit(e) {
            this.submitted = true;
            const { username, password } = this;
            if (username && password) {
                this.login({ username, password });
            }
        }
    }
};
</script>
