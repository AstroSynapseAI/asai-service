<script setup>
import { onMounted } from 'vue';

onMounted(() => {
  feather.replace();
})
</script>


<template>
  <div>
    <div class="row mb-5">
      <div class="col-8 text-center offset-2">
        <h1 class="display-6 mb-3">Your Avatar is getting ready</h1>
        <h2 class="subtitle mb-5">Your agents are being created and your avatar will be ready in a few moments.</h2>
      </div>
    </div>
    <div class="row">
      <div class="col-12 text-center">
        <img src="@/assets/avatar-rdy.svg" class="img-fluid mb-5" alt="Avatar created">
      </div>
    </div>

    <div class="row">
      <div class="col-4 offset-5">
        
        <div class="row mb-3">
          <div class="col-1">
            <div class="checkmark-icon d-flex align-items-center justify-content-center">
              <i data-feather="check" style="width: 24px; height: 24px; color: #1c64f2"></i>
            </div>
          </div>
          <div class="col-8">
            <p class="ml-2">Search agent is created</p>
          </div>
        </div>

        <div class="row mb-3">
          <div class="col-1">
            <div class="checkmark-icon d-flex align-items-center justify-content-center">
              <i data-feather="check" style="width: 24px; height: 24px; color: #1c64f2"></i>
            </div>
          </div>
          <div class="col-8">
            <p class="ml-2">Email agent is created</p>
          </div>
        </div>
      </div>
    </div>
    <div class="row" style="margin-top: 80px;">
      <div class="col-4 offset-2 text-center d-grid">
        <router-link :to="{name: 'select-agents'}" class="btn btn-primary btn-lg btn-back">Back</router-link>
      </div>
      <div class="col-4 text-center d-grid">
        <router-link :to="{name: 'select-agents'}" class="btn btn-primary btn-lg">Continue</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.btn {
  background-color: #1c64f2;
  border-color: #1c64f2;
}

.btn-back {
  background-color: transparent;
}

.checkmark-icon {
  background-color: white;
  border-radius: 50%;
  padding: 5px;
  height: 24px;
  width: 24px;
  display: flex; 
  align-items: center;
  justify-content: center;
}

</style>
