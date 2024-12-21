import { ref } from 'vue'
// 选择工具类(单选)
// 适用于弹窗列表选择
export function RadioList(formData: any, idKey: string, valKey: string) {
  // 路由添加层
  const check = ref(false)
  // 路由选择对象
  const radioData: any = ref({
    selectedKey: undefined,
    selectedItem: {}
  })
  // 初始化选择对象
  function init() {
    radioData.value = {
      selectedKey: undefined,
      selectedItem: {}
    }
  }
  // 打开选择画面
  function open() {
    init()
    check.value = true
  }
  // 取消选择画面
  function cancel() {
    init()
    check.value = false
  }
  // 确认选择
  function done() {
    // 与内部选择组件进行双向同步
    formData.value[valKey] = radioData.value.selectedItem
    formData.value[idKey] = radioData.value.selectedKey
    return true
  }
  // 子组件选择(传入方法)
  function doSelect(item: any, key: any) {
    radioData.value.selectedKey = item
    radioData.value.selectedItem = key
  }

  return {
    check,
    radioData,
    open,
    cancel,
    done,
    doSelect
  }
}
