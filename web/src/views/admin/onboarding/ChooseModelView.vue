<script setup>
import { onMounted } from 'vue';


onMounted(() => {
  feather.replace();
})
</script>


<template>
  <div class="container">
    <div class="row">
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5"><h2 class="circle-text"><i data-feather="check"style="width: 36px; height: 36px;"></i></h2></div>
        <h3 class="mb-3">Create Avatar</h3>
        <p class="lead mb-5">Give your AI avatar a name and describe how it should behave.</p>
      </div>
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle current mb-5"><h2 class="circle-text">2</h2></div>
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

      <div class="col-3">
        <div class="card">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
           
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>GPT</h2>
            </div>
            
          </div>
        </div>
      </div>

      <div class="col-3">
        <div class="card">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
           
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>LLama 2</h2>
            </div>
            <div class="card-overlay position-absolute"></div>
          </div>
        </div>
      </div>
      
      <div class="col-3">
        <div class="card">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
           
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>Mistral</h2>
            </div>
            
            <div class="card-overlay position-absolute"></div>
          </div>
        </div>
      </div>

      <div class="col-3">
        <div class="card">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
           
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>Falcon</h2>
            </div>
            
            <div class="card-overlay position-absolute"></div>
          </div>
        </div>
      </div>
    </div>

    <div class="row" style="margin-top: 80px;"> 
      <div class="col-4 offset-2 text-center d-grid">
        <router-link :to="{ name: 'avatar-create' }" class="btn btn-primary btn-lg btn-back">Back</router-link>
      </div>
      <div class="col-4 text-center d-grid">
        <router-link :to="{ name: 'select-agents' }" class="btn btn-primary btn-lg">Next</router-link>
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

.btn-back {
  background-color: transparent;
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

.current {
  background-color: #1c64f2;
}

.circle-text {
   text-align: center;
   line-height: 50px; /* match this with the height of .circle */
   margin: 0;
}

.card-icon {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 85px;
  height: 85px;
  background-color: #374151;
  border-radius: 5px;
}

.card {
  border: 2px solid transparent;
}

.card:hover  {
  border: 2px solid #1c64f2;
  cursor: pointer;
}

.card-overlay {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: rgba(0,0,0,0.5);
  z-index: 10;
}
</style>
