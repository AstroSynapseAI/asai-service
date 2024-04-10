<template>
    <div v-if="isActive" class="error-container px-3">
      <div class="row h-100 align-items-center justify-content-center text-center">
        <div class="col">
          <p class="mb-0">{{ errorText }}</p>
        </div>
        <div class="col-auto">
          <button class="btn-close btn-close-white" @click="closeError"></button>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      isActive: {
        type: Boolean,
        required: true
      },
      errorText: {
      type: String,
      default: ''
    }
    },
    methods: {
      closeError() {
        this.$emit('close');
      }
    }
  };
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
    top: 20px;
    transform: translateX(-50%);
    width: 600px;
    height: 60px;
    background-color: black;
    border-color: white !important;
    color: white !important;
    border: 1px solid white;
    border-radius: 3px;
    z-index: 9999;
  }

  </style>
  