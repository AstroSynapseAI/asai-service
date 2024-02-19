<script setup>
import { ref, onMounted, toRef } from 'vue'
import { useUserStore } from '@/stores/user.store';
import { useAuthStore } from '@/stores/auth.store';

const auth = useAuthStore();

const user = useUserStore();
const usersRecords = toRef(user, 'records')

const showModal = ref(false);
const inviteUsername = ref("");
const inviteToken = ref("");
const hostname = ref(import.meta.env.VITE_API_URL)

const avatarName = (roles) => {
  if (roles) {
    const ownerRole = roles.find(ownerRole => ownerRole.role.permission === 'owner')
    if (ownerRole) {
      return ownerRole.avatar.name
    }
    return ""
  }

  return ""
}

const inviteUser = async () => {
  try {
    const user = await auth.inviteUser(inviteUsername.value)
    if (user) {
      inviteToken.value = user.invite_token
    }
  }
  catch (error) {
    console.error(error)
  }
}

onMounted( async () => {
  try {
    await user.getUsers()
  } catch (error) {
    console.error(error)
  }

  feather.replace()
})
</script>
<template>
  <div class="container">

    <!-- Adjusted modal class bindings for centering with sidebar presence -->
    <div class="modal" :class="{ 'd-block': showModal, 'show': showModal }">
      <div class="modal-dialog">
        
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Invite New User</h5>
            <button type="button" class="btn-close" @click="showModal = false"></button>
          </div>
          
          <div class="modal-body container">
            <div class="row">
              <div class="col-12">
                <div class="form-floating mb-3">
                  <input v-model="inviteUsername" type="text" class="form-control" id="floatingInput" placeholder="Username">
                  <label for="floatingInput">Username</label>
                </div>
                <p>Invite Link: </p>
                <p><pre>{{ hostname }}/register/{{ inviteToken }}</pre></p>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" @click="showModal = false">Close</button>
            <button type="button" class="btn btn-primary" @click="inviteUser()">Invite</button>
          </div>
        </div>

      </div>
    </div>
  
    <h1>Users <button class="btn btn-primary float-end" @click="showModal = true">Invite User</button></h1>

    <div class="table-responsive">
      
      <table class="table table-striped">
        <thead>
          <tr>
            <th>Username</th>
            <th>Avatar</th>
            <th>Last Login</th>
            <th>Invited on</th>
            <th>Invite Link</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in usersRecords" :key="user.id">
            <td>{{ user.username }}</td>
            <td>{{ avatarName(user.roles) }}</td>
            <td>{{ new Date().toLocaleDateString() }}</td>
            <td>{{ new Date().toLocaleDateString() }}</td>
            <td><pre>{{ hostname }}/register/{{ user.invite_token }}</pre></td>
            <td>Delete</td>
          </tr>
        </tbody>
      </table>
      
    </div>
  </div>
</template>


<style scoped>

pre {
  white-space: pre-wrap;
  word-wrap: break-word;
}

.table-responsive {
    overflow-x: auto;
}

.table {
    width: 100%;
    table-layout: fixed;
    word-wrap: break-word;
}

td pre {
    word-wrap: break-word;
    max-width: 100%;
}

</style>