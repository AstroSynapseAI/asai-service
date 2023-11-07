<script setup>
// Tu ide send logika 
import { ref, onMounted } from 'vue';
import { Form, Field } from 'vee-validate';
import { useChatStore } from '../stores/chat.store.js';

const chatStore = useChatStore();
const prompt = ref("");
const maxRows = 8;

function resizeTextArea(event) {
  const lineHeight = parseInt(window.getComputedStyle(event.target).getPropertyValue("line-height"));
  const currentRows = event.target.scrollHeight / lineHeight;
  const maxRows = 10;

  if (currentRows <= maxRows) {
    event.target.style.height = 'auto';
    event.target.style.height = event.target.scrollHeight + 'px';
  }

  // Scroll to bottom
  var promptContainerHeight = document.querySelector('.prompt-container').offsetHeight;
  var scrollTo = document.body.scrollHeight + promptContainerHeight;
  window.scrollTo(0, scrollTo);
}

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

onMounted(() => {
  feather.replace();
});

</script>

<template>
  <Form>
    <div class="textarea-container">
      <Field
        v-on:input="resizeTextArea"
        @keydown.enter="submitPrompt"
        name="prompt"
        v-model="prompt"
        type="text"
        as="textarea"
        class="form-control"
        rows="2"
        placeholder="Send a message..."
      ></Field>
      <button class="send-button btn btn-light" @click="submitPrompt">
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
  background-color: transparent;
}

textarea:focus {
  background-color: transparent;
  color: white;
}
</style>