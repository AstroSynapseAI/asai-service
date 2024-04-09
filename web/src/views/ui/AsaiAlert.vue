<template>
    <div v-if="connectionErr.active" class="error-container text-white border border-1 border-white px-3">
      <div class="row h-100 align-items-center justify-content-center text-center">
        <div class="col">
          <p class="mb-0">An error has occurred! Please refresh the page and try again.</p>
        </div>
        <div class="col-auto">
          <button class="btn-close btn-close-white" @click="chatStore.closeError"></button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        connectionErr: {
          active: true // Set to true if error is active
        }
      };
    },
    methods: {
      closeError() {
        this.connectionErr.active = false;
      }
    }
  }
  </script>

<style scoped>
@font-face {
  font-family: 'SF Pro Text';
  src: url('./assets/fonts/SFProText-Regular.woff2') format('woff2');
  font-weight: 400;
  font-style: normal;
}

@font-face {
  font-family: 'SF Pro Text';
  src: url('./assets/fonts/SFProText-Bold.woff2') format('woff2');
  font-weight: 700;
  font-style: normal;
}

.error-container {
  position: fixed;
  left: 50%;
  transform: translateX(-50%);
  width: 600px;
  height: 60px;
  background-color: black;
}

</style>
