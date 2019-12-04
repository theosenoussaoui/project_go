<template>
    <div>
        <label :for="fieldId">{{ fieldLabel }} : </label>
        <component
            :is="htmlTag"
            :id="fieldId"
            :type="type"
            :name="name"
            :value="value"
            @input="bindField"
        >
            <slot></slot>
        </component>
    </div>
</template>

<script>
    export default {
        name: "FormField",
        props : {
            id: String,
            name: String,
            label: String,
            value: String,
            htmlTag: String,
            type: String,
            multiple: Boolean,
            checked: Boolean
        },
        inject: ['setValue'],
        computed: {
            fieldLabel: function()
            {
                return this.label = (this.label) ? this.label : this.name;
            },
            fieldId: function()
            {
                return this.id = (this.id) ? this.id : Math.floor(Math.random() * 100) + 1 + '-' + Date.now();
            }
        },
        methods: {
            bindField: function(event)
            {
                this.setValue(this.name, event.target.value);
                this.value = event.target.value;
            }
        }
    }
</script>

<style scoped>

</style>