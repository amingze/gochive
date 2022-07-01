import userApi from "@/api/modules/user";

const userStore = defineStore('counter', {
    state: () => {
        return {userInfo: {}, isLogin: false}
    },
    getters: {},
    actions: {
        clearUserInfo(user ?: any) {
            this.userInfo={}
        },
        setUserInfo(user ?: any) {
            this.userInfo = user
        },
        getUserInfo(): any {
            if (!this.userInfo || Object.keys(this.userInfo).length == 0) {
                userApi.getUserInfo().then((v) => {
                    this.userInfo = v.data
                })
            }
            return this.userInfo
        },
    },
})
export default userStore