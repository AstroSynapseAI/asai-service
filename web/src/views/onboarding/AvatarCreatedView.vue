<script setup>
import { onMounted, ref, toRef } from 'vue';
import { useUserStore } from '@/stores/user.store';

const user = useUserStore();

onMounted( async () => {
  try {
    await user.getUserAvatar(user.current.ID);
  }
  catch (error) {
    console.log(error);
  }
  feather.replace();
})
</script>


<template>
  <div class="container">
    <div class="row mb-5">
      <div class="col-8 text-center offset-2">
        <h1 class="display-6 mb-3">Your Avatar is ready!</h1>
        <h2 class="lead mb-5">You created your AI Avatar. Start using it Today.</h2>
      </div>
    </div>
    <div class="row">
      <div class="col-12 text-center">
        <img src="@/assets/avatar-rdy.png" class="img-fluid mb-5" alt="Avatar created">
      </div>
    </div>

    <!-- <div class="row"> -->
    <!--   <div class="col-4 offset-5"> -->
    <!--      -->
    <!--     <div class="row mb-3" if="avatarCreated"> -->
    <!--       <div class="col-1"> -->
    <!--         <div class="checkmark-icon d-flex align-items-center justify-content-center"> -->
    <!--           <i data-feather="check" style="width: 24px; height: 24px; color: #1c64f2"></i> -->
    <!--         </div> -->
    <!--       </div> -->
    <!--       <div class="col-8"> -->
    <!--         <p class="ml-2">Your AI Avatar is created</p> -->
    <!--       </div> -->
    <!--     </div> -->
    <!---->
    <!--     <div class="row mb-3" if="agentsCreated"> -->
    <!--       <div class="col-1"> -->
    <!--         <div class="checkmark-icon d-flex align-items-center justify-content-center"> -->
    <!--           <i data-feather="check" style="width: 24px; height: 24px; color: #1c64f2"></i> -->
    <!--         </div> -->
    <!--       </div> -->
    <!--       <div class="col-8"> -->
    <!--         <p class="ml-2">Agents are created</p> -->
    <!--       </div> -->
    <!--     </div> -->
    <!--   </div> -->
    <!-- </div> -->
    <div class="row" style="margin-top: 80px;" if="user.avatar.ID">
      <!-- <div class="col-4 offset-2 text-center d-grid"> -->
      <!--   <router-link :to="{name: 'select-agents'}" class="btn btn-primary btn-lg btn-back">Back</router-link> -->
      <!-- </div> -->
      <div class="col-4 text-center offset-4 d-grid">
        <router-link :to="{name: 'admin', params: { avatar_id: user.avatar.ID }}" class="btn btn-primary btn-lg">Continue</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.btn {
  background-color: #1c64f2;
  border-color: #1c64f2;
}

.lead {
  color: grey;
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
