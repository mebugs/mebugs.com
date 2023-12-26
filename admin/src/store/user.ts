import { defineStore } from 'pinia'
//import { logout as userLogout, getUserInfo } from '@/api/user'
//import { authLogin, AuthLoginReq } from '@/api/open/auth'
//import { setToken, clearToken } from '@/utils/auth'
//import { removeRouteListener } from '@/utils/route-listener'

interface UserState {
  name?: string
  avatar?: string
  permissions: string[] // 权限列表
}

const userStore = defineStore('user', {
  state: (): UserState => ({
    name: undefined,
    avatar: undefined,
    permissions: ['all']
  }),

  getters: {
    userInfo(state: UserState): UserState {
      return { ...state }
    }
  },

  actions: {
    // Set user's information
    setInfo(partial: Partial<UserState>) {
      this.$patch(partial)
    },

    // Reset user's information
    resetInfo() {
      this.$reset()
    },

    // Get user's information
    async info() {
      //const res = await getUserInfo()
      //this.setInfo(res.data)
    },

    // 账号登入
    async accountLogin(loginForm: any) {
      try {
        //const res = await authLogin(loginForm)
        // 设置登陆Token
        //setToken(res.data.token)
      } catch (err) {
        // 登陆失败清理Token
        // clearToken()
        throw err
      }
    },
    // 登出
    async logout() {
      try {
        // await userLogout()
      } finally {
        this.resetInfo()
        //clearToken()
        //removeRouteListener()
      }
    }
  }
})

export default userStore
