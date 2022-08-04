import { createApp } from 'vue';

import TDesign from 'tdesign-vue-next';
import 'tdesign-vue-next/es/style/index.css';

// VMdEditor MarkDown编辑器
import VMdEditor from '@kangc/v-md-editor/lib/codemirror-editor';
import '@kangc/v-md-editor/lib/style/codemirror-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';
// VMdEditor highlightjs
import hljs from 'highlight.js';
// VMdEditor codemirror
import Codemirror from 'codemirror';
// VMdEditor mode
import 'codemirror/mode/markdown/markdown.js';
import 'codemirror/mode/javascript/javascript.js';
import 'codemirror/mode/css/css.js';
import 'codemirror/mode/htmlmixed/htmlmixed.js';
import 'codemirror/mode/vue/vue.js';
// VMdEditor edit
import 'codemirror/addon/edit/closebrackets.js';
import 'codemirror/addon/edit/closetag.js';
import 'codemirror/addon/edit/matchbrackets.js';
// VMdEditor placeholder
import 'codemirror/addon/display/placeholder.js';
// VMdEditor active-line
import 'codemirror/addon/selection/active-line.js';
// VMdEditor scrollbar
import 'codemirror/addon/scroll/simplescrollbars.js';
import 'codemirror/addon/scroll/simplescrollbars.css';
// VMdEditor style
import 'codemirror/lib/codemirror.css';

import { store } from './store';
import router from './router';
import '@/style/index.less';
import './permission';
import App from './App.vue';

VMdEditor.Codemirror = Codemirror;
VMdEditor.use(githubTheme, {
  Hljs: hljs,
});

VMdEditor.Codemirror = Codemirror;

const app = createApp(App);

app.use(TDesign);
app.use(store);
app.use(router);
app.use(VMdEditor);

app.mount('#app');
