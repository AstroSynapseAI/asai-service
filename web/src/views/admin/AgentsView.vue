<script setup>
import { onMounted, toRef } from 'vue';
import { useAgentStore } from '@/stores/agent.store';


const avatarName = "asai";
const agent = useAgentStore();
const allAgents = toRef(agent, 'allAgents');

onMounted(async () => {
  await agent.getAgents();
  feather.replace();
});

</script>

<template>
          
  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Agents</h1>
    <div class="row">
      <div class="col-12">
        <div class="container">

          <div class="row" v-for="(agent, index) in allAgents.filter((a, i) => i % 2 === 0)" :key="'row' + index">
              <!-- Render the current and next agent (if it exists) within the same row -->
              <div class="col-6" v-for="activeAgent in allAgents.slice(index, index + 2)" :key="activeAgent.ID">
                <div class="card">
                  <div class="card-header">
                    <h5 class="card-title">{{ activeAgent.name }}</h5>
                    <div class="form-check form-switch float-end">
                      <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                      <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" v-model="activeAgent.isActive">
                    </div>
                  </div>
                  <div class="card-body">
                    <p>{{ activeAgent.description }}</p>
                    <div>
                      <router-link :to="{name: 'agent-config', params: { avatar: activeAgent.name, slug: activeAgent.slug }}" class="btn btn-primary">Configure</router-link>
                    </div>
                  </div>
                </div>
              </div>
          </div>

          
          <!-- <div class="row">
            <div class="col-6">
              <div class="card">
                <div class="card-header">
                  <div class="row">
                    <div class="col-6">
                      <h5 class="card-title">Search Agent</h5>
                    </div>
                    <div class="col-6">
                      <div class="form-check form-switch float-end">
                        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                      </div>
                    </div>
                  </div>

                  
                </div>
                <div class="card-body">
                  <p>Utilizes search engines such as Google, DuckDuckGo, and Metaphor for automated web searches.</p>
                  <div>
                    <router-link :to="{name: 'agent-config', params: { avatar: avatarName, slug: 'search' }}" class="btn btn-primary">Configure</router-link>
                  </div>
                </div>
              </div>
            </div>


            <div class="col-6">
              <div class="card">
                <div class="card-header">

                  <div class="row">
                    <div class="col-6">
                      <h5 class="card-title">Browser Agent</h5>
                    </div>
                    <div class="col-6">
                      <div class="form-check form-switch float-end">
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                      </div>
                    </div>
                  </div>

                </div>
                <div class="card-body">
                  <p>Equipped with the capability to scrape, read website contents, and interact with web pages and web applications.</p>
                  <div>
                    <router-link :to="{name: 'agent-config', params: { avatar: avatarName, slug: 'browser' }}" class="btn btn-primary">Configure</router-link>
                  </div>
                </div>
              </div>
            </div>
          </div> -->


        </div>
      </div>
    </div>
  </div>
      
</template>