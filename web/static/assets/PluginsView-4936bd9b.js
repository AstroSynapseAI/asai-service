import{u as y}from"./plugin.store-8d9f8004.js";import{_ as P,j as x,t as g,o as S,J as b,a as o,c as i,b as e,F as p,e as m,h as D,E as C,f as A,n as V,K as N,g as B,y as E,z as F}from"./index-6b44ca5b.js";import{u as R}from"./avatar.store-064ebf4e.js";const I=l=>(E("data-v-084fee13"),l=l(),F(),l),j={class:"container-fluid p-0"},z=I(()=>e("h1",{class:"h3 mb-3"},"Plugins",-1)),J={class:"row"},K={class:"col-12"},L={class:"container"},M={class:"row"},T={class:"card"},U={class:"card-header"},$={class:"row"},q={class:"col"},G={class:"card-title"},H={class:"col-auto"},O={key:0,class:"form-check form-switch d-flex align-items-center"},Q=["checked","onClick"],W=I(()=>e("label",{style:{"margin-bottom":"-5px"},for:"flexSwitchCheckDefault"},"Active",-1)),X={class:"card-body"},Y={__name:"PluginsView",setup(l){const _=x(),d=y(),v=g(d,"records"),h=R(),u=g(h,"activePlugins"),k=s=>{const t=u.value.find(c=>c.plugin.ID==s);return t?t.is_active:!1},w=async s=>{const t=u.value.find(n=>n.plugin.ID==s);t&&(t.is_active=!t.is_active);const c={is_active:t?t.is_active:!1,avatar_id:_.avatar.ID};try{await d.toggleActivePlugin(s,c)}catch(n){console.log(n)}},f=s=>{const t=u.value.find(c=>c.plugin.ID===s);return t?t.ID:null};return S(async()=>{try{await d.getPlugins(),await h.getActivePlugins(_.avatar.ID)}catch(s){console.log(s)}feather.replace()}),(s,t)=>{const c=b("router-link");return o(),i("div",j,[z,e("div",J,[e("div",K,[e("div",L,[e("div",M,[(o(!0),i(p,null,m(v.value.filter((n,r)=>r%2===0),(n,r)=>(o(),i("div",{class:"row",key:"row"+r},[(o(!0),i(p,null,m(v.value.slice(r,r+2),a=>(o(),i("div",{class:"col-6",key:a.ID},[e("div",T,[e("div",U,[e("div",$,[e("div",q,[e("h5",G,D(a.name),1)]),e("div",H,[f(a.ID)?(o(),i("div",O,[e("input",{class:"form-check-input me-2",type:"checkbox",id:"flexSwitchCheckDefault",checked:k(a.ID),onClick:Z=>w(a.ID)},null,8,Q),W])):C("",!0)])])]),e("div",X,[e("p",null,D(a.description),1),e("div",null,[A(c,{to:{name:"plugin-config",params:{avatar_id:B(_).avatar.ID,plugin_id:a.ID,active_plugin_id:f(a.ID)}},class:"btn btn-primary"},{default:V(()=>[N(" Configure ")]),_:2},1032,["to"])])])])]))),128))]))),128))])])])])])}}},ae=P(Y,[["__scopeId","data-v-084fee13"]]);export{ae as default};
