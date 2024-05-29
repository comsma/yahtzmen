import { createFileRoute, Outlet } from '@tanstack/react-router'
import { Header } from '@/packages/ui/navigation'
import Footer from '@/components/navigation/Footer.tsx'

export const Route = createFileRoute('/_store')({
    component: StoreLayout,
})

function StoreLayout() {
    return (
        <>
            <Header></Header>
            <Outlet />
            <Footer />
        </>
    )
}
