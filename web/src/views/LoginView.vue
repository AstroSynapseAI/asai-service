<script setup>
import { Form, Field, useForm } from 'vee-validate';
</script>


<template>
  <div class="container">
    <div class="row">

      <div class="col-md-6">
        <img class="logo" src="@/assets/Logo.v1.white.png" alt="">
        <div class="card">
          <div class="card-body">
            
            <Form class="form-control">
              <Field id="Email" name="Email" type="email" class="email-input d-block" placeholder="Email"></Field>
              <Field id="Password" name="Password" type="password" class="pass-input d-block" placeholder="Password"></Field>
              <button class="send-button btn btn-light" @click="''"> Login</button>
          </Form>
            
          </div>
        </div>
      </div>
      
      <div class="col-md-6">
        <h5 class="px-3 mb-4 mt-3 mt-md-0"> Asai cloud is currently in <b>closed beta</b>, and access is limited to <b>invite only</b>. Plese send us your email, if you are interested, and we will add you in the next onbaording batch of testers.</h5>
        <Form class="form-control d-flex" action="https://formspree.io/f/xyyqjdgr" method="POST">
          <Field id="waitlist-email" name="WaitList Email" type="email" class="email-input flex-fill mb-0" placeholder="Email"></Field>
          <button class="send-button btn btn-light" @click="''">Submit</button>
        </Form>
      </div>
    
    </div>
  </div>
</template>

<style scoped>
.logo {
  width: 100%;
  margin-bottom: 3em;
}
.card {
  background-color: black; 
  border: 1px solid white; 
  border-radius: 0;
  width: 100%;
}
.card-body {
  color: white; /* To make text visible in dark background */
}

.form-control {
  background-color: black;
  border: none;
  width: 100%;
}

.email-input, .pass-input {
  width: 100%;
  margin-bottom: 2em;
  height: 3em;
  background-color: transparent;
  border: 1px solid white;
}

.btn {
  background-color: black;
  border-color: white;
  border-radius: 0;
  text-decoration: none;
  color: white;
}

.btn:hover {
  color: black !important;
  background-color: white;
}
</style>