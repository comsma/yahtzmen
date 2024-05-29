import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import {
    createRouter,
    ErrorComponent,
    RouterProvider,
} from '@tanstack/react-router'
import { routeTree } from './routeTree.gen.ts'

export const queryClient = new QueryClient()
// Set up a Router instance
const router = createRouter({
    routeTree,
    defaultPreload: 'intent',
    context: {
        queryClient,
    },
    defaultErrorComponent: ({ error }) => <ErrorComponent error={error} />,
})

// Register things for typesafety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router
    }
}

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <QueryClientProvider client={queryClient}>
            <RouterProvider router={router} context={{}} />
        </QueryClientProvider>
    </React.StrictMode>
)
