var ie=Object.defineProperty,le=Object.defineProperties;var se=Object.getOwnPropertyDescriptors;var H=Object.getOwnPropertySymbols;var re=Object.prototype.hasOwnProperty,ue=Object.prototype.propertyIsEnumerable;var A=(e,o,a)=>o in e?ie(e,o,{enumerable:!0,configurable:!0,writable:!0,value:a}):e[o]=a,m=(e,o)=>{for(var a in o||(o={}))re.call(o,a)&&A(e,a,o[a]);if(H)for(var a of H(o))ue.call(o,a)&&A(e,a,o[a]);return e},S=(e,o)=>le(e,se(o));import{a as ce,r as q,c as h,w as D,L,a6 as M,a4 as j,h as Q,T as de,g as fe,b as me,V as ve}from"./index.e18600f8.js";import{u as he,v as W,b as ge,c as Te,a as ye,r as N,s as pe,p as R,d as be}from"./position-engine.e707f0b1.js";import{d as Se,e as Pe,u as we,f as Oe,g as Ce}from"./use-prevent-scroll.d4f2a0d7.js";import{c as ke,u as Ee,d as xe,e as He}from"./QDialog.07d9a48e.js";import{c as V}from"./QList.68ba6ad8.js";var Qe=ce({name:"QTooltip",inheritAttrs:!1,props:S(m(m(m({},he),Se),ke),{maxHeight:{type:String,default:null},maxWidth:{type:String,default:null},transitionShow:{default:"jump-down"},transitionHide:{default:"jump-up"},anchor:{type:String,default:"bottom middle",validator:W},self:{type:String,default:"top middle",validator:W},offset:{type:Array,default:()=>[14,14],validator:ge},scrollTarget:{default:void 0},delay:{type:Number,default:0},hideDelay:{type:Number,default:0}}),emits:[...Pe],setup(e,{slots:o,emit:a,attrs:g}){let s,r;const T=fe(),{proxy:{$q:i}}=T,u=q(null),c=q(!1),_=h(()=>R(e.anchor,i.lang.rtl)),B=h(()=>R(e.self,i.lang.rtl)),I=h(()=>e.persistent!==!0),{registerTick:U,removeTick:P}=Ee(),{registerTimeout:v,removeTimeout:y}=we(),{transition:$,transitionStyle:z}=xe(e,c),{localScrollTarget:w,changeScrollEvent:F,unconfigureScrollTarget:G}=Te(e,E),{anchorEl:n,canShow:J,anchorEvents:d}=ye({showing:c,configureAnchorEl:oe}),{show:K,hide:p}=Oe({showing:c,canShow:J,handleShow:Y,handleHide:Z,hideOnRouteChange:I,processOnMount:!0});Object.assign(d,{delayShow:ee,delayHide:te});const{showPortal:O,hidePortal:C,renderPortal:X}=He(T,u,ne);if(i.platform.is.mobile===!0){const t={anchorEl:n,innerRef:u,onClickOutside(l){return p(l),l.target.classList.contains("q-dialog__backdrop")&&ve(l),!0}},b=h(()=>e.modelValue===null&&e.persistent!==!0&&c.value===!0);D(b,l=>{(l===!0?be:N)(t)}),L(()=>{N(t)})}function Y(t){P(),y(),O(),U(()=>{r=new MutationObserver(()=>f()),r.observe(u.value,{attributes:!1,childList:!0,characterData:!0,subtree:!0}),f(),E()}),s===void 0&&(s=D(()=>i.screen.width+"|"+i.screen.height+"|"+e.self+"|"+e.anchor+"|"+i.lang.rtl,f)),v(()=>{O(!0),a("show",t)},e.transitionDuration)}function Z(t){P(),y(),C(),k(),v(()=>{C(!0),a("hide",t)},e.transitionDuration)}function k(){r!==void 0&&(r.disconnect(),r=void 0),s!==void 0&&(s(),s=void 0),G(),M(d,"tooltipTemp")}function f(){const t=u.value;n.value===null||!t||pe({el:t,offset:e.offset,anchorEl:n.value,anchorOrigin:_.value,selfOrigin:B.value,maxHeight:e.maxHeight,maxWidth:e.maxWidth})}function ee(t){if(i.platform.is.mobile===!0){V(),document.body.classList.add("non-selectable");const b=n.value,l=["touchmove","touchcancel","touchend","click"].map(x=>[b,x,"delayHide","passiveCapture"]);j(d,"tooltipTemp",l)}v(()=>{K(t)},e.delay)}function te(t){y(),i.platform.is.mobile===!0&&(M(d,"tooltipTemp"),V(),setTimeout(()=>{document.body.classList.remove("non-selectable")},10)),v(()=>{p(t)},e.hideDelay)}function oe(){if(e.noParentEvent===!0||n.value===null)return;const t=i.platform.is.mobile===!0?[[n.value,"touchstart","delayShow","passive"]]:[[n.value,"mouseenter","delayShow","passive"],[n.value,"mouseleave","delayHide","passive"]];j(d,"anchor",t)}function E(){if(n.value!==null||e.scrollTarget!==void 0){w.value=Ce(n.value,e.scrollTarget);const t=e.noParentEvent===!0?f:p;F(w.value,t)}}function ae(){return c.value===!0?Q("div",S(m({},g),{ref:u,class:["q-tooltip q-tooltip--style q-position-engine no-pointer-events",g.class],style:[g.style,z.value],role:"complementary"}),me(o.default)):null}function ne(){return Q(de,{name:$.value,appear:!0},ae)}return L(k),Object.assign(T.proxy,{updatePosition:f}),X}});export{Qe as Q};