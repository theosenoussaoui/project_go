<template>
    <form @submit="submitForm">
        <slot></slot>
        <button @click.prevent="submitForm">Valider</button>
    </form>
</template>

<script>
    export default {
        name: "FormWrapper",
        data: function()
        {
            return {
                values: {}
            };
        },
        provide: function()
        {
            return {
                setValue: this.setValue
            }
        },
        methods: {
            submitForm: function()
            {
                this.$emit('onSubmit', this.values);
                alert('form submitted');
            },
            setValue: function(name, value)
            {
                this.values[name] = value;
            }
        }
    }
</script>

<style scoped>

</style>