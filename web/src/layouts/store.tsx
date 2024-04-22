import NavigationHeader from '../components/navigation/Header.tsx'
import { Outlet } from 'react-router-dom'
import NavigationFooter from '../components/navigation/Footer.tsx'

export function StoreLayout() {
    return (
        <>
            <NavigationHeader />
            <Outlet />
            <NavigationFooter />
        </>
    )
}
