import { Route, Routes } from 'react-router-dom'
import Home from '../components/store/home'
import { StoreLayout } from '../layouts/store.tsx'
import PrivacyPage from '../components/store/privacy'
import Product from '../components/store/product'

export function StoreRoutes() {
    return (
        <Routes>
            <Route path={'/'} element={<StoreLayout />}>
                <Route index element={<Home />} />
                <Route path="products" element={<></>} />
                <Route path={'product/:id'} element={<Product />} />
                <Route path="cart" element={<></>} />
                <Route path="checkout" element={<></>} />
                <Route path={'privacy'} element={<PrivacyPage />} />
            </Route>
        </Routes>
    )
}
