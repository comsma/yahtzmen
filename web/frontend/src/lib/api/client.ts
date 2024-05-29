import createClient, { Middleware } from 'openapi-fetch'
import { paths } from './v1'
import { notFound } from '@tanstack/react-router'

const throwOnError: Middleware = {
    async onResponse(res) {
        if (res.status >= 404) {
            throw notFound()
        }
        return undefined
    },
}

const client = createClient<paths>({ baseUrl: '/api/v1' })

client.use(throwOnError)

export default client
