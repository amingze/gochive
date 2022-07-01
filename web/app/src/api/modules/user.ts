import service from '@/api/http';

const userApi = {
    // 验证登录实例
    getUserInfo: () => service.get(`/user`),

};

export default userApi;
