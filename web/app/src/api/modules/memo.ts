import service from '@/api/http';

const memoApi = {
    addMemo: (params: object) => service.post(`/memo`, params).then((v) => {
        return v
    }),
    listMemo: () => service.get(`/memo/list`),
    deleteMemo: (id: string) => service.delete("/memo/" + id),
}
export default memoApi