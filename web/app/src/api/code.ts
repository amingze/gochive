

declare interface codeMessageMapTypes {
    200: string;
    401: string;
    403: string;
    404: string;
    405: string;
    500: string;
    520: string;

    [key: string]: string;
}

const codeMessageMap: codeMessageMapTypes = {
    200: '请求成功',
    401: '[401]:账户未登录',
    403: '[403]:拒绝访问',
    404: '[404]:请求路径错误',
    405: '[405]:请求方法错误',
    500: '[500]:服务器错误',
    520: '[520]:服务器未初始化',
};

const showCodeMessage = (code: number | string): string => {
    return codeMessageMap[JSON.stringify(code)] || '网络连接异常,请稍后再试!';
};

export default showCodeMessage;
