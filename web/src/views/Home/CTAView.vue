<template>
  <h4 class="text-center"> Asai comes in 3 differnte flavors, make your choice: </h4>
  <div class="cta row">
    <div class="col-md-4">
      <button class="btn btn-primary w-100 btn-lg">CLOUD</button>
    </div>
    <div class="col-md-4">
      <button class="btn btn-primary w-100 btn-lg">SELF-HOSTED</button>
    </div>
    <div class="col-md-4">
      <button class="btn btn-primary w-100 btn-lg">LOCAL</button>
    </div>
  </div>
</template>

<style scoped>
.cta {
  margin-top: 100px;  
  margin-bottom: 100px;
}
.cta .btn {
  background-color: black;
  border-color: white;
}

.btn:hover {
  color: black;
  background-color: white;
}

@media only screen and (max-width: 576px) {
  .cta {
    margin-top: 25px;
    margin-bottom: 25px;
  }
  .cta .btn {
    margin-bottom: 15px;
  }
}
</style>