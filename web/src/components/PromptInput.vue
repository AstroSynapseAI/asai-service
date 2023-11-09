<script setup>
// Tu ide send logika 
import { ref, onMounted } from 'vue';
import { Form, Field, useForm } from 'vee-validate';
import { useChatStore } from '../stores/chat.store.js';

const chatStore = useChatStore();
const prompt = ref("");

let { resetForm } = useForm();

function resizeTextArea(event) {
  const lineHeight = parseInt(window.getComputedStyle(event.target).getPropertyValue("line-height"));
  const currentRows = event.target.scrollHeight / lineHeight;
  const maxRows = 10;

  if (currentRows <= maxRows) {
    event.target.style.height = 'auto';
    event.target.style.height = event.target.scrollHeight + 'px';
  }
}

function submitPrompt(event, resetForm) {
  if (event.shiftKey && event.key == 'Enter') {
    event.preventDefault();
    let cursorPos = event.target.selectionStart;
    let textBeforeCursor = prompt.value.substring(0, cursorPos);
    let textAfterCursor = prompt.value.substring(cursorPos);
    prompt.value = textBeforeCursor + '\n' + textAfterCursor;
  }
  else if (event.key == 'Enter') {
    event.preventDefault();
    if (prompt.value != '') {
      chatStore.sendPrompt(prompt.value);
      resetForm();
    }
  }
}

onMounted(() => {
  feather.replace();
});

</script>

<template>
  <Form v-slot="{ resetForm }">
    <div class="textarea-container">
      <Field
        v-on:input="resizeTextArea"
        @keydown.enter="submitPrompt($event, resetForm)"
        name="prompt"
        v-model="prompt"
        type="text"
        as="textarea"
        class="form-control"
        rows="2"
        placeholder="Send a message..."
        :disabled="chatStore.isLoading"
      ></Field>
      <button class="send-button btn btn-light" @click="submitPrompt($event, resetForm)">
        <i class="align-middle" data-feather="send"></i>
      </button>
    </div>
  </Form>
  
</template>

<style scoped>
.textarea-container {
  position: relative;
}

.message-input {
  padding: 1rem 3rem 1rem 1rem;  /* make space for the button */
}

.send-button {
  position: absolute;
  right: 0.5rem;
  bottom: 0.7rem;
  background-color: black;
  color: white;
  border: 1px solid white;
}

textarea {
  width: 100%;
  background-color: black;
  color: white !important;
  z-index: 999;
}

textarea:focus {
  background-color: black;
  color: white;
}

textarea:disabled {
  background-color: grey;
  color: white;
}

</style>