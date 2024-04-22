import { Route, Routes } from 'react-router-dom'

export function AdminRoutes() {
    return (
        <Routes>
            <Route path="/admin" element={<></>}>
                <Route index element={<></>} />
                <Route path="users" element={<></>} />
                <Route path="products" element={<></>} />
                <Route path={'product/:id'} element={<></>} />
                <Route path="orders" element={<></>} />
            </Route>
        </Routes>
    )
}
