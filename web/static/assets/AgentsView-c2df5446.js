import{u as y}from"./agent.store-d7674f43.js";import{_ as x,u as C,I as h,v as S,o as b,x as o,f as t,y as p,z as g,K as V,c as n,C as f,A as N,h as B,e as R,L,k as q,q as z,s as E}from"./index-a4c4dc87.js";import{u as F}from"./avatar.store-df433578.js";const m=r=>(z("data-v-4468a86a"),r=r(),E(),r),K={class:"container-fluid p-0"},M=m(()=>t("h1",{class:"h3 mb-3"},"Agents",-1)),T={class:"row"},U={class:"col-12"},$={class:"container"},j={class:"card"},G={class:"card-header"},H={class:"row"},J={class:"col"},O={class:"card-title"},P={class:"col-auto"},Q={key:0,class:"form-check form-switch d-flex align-items-center"},W=["checked","onClick"],X=m(()=>t("label",{style:{"margin-bottom":"-5px"},for:"flexSwitchCheckDefault"},"Active",-1)),Y={class:"card-body"},Z={__name:"AgentsView",setup(r){const l=C(),d=y(),A=h(d,"records"),u=F(),_=h(u,"activeAgents"),I=s=>{const e=_.value.find(c=>c.agent.ID===s);return e?e.is_active:!1},k=async s=>{const e=_.value.find(i=>i.agent.ID==s);e&&(e.is_active=!e.is_active);const c={is_active:e?e.is_active:!1,avatar_id:l.avatar.ID};try{await d.toggleActiveAgent(s,c)}catch(i){console.log(i)}},v=s=>{const e=_.value.find(c=>c.agent.ID===s);return e?e.ID:null},D=S(()=>{const s=[],e=[...A.value];for(;e.length;)s.push(e.splice(0,2));return s});return b(async()=>{await d.getAgents();try{await u.getActiveAgents(l.avatar.ID)}catch(s){console.log(s)}feather.replace()}),(s,e)=>{const c=V("router-link");return n(),o("div",K,[M,t("div",T,[t("div",U,[t("div",$,[(n(!0),o(p,null,g(D.value,(i,w)=>(n(),o("div",{key:w,class:"row"},[(n(!0),o(p,null,g(i,a=>(n(),o("div",{class:"col-6",key:a.ID},[t("div",j,[t("div",G,[t("div",H,[t("div",J,[t("h5",O,f(a.name),1)]),t("div",P,[v(a.ID)?(n(),o("div",Q,[t("input",{class:"form-check-input me-2",type:"checkbox",id:"flexSwitchCheckDefault",checked:I(a.ID),onClick:ee=>k(a.ID)},null,8,W),X])):N("",!0)])])]),t("div",Y,[t("p",null,f(a.description),1),t("div",null,[B(c,{to:{name:a.slug,params:{avatar_id:q(l).avatar.ID,agent_id:a.ID,active_agent_id:v(a.ID)}},class:"btn btn-primary"},{default:R(()=>[L("Configure")]),_:2},1032,["to"])])])])]))),128))]))),128))])])])])}}},ce=x(Z,[["__scopeId","data-v-4468a86a"]]);export{ce as default};
