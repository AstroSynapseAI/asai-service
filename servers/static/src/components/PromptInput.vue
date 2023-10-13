<script setup>
// Tu ide send logika 
import { ref } from 'vue';
import { Form, Field } from 'vee-validate';
import { useChatStore } from '../stores/chat.store.js';

const chatStore = useChatStore();
const prompt = ref("");

function submitPrompt() {
  chatStore.sendPrompt(prompt.value);
  prompt.value = "";
}

</script>

<template>
  <Form>
    <Field 
      @keyup.enter="submitPrompt" 
      name="prompt" 
      v-model="prompt" 
      type="text" 
      as="textarea" 
      class="form-control" 
      rows="3" 
      placeholder="Send a message..."
    ></Field>
  </Form>
  
</template>

<style scoped>
textarea {
  width: 100%;
  background-color: transparent;
}

textarea:focus {
  background-color: transparent;
  color: white;
}
</style>