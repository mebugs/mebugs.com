import Vditor from 'vditor'
import { h, reactive, ref } from 'vue'

// mardown编辑器对象
export interface markdown {
  md: string | undefined // MD 基础文本
  html: string | undefined // HTML 解析结果
  outline: string | undefined // 大纲 HTML
}

const vd = ref<Vditor>({} as any)

// 编辑器插件初始化
export default function vditor(nodeId: string, openSource: Function) {
  const md: markdown = reactive<markdown>({
    md: '',
    html: '',
    outline: ''
  })

  // 编辑器工具栏
  const toolbar = [
    'emoji',
    'headings',
    'bold',
    'italic',
    'strike',
    'link',
    '|',
    'list',
    'ordered-list',
    'check',
    'outdent',
    'indent',
    '|',
    'quote',
    'line',
    'code',
    'inline-code',
    'insert-before',
    'insert-after',
    '|',
    {
      hotkey: '⇧⌘u',
      name: 'uploadSource',
      tipPosition: 'n',
      tip: '上传/选择资源',
      className: 'right',
      icon: '<svg><use xlink:href="#vditor-icon-upload"></use></svg>',
      click() {
        openSource()
      }
    },
    'table',
    '|',
    'undo',
    'redo',
    '|',
    'fullscreen',
    'preview'
  ]

  // 表情对象
  const emoji = {
    '100': '💯',
    anchor: '⚓️',
    anger: '💢',
    boom: '💥',
    dart: '🎯',
    fire: '🔥',
    gem: '💎',
    construction: '🚧',
    crossed_swords: '⚔️',
    crown: '👑',
    art: '🎨',
    key: '🔑',
    book: '📖',
    heavy_check_mark: '✔️',
    heavy_multiplication_x: '✖️',
    pushpin: '📌',
    bell: '🔔',
    heart: '❤️',
    broken_heart: '💔',
    bulb: '💡',
    alarm_clock: '⏰',
    balance_scale: '⚖️',
    basketball: '🏀',
    blossom: '🌼',
    bomb: '💣',
    closed_umbrella: '🌂',
    lemon: '🍋',
    cheese: '🧀',
    cherries: '🍒',
    beers: '🍻',
    banana: '🍌',
    apple: '🍎',
    reminder_ribbon: '🎗',
    ring: '💍',
    trophy: '🏆',
    underage: '🔞',
    zap: '⚡️',
    zzz: '💤',
    airplane: '✈️',
    rocket: '🚀',
    car: '🚗',
    bus: '🚌',
    sunny: '☀️',
    sparkles: '✨',
    star: '⭐️',
    calendar: '📆',
    camera: '📷',
    computer: '💻',
    shit: '💩',
    christmas_tree: '🎄',
    ferris_wheel: '🎡',
    bee: '🐝',
    cactus: '🌵',
    fallen_leaf: '🍂',
    strawberry: '🍓',
    balloon: '🎈',
    '+1': '👍',
    '-1': '👎',
    crossed_fingers: '🤞',
    fist: '✊',
    point_right: '👉',
    raised_hand_with_fingers_splayed: '🖐',
    vulcan_salute: '🖖',
    boy: '👦',
    girl: '👧',
    bear: '🐻',
    cat: '🐱',
    cow: '🐮',
    dog: '🐶',
    fox_face: '🦊',
    lion: '🦁',
    mouse: '🐭',
    panda_face: '🐼',
    pig: '🐷',
    rabbit: '🐰',
    frog: '🐸',
    skull: '💀',
    kissing_heart: '😘',
    smile: '😄',
    smirk: '😏',
    wink: '😉',
    heart_eyes: '😍',
    laughing: '😆',
    yum: '😋',
    angry: '😠',
    unamused: '😒',
    worried: '😟',
    cry: '😢',
    face_with_head_bandage: '🤕',
    persevere: '😣',
    triumph: '😤',
    confused: '😕',
    confounded: '😖',
    expressionless: '😑',
    frowning_face: '☹️',
    zipper_mouth_face: '🤐',
    face_with_thermometer: '🤒',
    thinking: '🤔',
    astonished: '😲',
    dizzy_face: '😵',
    cold_sweat: '😰',
    fearful: '😨',
    flushed: '😳',
    grimacing: '😬',
    stuck_out_tongue: '😛',
    sleeping: '😴',
    sneezing_face: '🤧',
    sob: '😭',
    tired_face: '😫',
    weary: '😩',
    hushed: '😯',
    joy: '😂'
  }

  // 初始化VD
  const getNew = (initVal: string | undefined) => {
    vd.value = new Vditor(nodeId, {
      toolbar,
      lang: 'zh_CN',
      mode: 'ir',
      value: initVal,
      icon: 'material',
      height: window.innerHeight - 200,
      cache: {
        enable: false
      },
      outline: {
        enable: true,
        position: 'right'
      },
      debugger: true,
      typewriterMode: true,
      placeholder: "❤️你好！米虫！开始编写吧！❤️推荐您Ctrl+'或点击↕按钮打开全屏专注编写↗",
      cdn: '/static/lib/vditor',
      preview: {
        maxWidth: 920,
        actions: [],
        markdown: {
          toc: true,
          mark: true,
          footnotes: true,
          autoSpace: true
        },
        theme: {
          current: 'ant-design',
          path: '/static/lib/vditor'
        },
        hljs: {
          enable: true,
          lineNumber: true,
          defaultLang: '',
          style: 'vs'
        }
      },
      toolbarConfig: {
        pin: true
      },
      counter: {
        enable: true,
        type: 'text'
      },
      tab: '\t',
      hint: {
        emoji: emoji
      }
    })
  }

  const preViewSet = {
    mode: 'light' as 'light',
    theme: {
      current: 'ant-design',
      path: '/static/lib/vditor'
    },
    hljs: {
      enable: true,
      lineNumber: true,
      defaultLang: '',
      style: 'vs'
    },
    cdn: '/static/lib/vditor'
  }
  const sleep = (ms: number) => {
    return new Promise((resolve) => setTimeout(resolve, ms))
  }
  const hasNum = (haystack: string, needle: string) => {
    let regex = new RegExp(needle, 'g')
    let matches = haystack.match(regex)
    return matches ? matches.length : 0
  }
  // 模拟预览
  const toPreview = () => {
    var elements = document.querySelectorAll('#' + nodeId + ' .vditor-toolbar div button')
    var lastElement = elements[elements.length - 1]
    lastElement.dispatchEvent(new MouseEvent('click'))
  }
  // 获取提交结果
  const getResponse = async () => {
    // 读取预览HTML
    md.md = vd.value.getValue()
    var bashHtml = vd.value.getHTML()
    toPreview()
    var outline = document.querySelector('#' + nodeId + ' .vditor-outline')
    var myPreView = document.querySelector('#' + nodeId + ' .vditor-preview')
    if (myPreView != null) {
      await sleep(200)
      if (bashHtml.includes('<pre><code ')) {
        const hasCount = hasNum(bashHtml, '<pre><code ')
        let runRender = true
        while (runRender) {
          const renderHtml = myPreView.innerHTML
          if (renderHtml.includes('vditor-linenumber')) {
            // 优化预览比对细节，部分代码的Class排列有差异
            const hasNCount = hasNum(renderHtml, ' vditor-linenumber"')
            if (hasCount === hasNCount) {
              runRender = false
            }
          }
          await sleep(200)
        }
      }
      if (outline != null) {
        let runRender = true
        while (runRender) {
          let outHtml = outline.innerHTML
          if (!outHtml.includes('data-target-id="ir-')) {
            runRender = false
          }
          await sleep(200)
        }
        md.outline = outline.innerHTML
      }
      md.html = myPreView.innerHTML
      md.html = md.html.replaceAll('<p><img src="', '<p class="bg"><img src="')
      md.html = md.html.replaceAll(
        `<div class="vditor-copy"><textarea></textarea><span aria-label="复制" onmouseover="this.setAttribute('aria-label', '复制')" class="vditor-tooltipped vditor-tooltipped__w" onclick="this.previousElementSibling.select();document.execCommand('copy');this.setAttribute('aria-label', '已复制');this.previousElementSibling.blur()"><svg><use xlink:href="#vditor-icon-copy"></use></svg></span></div>`,
        ''
      )
      md.outline = md.outline?.replaceAll('<svg class="vditor-outline__action"><use xlink:href="#vditor-icon-down"></use></svg>', '')
      md.outline = md.outline?.replaceAll(
        '<svg class="vditor-outline__action" viewBox="0 0 32 32"><path d="M3.76 6.12l12.24 12.213 12.24-12.213 3.76 3.76-16 16-16-16 3.76-3.76z"></path></svg>',
        ''
      )
      md.outline = md.outline?.replaceAll('<svg></svg>', '')
      md.outline = md.outline?.replaceAll('vditor-outline__title', 'outline')
      md.outline = md.outline?.replaceAll('vditor-outline__content', 'outline_box')
    }
    // 处理图片关联
    let imgNodes = document.querySelectorAll('#' + nodeId + ' .vditor-preview img')
    var soucesList: number[] = []
    if (imgNodes) {
      imgNodes?.forEach((node) => {
        // ALT = SC-ID-NAME
        let alt = node.getAttribute('alt')
        if (alt) {
          const alts = alt.split('-')
          if (alts && alts.length > 1) {
            soucesList.push(Number(alts[1]))
          }
        }
      })
    }
    return soucesList
  }
  return {
    vd,
    md,
    getNew,
    getResponse
  }
}
