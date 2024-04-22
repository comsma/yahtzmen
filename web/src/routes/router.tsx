import { BrowserRouter } from 'react-router-dom'
import { AdminRoutes } from './admin.tsx'
import { StoreRoutes } from './store.tsx'

export function Router() {
    return (
        <BrowserRouter>
            <StoreRoutes />
            <AdminRoutes />
        </BrowserRouter>
    )
}
