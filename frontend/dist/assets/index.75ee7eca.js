(function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const r of document.querySelectorAll('link[rel="modulepreload"]'))s(r);new MutationObserver(r=>{for(const i of r)if(i.type==="childList")for(const l of i.addedNodes)l.tagName==="LINK"&&l.rel==="modulepreload"&&s(l)}).observe(document,{childList:!0,subtree:!0});function n(r){const i={};return r.integrity&&(i.integrity=r.integrity),r.referrerpolicy&&(i.referrerPolicy=r.referrerpolicy),r.crossorigin==="use-credentials"?i.credentials="include":r.crossorigin==="anonymous"?i.credentials="omit":i.credentials="same-origin",i}function s(r){if(r.ep)return;r.ep=!0;const i=n(r);fetch(r.href,i)}})();function x(){}function G(e){return e()}function S(){return Object.create(null)}function k(e){e.forEach(G)}function I(e){return typeof e=="function"}function D(e,t){return e!=e?t==t:e!==t||e&&typeof e=="object"||typeof e=="function"}let E;function F(e,t){return E||(E=document.createElement("a")),E.href=t,e===E.href}function K(e){return Object.keys(e).length===0}function d(e,t){e.appendChild(t)}function z(e,t,n){e.insertBefore(t,n||null)}function M(e){e.parentNode&&e.parentNode.removeChild(e)}function _(e){return document.createElement(e)}function T(e){return document.createTextNode(e)}function j(){return T(" ")}function C(e,t,n,s){return e.addEventListener(t,n,s),()=>e.removeEventListener(t,n,s)}function c(e,t,n){n==null?e.removeAttribute(t):e.getAttribute(t)!==n&&e.setAttribute(t,n)}function H(e){return Array.from(e.childNodes)}function J(e,t){t=""+t,e.data!==t&&(e.data=t)}function q(e,t){e.value=t==null?"":t}let N;function v(e){N=e}const g=[],P=[];let y=[];const B=[],Q=Promise.resolve();let O=!1;function R(){O||(O=!0,Q.then(W))}function L(e){y.push(e)}const A=new Set;let m=0;function W(){if(m!==0)return;const e=N;do{try{for(;m<g.length;){const t=g[m];m++,v(t),U(t.$$)}}catch(t){throw g.length=0,m=0,t}for(v(null),g.length=0,m=0;P.length;)P.pop()();for(let t=0;t<y.length;t+=1){const n=y[t];A.has(n)||(A.add(n),n())}y.length=0}while(g.length);for(;B.length;)B.pop()();O=!1,A.clear(),v(e)}function U(e){if(e.fragment!==null){e.update(),k(e.before_update);const t=e.dirty;e.dirty=[-1],e.fragment&&e.fragment.p(e.ctx,t),e.after_update.forEach(L)}}function V(e){const t=[],n=[];y.forEach(s=>e.indexOf(s)===-1?t.push(s):n.push(s)),n.forEach(s=>s()),y=t}const X=new Set;function Y(e,t){e&&e.i&&(X.delete(e),e.i(t))}function Z(e,t,n,s){const{fragment:r,after_update:i}=e.$$;r&&r.m(t,n),s||L(()=>{const l=e.$$.on_mount.map(G).filter(I);e.$$.on_destroy?e.$$.on_destroy.push(...l):k(l),e.$$.on_mount=[]}),i.forEach(L)}function ee(e,t){const n=e.$$;n.fragment!==null&&(V(n.after_update),k(n.on_destroy),n.fragment&&n.fragment.d(t),n.on_destroy=n.fragment=null,n.ctx=[])}function te(e,t){e.$$.dirty[0]===-1&&(g.push(e),R(),e.$$.dirty.fill(0)),e.$$.dirty[t/31|0]|=1<<t%31}function ne(e,t,n,s,r,i,l,p=[-1]){const f=N;v(e);const o=e.$$={fragment:null,ctx:[],props:i,update:x,not_equal:r,bound:S(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(t.context||(f?f.$$.context:[])),callbacks:S(),dirty:p,skip_bound:!1,root:t.target||f.$$.root};l&&l(o.root);let $=!1;if(o.ctx=n?n(e,t.props||{},(u,h,...b)=>{const a=b.length?b[0]:h;return o.ctx&&r(o.ctx[u],o.ctx[u]=a)&&(!o.skip_bound&&o.bound[u]&&o.bound[u](a),$&&te(e,u)),h}):[],o.update(),$=!0,k(o.before_update),o.fragment=s?s(o.ctx):!1,t.target){if(t.hydrate){const u=H(t.target);o.fragment&&o.fragment.l(u),u.forEach(M)}else o.fragment&&o.fragment.c();t.intro&&Y(e.$$.fragment),Z(e,t.target,t.anchor,t.customElement),W()}v(f)}class re{$destroy(){ee(this,1),this.$destroy=x}$on(t,n){if(!I(n))return x;const s=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return s.push(n),()=>{const r=s.indexOf(n);r!==-1&&s.splice(r,1)}}$set(t){this.$$set&&!K(t)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}function oe(e){return window.go.main.App.Search(e)}const se="/assets/Dramatic_Chipmunk.53260775.png";function ie(e){let t,n,s,r,i,l,p,f,o,$,u,h,b;return{c(){t=_("main"),n=_("img"),r=j(),i=_("div"),l=T(e[0]),p=j(),f=_("div"),o=_("input"),$=j(),u=_("button"),u.textContent="Search",c(n,"alt","Wails logo"),c(n,"id","logo"),F(n.src,s=se)||c(n,"src",s),c(n,"class","svelte-sjdsfp"),c(i,"class","result svelte-sjdsfp"),c(i,"id","result"),c(o,"autocomplete","off"),c(o,"class","input svelte-sjdsfp"),c(o,"id","query"),c(o,"type","text"),c(u,"class","btn svelte-sjdsfp"),c(f,"class","input-box svelte-sjdsfp"),c(f,"id","input")},m(a,w){z(a,t,w),d(t,n),d(t,r),d(t,i),d(i,l),d(t,p),d(t,f),d(f,o),q(o,e[1]),d(f,$),d(f,u),h||(b=[C(o,"input",e[3]),C(u,"click",e[2])],h=!0)},p(a,[w]){w&1&&J(l,a[0]),w&2&&o.value!==a[1]&&q(o,a[1])},i:x,o:x,d(a){a&&M(t),h=!1,k(b)}}}function le(e,t,n){let s="Welcome to the Go Chatbot! Please enter a Go concept that you would like to explore.",r;function i(){oe(r).then(p=>n(0,s=p))}function l(){r=this.value,n(1,r)}return[s,r,i,l]}class ue extends re{constructor(t){super(),ne(this,t,le,ie,D,{})}}new ue({target:document.getElementById("app")});
