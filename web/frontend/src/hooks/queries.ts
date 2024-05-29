import { queryOptions } from '@tanstack/react-query'
import client from '../lib/api/client.ts'

export const createCheckoutSessionQueryOptions = () =>
    queryOptions({
        queryKey: ['checkout-session'],
        queryFn: async ({ signal }) => {
            const { data } = await client.POST('/checkout/session', {
                signal,
                body: {
                    products: [],
                },
            })
            return data
        },
    })
export const listStoreProductsQueryOptions = () =>
    queryOptions({
        queryKey: ['store-products'],
        queryFn: async ({ signal }) => {
            const { data } = await client.GET('/store/products', { signal })
            return data
        },
    })

export const getStoreProductByIdQueryOptions = (productId: string) =>
    queryOptions({
        queryKey: ['store-product', { id: productId }],
        queryFn: async ({ signal }) => {
            const { data } = await client.GET(`/store/products/{id}`, {
                params: { path: { id: productId } },
                signal,
            })
            return data
        },
    })
