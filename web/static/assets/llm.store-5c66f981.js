import{P as c,Q as e}from"./index-3bc6b851.js";const a="https://test.astrosynapse.ai/api",o=`${a}/llms`,i=c({id:"llm",state:()=>({records:{},record:{}}),actions:{async getLLMs(){try{const r=await e.get(`${o}`);this.records=r}catch(r){console.error(r)}},async getLLM(r){try{const t=await e.get(`${o}/${r}`);this.record=t}catch(t){console.error(t)}},async saveLLM(r){try{const t=await e.post(`${o}/save/active`,r);this.record=t}catch(t){console.error(t)}},async toggleActiveLLM(r,t){try{await e.post(`${o}/${r}/toggle/active`,t)}catch(s){console.error(s)}}}});export{i as u};
