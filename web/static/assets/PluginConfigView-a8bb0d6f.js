import{r as l,o as f,c as g,a as e,J as m,h as k,w as b,v as w,H as P,f as x,g as y}from"./index-92c710bd.js";import{u as C}from"./plugin.store-148f76d7.js";import{u as S}from"./avatar.store-8255a3bd.js";const D={class:"container-fluid p-0"},A={class:"h3 mb-3"},I={class:"form-check form-switch float-end me-5"},T=e("label",{class:"form-check-label",for:"flexSwitchCheckDefault"},"Active",-1),V=["checked"],B={class:"form-check form-switch float-end me-3"},N=e("label",{class:"form-check-label",for:"flexSwitchCheckDefault"},"Public",-1),M=["checked"],R={class:"row"},E={class:"col-12"},H={class:"card"},J={class:"card-body"},U={class:"container"},j={class:"row"},q={class:"col-12"},z={class:"form-floating mb-3"},F=e("label",{for:"floatingInput"},"Token",-1),Q={__name:"PluginConfigView",setup(G){const a=P(),d=x(),n=C(),t=S(),o=l(""),s=l(!1),c=l(!1),r=l(""),_=()=>{s.value=!s.value},v=()=>{c.value=!c.value},p=async()=>{try{await n.saveActivePlugin({ID:parseInt(a.params.active_plugin_id),avatar_id:parseInt(a.params.avatar_id),plugin_id:parseInt(a.params.plugin_id),token:o.value,is_active:s.value,is_public:c.value}),d.push({name:"plugins",params:{avatar_id:a.params.avatar_id}})}catch(i){console.log(i)}};return f(async()=>{try{await n.getPlugin(a.params.plugin_id),r.value=n.record.name,a.params.active_plugin_id&&(await t.getActivePlugin(a.params.avatar_id,a.params.plugin_id),t.activePlugin&&(s.value=t.activePlugin.is_active,c.value=t.activePlugin.is_public,o.value=t.activePlugin.token))}catch(i){console.log(i)}feather.replace()}),(i,u)=>(y(),g("div",D,[e("h1",A,[m("Configure: "+k(r.value)+" ",1),e("div",I,[T,e("input",{class:"form-check-input",type:"checkbox",id:"flexSwitchCheckDefault",checked:s.value,onClick:_},null,8,V)]),e("div",B,[N,e("input",{class:"form-check-input",type:"checkbox",id:"flexSwitchCheckDefault",checked:c.value,onClick:v},null,8,M)])]),e("div",R,[e("div",E,[e("div",H,[e("div",J,[e("div",U,[e("div",j,[e("div",q,[e("div",z,[b(e("input",{"onUpdate:modelValue":u[0]||(u[0]=h=>o.value=h),type:"text",class:"form-control",id:"floatingInput",placeholder:"Token..."},null,512),[[w,o.value]]),F])])]),e("div",{class:"row mt-3"},[e("div",{class:"col-12"},[e("button",{type:"button",class:"btn btn-secondary",onClick:p},"Save")])])])])])])])]))}};export{Q as default};