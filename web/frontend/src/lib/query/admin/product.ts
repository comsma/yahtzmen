import client from '@/lib/api/client.ts'
import { paths } from '@/lib/api/v1'
import { queryOptions } from '@tanstack/react-query'

const SEGMENT = 'admin-product'

export function adminGetProductImagesQueryOptions(productId: string) {
    return queryOptions({
        queryKey: [SEGMENT, 'images'],
        queryFn: async ({ signal }) => {
            const { data } = await client.GET('/admin/product/{id}/image', {
                params: { path: { id: productId } },
                signal,
            })
            return data
        },
    })
}
export function adminCreateProductMutationOptions() {
    return {
        mutationKey: [SEGMENT, 'create'],
        mutationFn: (
            body: paths['/admin/product']['post']['requestBody']['content']['application/json']
        ) => client.POST('/admin/product', { body: body }),
    } as const
}

export function adminDeleteProductMutationOptions() {
    return {
        mutationKey: [SEGMENT, 'delete'],
        mutationFn: (id: string) =>
            client.DELETE('/admin/products/{id}', {
                params: { path: { id: id } },
            }),
    } as const
}

export function adminGetProductQueryOptions(productId: string) {
    return queryOptions({
        queryKey: [SEGMENT, 'get'],
        queryFn: async ({ signal }) =>
            client.GET('/admin/products/{id}', {
                params: { path: { id: productId } },
                signal,
            }),
    })
}
