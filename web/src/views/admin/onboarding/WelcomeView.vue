<script setup>

</script>

<template>
  <div class="container">
    <div class="row mb-5">
      <div class="col-8 text-center offset-2">
        <h1 class="display-6 mb-3">Welcome to ASAI</h1>
        <h2 class="subtitle mb-5">Get started by creating your AI Avatar</h2>
      </div>
    </div>
    <div class="row">
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5"><h2 class="circle-text">1</h2></div>
        <h3 class="mb-3">Create Avatar</h3>
        <p class="lead mb-5">Give your AI avatar a name and describe how it should behave.</p>
      </div>
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5"><h2 class="circle-text">2</h2></div>
        <h3 class="mb-3">Choose models</h3>
        <p class="lead mb-5">Select one or more LLM models yor avatar will be using.</p>
      </div>
      <div class="col-4 text-center d-flex flex-column align-items-center">
        <div class="circle mb-5"><h2 class="circle-text">3</h2></div>
        <h3 class="mb-3">Select Agents</h3>
        <p class="lead mb-5">Your AI Avatar can browse the internet, answer emails, post on social media nad more!</p>
      </div>
    </div>
    <div class="row">
      <div class="col-4 offset-4 text-center d-grid mx-auto">
        <router-link :to="{ name: 'avatar-create' }" class="btn btn-primary btn-lg">Create your Avatar</router-link>
      </div>
    </div>

  </div>
</template>

<style scoped>
h1, h3 {
  font-weight: 700;
}

.btn {
  background-color: #1c64f2;
  border-color: #1c64f2;
}

.subtitle, .lead {
  color: grey;
}

.circle {
   display: flex;
   justify-content: center;
   align-items: center;
   width: 85px;
   height: 85px;
   border: 2px solid #1c64f2;
   border-radius: 50%;
}

.circle-text {
   text-align: center;
   line-height: 50px; /* match this with the height of .circle */
   margin: 0;
}

</style>
