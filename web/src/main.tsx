import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { QueryClient } from '@tanstack/react-query'
import { Router } from './routes/router.tsx'

const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <Router />
        {/*<Provider store={store}>*/}
        {/*    <QueryClientProvider client={queryClient}>*/}
        {/*        */}
        {/*    </QueryClientProvider>*/}
        {/*</Provider>*/}
    </React.StrictMode>
)
