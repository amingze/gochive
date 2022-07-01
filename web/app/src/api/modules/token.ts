import service from '@/api/http';
import userStore from "@/store/user";

const tokenApi = {
    // 验证登录实例
    postToken: (params: object) => service.post(`/tokens`, params).then((v) => {
        if (v.code == 0) {
            userStore().setUserInfo(v.data)
            ElMessage.success(`${'登陆成功'}`)
        }
        return v
    }),
    deleteToken: () => service.delete(`/tokens`,).then((v) => {
        userStore().clearUserInfo()
        ElMessage.success(`${'注销成功'}`)
        return v
    }),

};

export default tokenApi;
