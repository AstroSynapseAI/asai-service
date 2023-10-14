<script setup>
// Tu ide send logika 
import { ref } from 'vue';
import { Form, Field } from 'vee-validate';
import { useChatStore } from '../stores/chat.store.js';

const chatStore = useChatStore();
const prompt = ref("");

function submitPrompt(event) {
  if (event.shiftKey && event.key == 'Enter') {
    event.preventDefault();
    let cursorPos = event.target.selectionStart;
    let textBeforeCursor = prompt.value.substring(0, cursorPos);
    let textAfterCursor = prompt.value.substring(cursorPos);
    prompt.value = textBeforeCursor + '\n' + textAfterCursor;
  }
  else if (event.key == 'Enter') {
    event.preventDefault();
    chatStore.sendPrompt(prompt.value);
    prompt.value = "";
  }
}

</script>

<template>
  <Form>
    <Field 
      @keydown.enter="submitPrompt"
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