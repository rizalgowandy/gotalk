package gotalk

const JSLibSHA1Base64 = "79W3b3jTN9Zi/JMrJFT7u/cNGbE="
const JSLibString = "var gotalk=(function(){var I=Object.defineProperty,ce=Object.prototype.hasOwnProperty,q=function(e,t){return function(){return t||(t={exports:{}},e(t.exports,t)),t.exports}},Y=function(e){return I(e,\"__esModule\",{value:!0})},Z=function(e,t){Y(e);for(var r in t)I(e,r,{get:t[r],enumerable:!0})},pe=function(e,t){if(Y(e),typeof t==\"object\"||typeof t==\"function\")for(var r in t)!ce.call(e,r)&&r!==\"default\"&&(function(o){I(e,o,{get:function(){return t[o]},enumerable:!0})})(r);return e},w=function(e){return e&&e.__esModule?e:pe(I({},\"default\",{value:e,enumerable:!0}),e)};var B=q(function(R){Z(R,{sizeOf:function(){return $}});function $(e){for(var t=0,r=0,o;o=e.charCodeAt(r++);t+=o>>11?3:o>>7?2:1);return t}function O(e){return 255&e}if(typeof TextDecoder!=\"undefined\"){var ge=new TextDecoder(\"utf8\"),he=new TextEncoder(\"utf8\");R.decode=function(t){return ge.decode(t)},R.encode=function(t){return he.encode(t)}}else R.decode=function(t){var r=0,o=t.length-r,n,s,i=\"\";for(r=0;r<o;)n=t[r++],s=O(n),s<128||(s>>5==6?n=(n<<6&2047)+(t[r++]&63):s>>4==14?(n=(n<<12&65535)+(O(t[r++])<<6&4095),n+=t[r++]&63):s>>3==30&&(n=(n<<18&2097151)+(O(t[r++])<<12&262143),n+=O(t[r++])<<6&4095,n+=t[r++]&63)),i+=String.fromCharCode(n);return i},R.encode=function(t){for(var r=0,o=t.length,n,s=0,i=new Uint8Array($(t));r!==o;)n=t.charCodeAt(r++),n<128?i[s++]=n:n<2048?(i[s++]=n>>6|192,i[s++]=n&63|128):n<65536?(i[s++]=n>>12|224,i[s++]=n>>6&63|128,i[s++]=n&63|128):(i[s++]=n>>18|240,i[s++]=n>>12&63|128,i[s++]=n>>6&63|128,i[s++]=n&63|128);return i}});var D=q(function(_){function W(){}_.global=typeof global!=\"undefined\"?global:typeof self!=\"undefined\"?self:typeof window!=\"undefined\"?window:_;_.console=typeof console!=\"undefined\"?console:{log:W,warn:W,error:W};_.document=typeof document!=\"undefined\"?document:null});var V=q(function(d){const k=w(B());d.Version=1;var ee=d.MsgTypeSingleReq=114,te=d.MsgTypeStreamReq=115,_e=d.MsgTypeStreamReqPart=112,Ce=d.MsgTypeSingleRes=82,Ae=d.MsgTypeStreamRes=83,He=d.MsgTypeErrorRes=69,P=d.MsgTypeRetryRes=101,z=d.MsgTypeNotification=110,N=d.MsgTypeHeartbeat=104,re=d.MsgTypeProtocolError=102;d.ErrorAbnormal=0;d.ErrorUnsupported=1;d.ErrorInvalidMsg=2;d.ErrorTimeout=3;d.HeartbeatMsgMaxLoad=65535;function S(e,t,r,o){for(var n=t||0,s=0,i,a=r.toString(16),m=o-a.length;m--;)e[n++]=48;for(;!isNaN(i=a.charCodeAt(s++));)e[n++]=i}function oe(e,t){var r=g(t);return S(r,0,e,t),r}function C(e){for(var t=0,r=0,o=0,n=!1;r<e.length;r++)o=e[r],o<64?(o<48&&(n=!0),o-=48):o<71?(o<65&&(n=!0),o-=55):o<103&&96<o?o-=87:n=!0,t=t<<4|o&15;if(n)throw new Error(\"invalid hexint \"+String.fromCharCode.apply(null,e));return t}d.binary={makeFixnum:oe,versionBuf:oe(d.Version,2),parseVersion:C,parseMsg:function(e){var t,r,o,n,s=0,i=0,a;return t=e[0],a=1,t===N?(s=C(e.subarray(a,a+4)),a+=4):t!==z&&t!==re&&(r=e.subarray(a,a+4),a+=4),t==ee||t==te||t==z?(n=C(e.subarray(a,a+3)),a+=3,o=k.decode(e.subarray(a,a+n)),a+=n):t===P&&(s=C(e.subarray(a,a+8)),a+=8),i=C(e.subarray(a,a+8)),{t:t,id:r,name:o,wait:s,size:i}},makeMsg:function(e,t,r,o,n){var s,i,a=t?13:9;return r&&r.length!==0&&(i=k.encode(r),a+=3+i.length),s=g(a),s[0]=e,a=1,t&&t.length===4&&(typeof t==\"string\"?(s[1]=t.charCodeAt(0),s[2]=t.charCodeAt(1),s[3]=t.charCodeAt(2),s[4]=t.charCodeAt(3)):(s[1]=t[0],s[2]=t[1],s[3]=t[2],s[4]=t[3]),a+=4),i&&(S(s,a,i.length,3),a+=3,s.set(i,a),a+=i.length),e===P&&(S(s,a,o,8),a+=8),S(s,a,n,8),s},makeHeartbeatMsg:function(e){var t=g(13),r=1;return t[0]=N,S(t,r,e,4),r+=4,S(t,r,Math.round(new Date().getTime()/1e3),8),r+=8,t}};var me=\"00000000\";function x(e,t){var r=e.toString(16);return me.substr(0,t-r.length)+r}d.text={makeFixnum:x,versionBuf:x(d.Version,2),parseVersion:function(e){return parseInt(e.substr(0,2),16)},parseMsg:function(e){var t,r,o,n=0,s=0,i;return t=e.charCodeAt(0),i=1,t===N?(n=parseInt(e.substr(i,4),16),i+=4):t!==z&&t!==re&&(r=e.substr(i,4),i+=4),t==ee||t==te||t==z?o=e.substring(i+3,e.length-8):t==P&&(n=parseInt(e.substr(i,8),16),i+=8),s=parseInt(e.substr(e.length-8),16),{t:t,id:r,name:o,wait:n,size:s}},makeMsg:function(e,t,r,o,n){var s=String.fromCharCode(e);return t&&t.length===4&&(s+=t),r&&r.length!==0&&(s+=x(k.sizeOf(r),3),s+=r),e===P&&(s+=x(o,8)),s+=x(n,8),s},makeHeartbeatMsg:function(e){var t=String.fromCharCode(N);return t+=x(e,4),t+=x(Math.round(new Date().getTime()/1e3),8),t}}});var de=q(function(v){Z(v,{default:function(){return se}});const T=w(D()),l=w(V()),A=w(B());var u=v,se=v,ie=l.text,ae=l.binary;u.protocol=l;u.Buf=g;u.developmentMode=!1;u.defaultResponderAddress=\"\";var le=\"\",c={wsproto:\"\",proto:\"\",host:\"\",path:\"\"};function ye(){T.document&&ve(),u.developmentMode=be(c.host)}function ve(){var e=T.document.currentScript.src;if(!e)return;var t=e.indexOf(\"://\")+3;if(t==2)return;c.proto=e.substr(0,t-2);var r=e.indexOf(\"/\",t);if(r==-1)return;c.wsproto=c.proto==\"https:\"?\"wss://\":\"ws://\",c.host=e.substring(t,r),e=e.substr(r),t=e.lastIndexOf(\"?\"),t!=-1&&(e=e.substr(0,t)),c.path=e.substring(e.indexOf(\"/\"),e.lastIndexOf(\"/\")+1),u.defaultResponderAddress=c.wsproto+c.host+c.path,le=u.defaultResponderAddress}function be(e){var t=e,r=t.lastIndexOf(\":\");return r!=-1&&(t=t.substr(0,r)),r=t.lastIndexOf(\".\"),r==-1?t==\"localhost\":t==\"127.0.0.1\"||t.substr(r)==\".local\"}function U(){u.developmentMode&&T.console.warn.apply(T.console,Array.prototype.slice.call(arguments))}function j(e){if(!e||e.length==0)return null;typeof e!=\"string\"&&(e=A.decode(e));try{return JSON.parse(e)}catch(t){U(\"[gotalk] ignoring invalid json\",e)}}function f(e,t){return Object.create(f.prototype,{handlers:{value:e,enumerable:!0},protocol:{value:t||(g?l.binary:l.text),enumerable:!0,writable:!0},heartbeatInterval:{value:20*1e3,enumerable:!0,writable:!0},ws:{value:null,writable:!0},keepalive:{value:null,writable:!0},nextOpID:{value:0,writable:!0},nextStreamID:{value:0,writable:!0},pendingRes:{value:{},writable:!0},hasPendingRes:{get:function(){for(var r in this.pendingRes)return!0}},pendingClose:{value:!1,writable:!0}})}f.prototype=p.mixin(f.prototype);v.Sock=f;function we(e,t){if(e.pendingClose=!1,e.stopSendingHeartbeats(),e.ws&&(e.ws.onmessage=null,e.ws.onerror=null,e.ws.onclose=null,e.ws=null),e.nextOpID=0,e.hasPendingRes){var r=t||new Error(\"connection closed\");for(var o in e.pendingRes)e.pendingRes[o](r);e.pendingRes={}}}var xe={1e3:\"normal\",1001:\"going away\",1002:\"protocol error\",1003:\"unsupported\",1005:\"no status\",1006:\"abnormal\",1007:\"inconsistent\",1008:\"invalid message\",1009:\"too large\"};function J(e){var t=xe[e];return\"#\"+e+(t?\" (\"+t+\")\":\"\")}f.prototype.adoptWebSocket=function(e){var t=this;if(e.readyState!==WebSocket.OPEN)throw new Error(\"web socket readyState != OPEN\");e.binaryType=\"arraybuffer\",t.ws=e,e.onclose=function(r){var o=e._gotalkCloseError;!o&&r.code!==1e3&&(o=new Error(\"websocket closed: \"+J(r.code))),we(t,o),t.emit(\"close\",o)},e.onmessage=function(r){e._bufferedMessages||(e._bufferedMessages=[]),e._bufferedMessages.push(r.data)}};f.prototype.adopt=function(e){if(adopt instanceof WebSocket)return this.adoptWebSocket(e);throw new Error(\"unsupported transport\")};f.prototype.handshake=function(){this.ws.send(this.protocol.versionBuf)};f.prototype.end=function(){var e=this;e.keepalive&&(e.keepalive.disable(),e.keepalive=null),!e.pendingClose&&e.hasPendingRes?e.pendingClose=!0:e.ws&&e.ws.close(1e3)};f.prototype.address=function(){var e=this;return e.ws?e.ws.url:null};var G=v.ErrAbnormal=Error(\"unsupported protocol\");G.isGotalkProtocolError=!0;G.code=l.ErrorAbnormal;var L=v.ErrUnsupported=Error(\"unsupported protocol\");L.isGotalkProtocolError=!0;L.code=l.ErrorUnsupported;var F=v.ErrInvalidMsg=Error(\"invalid protocol message\");F.isGotalkProtocolError=!0;F.code=l.ErrorInvalidMsg;var K=v.ErrTimeout=Error(\"timeout\");K.isGotalkProtocolError=!0;K.code=l.ErrorTimeout;f.prototype.sendHeartbeat=function(e){var t=this,r=t.protocol.makeHeartbeatMsg(Math.round(e*l.HeartbeatMsgMaxLoad));try{t.ws.send(r)}catch(o){throw(!this.ws||this.ws.readyState>WebSocket.OPEN)&&(o=new Error(\"socket is closed\")),o}};f.prototype.startSendingHeartbeats=function(){var e=this;if(e.heartbeatInterval<10)throw new Error(\"Sock.heartbeatInterval is too low\");clearTimeout(e._sendHeartbeatsTimer);var t=function(){clearTimeout(e._sendHeartbeatsTimer),e.sendHeartbeat(0),e._sendHeartbeatsTimer=setTimeout(t,e.heartbeatInterval)};e._sendHeartbeatsTimer=setTimeout(t,1)};f.prototype.stopSendingHeartbeats=function(){var e=this;clearTimeout(e._sendHeartbeatsTimer)};f.prototype.startReading=function(){var e=this,t=e.ws,r;function o(i){if(r=typeof i.data==\"string\"?ie.parseMsg(i.data):ae.parseMsg(g(i.data)),r.t===l.MsgTypeProtocolError){var a=r.size;a===l.ErrorAbnormal?t._gotalkCloseError=G:a===l.ErrorUnsupported?t._gotalkCloseError=L:a===l.ErrorTimeout?t._gotalkCloseError=K:t._gotalkCloseError=F,t.close(4e3+a)}else r.size!==0&&r.t!==l.MsgTypeHeartbeat?t.onmessage=n:(e.handleMsg(r),r=null)}function n(i){var a=i.data;t.onmessage=o,e.handleMsg(r,typeof a==\"string\"?a:g(a)),r=null}function s(i){var a=typeof i.data==\"string\"?ie.parseVersion(i.data):ae.parseVersion(g(i.data));a!==l.Version?(t._gotalkCloseError=L,e.closeError(l.ErrorUnsupported)):(t.onmessage=o,e.heartbeatInterval>0&&e.startSendingHeartbeats())}t.onmessage=s,t._bufferedMessages&&(t._bufferedMessages.forEach(function(i){t.onmessage({data:i})}),t._bufferedMessages=null)};var M={};f.prototype.handleMsg=function(e,t){var r=this,o=M[e.t];o?o.call(r,e,t):(r.ws&&(r.ws._gotalkCloseError=F),r.closeError(l.ErrorInvalidMsg))};M[l.MsgTypeSingleReq]=function(e,t){var r=this,o,n;if(o=r.handlers.findRequestHandler(e.name),n=function(s){r.sendMsg(l.MsgTypeSingleRes,e.id,null,0,s)},n.error=function(s){var i=s.message||String(s);r.sendMsg(l.MsgTypeErrorRes,e.id,null,0,i)},typeof o!=\"function\")n.error('no such operation \"'+e.name+'\"');else try{o(t,n,e.name)}catch(s){U(\"[gotalk] handler error:\",s.stack||\"\"+s),n.error(\"internal error\")}};function X(e,t){var r=e.id;typeof r!=\"string\"&&(r=String.fromCharCode.apply(null,r));var o=this,n=o.pendingRes[r];if((e.t!==l.MsgTypeStreamRes||!t||(t.length||t.size)===0)&&(delete o.pendingRes[r],o.pendingClose&&!o.hasPendingRes&&o.end()),typeof n!=\"function\")return;e.t===l.MsgTypeErrorRes?(typeof t!=\"string\"&&(t=A.decode(t)),n(new Error(t),null)):n(null,t)}M[l.MsgTypeSingleRes]=X;M[l.MsgTypeStreamRes]=X;M[l.MsgTypeErrorRes]=X;M[l.MsgTypeNotification]=function(e,t){var r=this.handlers.findNotificationHandler(e.name);r&&r(t,e.name)};M[l.MsgTypeHeartbeat]=function(e){this.emit(\"heartbeat\",{time:new Date(e.size*1e3),load:e.wait})};f.prototype.sendMsg=function(e,t,r,o,n){var s=n&&typeof n==\"string\"&&this.protocol===l.binary?A.sizeOf(n):n?n.length||n.size:0,i=this,a=i.protocol.makeMsg(e,t,r,o,s);try{i.ws.send(a),s!==0&&i.ws.send(n)}catch(m){throw(!this.ws||this.ws.readyState>WebSocket.OPEN)&&(m=new Error(\"socket is closed\")),m}};f.prototype.closeError=function(e){var t=this,r;if(t.ws){try{t.ws.send(t.protocol.makeMsg(l.MsgTypeProtocolError,null,null,0,e))}catch(o){}t.ws.close(4e3+e)}};var fe=\"0000\";f.prototype.bufferRequest=function(e,t,r){var o=this,n=o.nextOpID++;o.nextOpID===1679616&&(o.nextOpID=0),n=n.toString(36),n=fe.substr(0,4-n.length)+n,o.pendingRes[n]=r;try{o.sendMsg(l.MsgTypeSingleReq,n,e,0,t)}catch(s){delete o.pendingRes[n],r(s)}};f.prototype.bufferNotify=function(e,t){this.sendMsg(l.MsgTypeNotification,null,e,0,t)};f.prototype.request=function(e,t,r){var o;return r?o=JSON.stringify(t):r=t,this.bufferRequest(e,o,function(n,s){var i=j(s);return r(n,i)})};f.prototype.notify=function(e,t){var r=JSON.stringify(t);return this.bufferNotify(e,r)};f.prototype.requestp=function(e,t){var r=this;return new Promise(function(o,n){r.request(e,t,function(s,i){s?n(s):o(i)})})};f.prototype.bufferRequestp=function(e,t){var r=this;return new Promise(function(o,n){r.bufferRequest(e,t,function(s,i){s?n(s):o(i)})})};var H=function(e,t,r){return Object.create(H.prototype,{s:{value:e},op:{value:t,enumerable:!0},id:{value:r,enumerable:!0}})};p.mixin(H.prototype);H.prototype.write=function(e){this.ended||(this.started?this.s.sendMsg(l.MsgTypeStreamReqPart,this.id,null,0,e):(this.started=!0,this.s.sendMsg(l.MsgTypeStreamReq,this.id,this.op,0,e)),(!e||e.length===0||e.size===0)&&(this.ended=!0))};H.prototype.end=function(){this.write(null)};f.prototype.streamRequest=function(e){var t=this,r=t.nextStreamID++;t.nextStreamID===46656&&(t.nextStreamID=0),r=r.toString(36),r=\"!\"+fe.substr(0,3-r.length)+r;var o=H(t,e,r);return t.pendingRes[r]=function(n,s){n?o.emit(\"end\",n):!s||s.length===0?o.emit(\"end\",null):o.emit(\"data\",s)},o};function b(){return Object.create(b.prototype,{reqHandlers:{value:{}},reqFallbackHandler:{value:null,writable:!0},noteHandlers:{value:{}},noteFallbackHandler:{value:null,writable:!0}})}v.Handlers=b;b.prototype.handleBufferRequest=function(e,t){e?this.reqHandlers[e]=t:this.reqFallbackHandler=t};b.prototype.handleRequest=function(e,t){return this.handleBufferRequest(e,function(r,o,n){var s=function(a){return o(JSON.stringify(a))};s.error=o.error;var i=j(r);t(i,s,n)})};b.prototype.handleBufferNotification=function(e,t){e?this.noteHandlers[e]=t:this.noteFallbackHandler=t};b.prototype.handleNotification=function(e,t){this.handleBufferNotification(e,function(r,o){t(j(r),o)})};b.prototype.findRequestHandler=function(e){var t=this.reqHandlers[e];return t||this.reqFallbackHandler};b.prototype.findNotificationHandler=function(e){var t=this.noteHandlers[e];return t||this.noteFallbackHandler};var ue=!1;function Me(e,t,r){var o;try{o=new WebSocket(t),o.binaryType=\"arraybuffer\",o.onclose=function(n){u.developmentMode&&!ue&&le==u.defaultResponderAddress&&(ue=!0,U(\"gotalk connection failed with code \"+J(n.code)+\". If you are serving gotalk.js yourself, remember to set gotalk.defaultResponderAddress to the gotalk websocket endpoint.\"));var s=new Error(\"connection failed: \"+J(n.code));r&&r(s)},o.onopen=function(n){o.onerror=void 0,e.adoptWebSocket(o),e.handshake(),r&&r(null,e),e.emit(\"open\"),e.startReading()},o.onmessage=function(n){o._bufferedMessages||(o._bufferedMessages=[]),o._bufferedMessages.push(n.data)}}catch(n){U(\"[gotalk] WebSocket init error:\",n.stack||\"\"+n),r&&r(n),e.emit(\"close\",n)}}function Ee(e){e||(e=u.defaultResponderAddress);var t=e.substr(0,4);if(t!=\"ws:/\"&&t!=\"wss:\"&&(c.proto&&(e[0]==\"/\"?e[1]==\"/\"?e=c.wsproto+e:e=c.wsproto+c.host+e:e=c.wsproto+c.host+\"/\"+e)),!e)throw new Error(\"address not specified\");return e}f.prototype.open=function(e,t){var r=this;return!t&&typeof e==\"function\"&&(t=e,e=null),Me(r,Ee(e),t),r};u.open=function(e,t,r,o){return f(r||u.defaultHandlers,o).open(e,t)};f.prototype.openKeepAlive=function(e){var t=this;return t.keepalive&&t.keepalive.disable(),t.keepalive=ne(t,e),t.keepalive.enable(),t};u.connection=function(e,t,r){return f(t||u.defaultHandlers,r).openKeepAlive(e)};u.defaultHandlers=b();u.handleBufferRequest=function(e,t){return u.defaultHandlers.handleBufferRequest(e,t)};u.handle=function(e,t){return u.defaultHandlers.handleRequest(e,t)};u.handleBufferNotification=function(e,t){return u.defaultHandlers.handleBufferNotification(e,t)};u.handleNotification=function(e,t){return u.defaultHandlers.handleNotification(e,t)};ye()});const Ne=w(B());var g=function(){return typeof Uint8Array==\"undefined\"?null:function(t){return t instanceof Uint8Array?t:new Uint8Array(t instanceof ArrayBuffer?t:new ArrayBuffer(t))}}();function p(){}p.prototype.addListener=function(e,t){if(typeof t!=\"function\")throw TypeError(\"listener must be a function\");if(!this.__events)return Object.defineProperty(this,\"__events\",{value:{},enumerable:!1,writable:!0}),this.__events[e]=[t],this;var r=this.__events[e];return r===void 0?(this.__events[e]=[t],this):(r.push(t),this)};p.prototype.on=p.prototype.addListener;p.prototype.once=function(e,t){var r=!1,o=function(){this.removeListener(e,o),r||(r=!0,t.apply(this,arguments))};return this.on(e,o)};p.prototype.removeListener=function(e,t){var r,o=this.__events?this.__events[e]:void 0;if(o!==void 0){for(;(r=o.indexOf(t))!==-1;)o.splice(r,1);return o.length===0&&delete this.__events[e],o.length}return this};p.prototype.removeAllListeners=function(e){this.__events&&(e?delete this.__events[e]:delete this.__events)};p.prototype.listeners=function(e){return e?this.__events?this.__events[e]:void 0:this.__events};p.prototype.emit=function(e){var t=this.__events?this.__events[e]:void 0;if(t===void 0)return!1;for(var r=0,o=t.length,n=Array.prototype.slice.call(arguments,1);r!==o;++r)t[r].apply(this,n);return!0};p.mixin=function(t){for(var r=t;r;){if(r.__proto__===Object.prototype)return r.__proto__=p.prototype,t;r=r.__proto__}return t};const h=w(D()),Fe=w(V());var y={available:!1,onLine:!0};h.global.addEventListener&&(y.available=!0,y.onLine=typeof navigator!=\"undefined\"?navigator.onLine:!0,h.global.addEventListener(\"offline\",function(e){y.onLine=!1,y.emit(\"offline\")}),h.global.addEventListener(\"online\",function(e){y.onLine=!0,y.emit(\"online\")}));function ne(e,t,r,o){r?r<100&&(r=100):r=500,(!o||o<r)&&(o=5e3);var n,s,i,a=0,m,Q;return n={isEnabled:!1,isConnected:!1,enable:function(){n.enabled||(n.enabled=!0,a=0,n.isConnected||s())},disable:function(){n.enabled&&(clearTimeout(m),n.enabled=!1,a=0)}},s=function(){clearTimeout(m),e.open(t,function(E){Q=new Date,E?i(E):(a=0,n.isConnected=!0,e.once(\"close\",i))})},i=function(E){if(clearTimeout(m),n.isConnected=!1,!n.enabled)return;if(y.available&&!y.onLine&&!(h.document&&h.document.location&&h.document.location.hostname!==\"localhost\"&&h.document.location.hostname!==\"127.0.0.1\"&&h.document.location.hostname!==\"[::1]\")){y.once(\"online\",i),a=0;return}E?E.isGotalkProtocolError?E.code===protocol.ErrorTimeout?a=0:a=o:a=a?Math.min(o,a*2):r:a=Math.max(0,r-(new Date-Q)),m=setTimeout(s,a)},n}return de();})();\n"
