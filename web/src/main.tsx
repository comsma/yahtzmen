import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import DefaultLayout from './layouts/Default.tsx'
import HomePage from './pages'
import { Provider } from 'react-redux'
import store from './app/store.ts'
import PrivacyPage from './pages/privacy'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

const queryClient = new QueryClient()

const router = createBrowserRouter([
    {
        element: <DefaultLayout />,
        children: [
            {
                path: '/',
                element: <HomePage />,
            },
            {
                path: '/product/:id',
            },
            {
                path: '/privacy',
                element: <PrivacyPage />,
            },
        ],
    },
])
ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <Provider store={store}>
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
            </QueryClientProvider>
        </Provider>
    </React.StrictMode>
)
