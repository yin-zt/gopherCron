var d=Object.defineProperty,l=Object.defineProperties;var u=Object.getOwnPropertyDescriptors;var a=Object.getOwnPropertySymbols;var h=Object.prototype.hasOwnProperty,m=Object.prototype.propertyIsEnumerable;var n=(e,t,o)=>t in e?d(e,t,{enumerable:!0,configurable:!0,writable:!0,value:o}):e[t]=o,c=(e,t)=>{for(var o in t||(t={}))h.call(t,o)&&n(e,o,t[o]);if(a)for(var o of a(t))m.call(t,o)&&n(e,o,t[o]);return e},p=(e,t)=>l(e,u(t));import{a as v,ck as f,cl as _,c as C,h as g,b as q,a1 as y,S as k}from"./index.e18600f8.js";import{g as A,f as E}from"./QDialog.07d9a48e.js";var b=v({name:"QCardActions",props:p(c({},f),{vertical:Boolean}),setup(e,{slots:t}){const o=_(e),r=C(()=>`q-card__actions ${o.value} q-card__actions--${e.vertical===!0?"vert column":"horiz row"}`);return()=>g("div",{class:r.value},q(t.default))}});function i(e){if(e===!1)return 0;if(e===!0||e===void 0)return 1;const t=parseInt(e,10);return isNaN(t)?0:t}var x=y({name:"close-popup",beforeMount(e,{value:t}){const o={depth:i(t),handler(r){o.depth!==0&&setTimeout(()=>{const s=A(e);s!==void 0&&E(s,r,o.depth)})},handlerKey(r){k(r,13)===!0&&o.handler(r)}};e.__qclosepopup=o,e.addEventListener("click",o.handler),e.addEventListener("keyup",o.handlerKey)},updated(e,{value:t,oldValue:o}){t!==o&&(e.__qclosepopup.depth=i(t))},beforeUnmount(e){const t=e.__qclosepopup;e.removeEventListener("click",t.handler),e.removeEventListener("keyup",t.handlerKey),delete e.__qclosepopup}});export{x as C,b as Q};