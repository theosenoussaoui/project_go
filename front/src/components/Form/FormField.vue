<template>
  <div class="fromfield-wrapper">
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
  props: {
    id: String,
    name: String,
    label: String,
    value: String,
    htmlTag: String,
    type: String,
    multiple: Boolean,
    checked: Boolean
  },
  inject: ["setValue"],
  computed: {
    fieldLabel: function() {
      const computedLabel = (this.getLabel() != null ? this.getLabel() : this.name);
      this.setLabel(computedLabel);
      return computedLabel;
    },
    fieldId: function() {
      const computedId = (this.getId() != null ? this.getId() : Math.floor(Math.random() * 100) + 1 + "-" + Date.now());
      this.setId(computedId);
      return computedId;
    }
  },
  methods: {
    getLabel: function() {
      return this.label;
    },
    setLabel: function(label) {
      this.label = label;
    },
    getId: function() {
      return this.id;
    },
    setId: function(id) {
      this.id = id;
    },
    bindField: function(event) {
      this.setValue(this.name, event.target.value);
      this.value = event.target.value;
    }
  }
};
</script>

<style scoped>
.fromfield-wrapper {
  margin-bottom: 20px;
}

input {
  -webkit-appearance: none;
  background-color: #fff;
  background-image: none;
  border-radius: 4px;
  border: 1px solid #dcdfe6;
  box-sizing: border-box;
  color: #606266;
  display: inline-block;
  font-size: inherit;
  height: 40px;
  line-height: 40px;
  outline: none;
  padding: 0 15px;
  transition: border-color .2s cubic-bezier(.645,.045,.355,1);
  width: 50%;
}
</style>
