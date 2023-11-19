<script setup>
import { ref, onMounted } from 'vue';
import { Form, Field, useForm } from 'vee-validate';
import { useChatStore } from '../stores/chat.store.js';

const chatStore = useChatStore();
const prompt = ref("");

let { resetForm } = useForm();

function resizeTextArea(event) {
  event.target.style.height = 'auto';

  const lineHeight = parseInt(window.getComputedStyle(event.target).getPropertyValue("line-height"));
  let currentRows = Math.floor(event.target.scrollHeight / lineHeight);
  const maxRows = 10;

  currentRows = currentRows > maxRows ? maxRows : currentRows;
  event.target.style.height = `${currentRows * lineHeight}px`;
}

function submitPrompt(event, resetForm) {
  event.preventDefault();
  if (event.shiftKey && event.key == 'Enter') {
    let cursorPos = event.target.selectionStart;
    let textBeforeCursor = prompt.value.substring(0, cursorPos);
    let textAfterCursor = prompt.value.substring(cursorPos);
    prompt.value = textBeforeCursor + '\n' + textAfterCursor;
    resizeTextArea(event);
  }
  else if (event.key == 'Enter') {
    if (prompt.value.trim() !== '') {
      chatStore.sendPrompt(prompt.value);
      event.target.style.height = 'auto';
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
    <div :class="{ 'textarea-container': true, 'loading': chatStore.isLoading, 'form-control': true}" class="form-control">
      <Field
        v-on:input="resizeTextArea"
        @keydown.enter="submitPrompt($event, resetForm)"
        name="prompt"
        v-model="prompt"
        type="text"
        as="textarea"
        class=""
        rows="2"
        placeholder="Send a message..."
        :disabled="chatStore.isLoading"
      ></Field>
      <button class="send-button btn btn-light" :disabled="chatStore.isLoading" @click="submitPrompt($event, resetForm)">
        <i class="align-middle" data-feather="send"></i>
      </button>
    </div>
  </Form>
  
</template>

<style scoped>
.textarea-container {
  position: relative;
  background-color: black;
}

.loading {
  background-color: rgba(255, 255, 255, 0.1);
}

.send-button {
  position: absolute;
  right: 0.5rem;
  bottom: 0.7rem;
  background-color: black;
  color: white;
  border: 1px solid white;
}

.send-button:disabled {
  background-color: rgba(255, 255, 255, 0.1);
  color: black;
}

textarea {
  width: calc(100% - 50px);
  background-color: transparent;
  color: white !important;
  z-index: 999;
  border: none;
}

textarea:focus {
  background-color: transparent;
  color: white;
  border: none !important;
  outline: none !important;
}

textarea:disabled {
  color: white;
}

</style>