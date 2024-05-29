import { createFileRoute, Outlet } from '@tanstack/react-router'
import AdminHeader from '@/routes/_admin/admin/-components/header'
import { DialogOverlay } from '@/components/ui/dialog.tsx'

export const Route = createFileRoute('/_admin')({
    component: AdminLayout,
})

function AdminLayout() {
    return (
        <>
            <AdminHeader />
            <Outlet />
        </>
    )
}
