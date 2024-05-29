import { createRootRouteWithContext } from '@tanstack/react-router'
import { Outlet } from '@tanstack/react-router'
import { QueryClient } from '@tanstack/react-query'
import { Toaster } from '@/components/ui/sonner'

export const Route = createRootRouteWithContext<{
    queryClient: QueryClient
}>()({
    component: RootComponent,
})

function RootComponent() {
    return (
        <>
            <Outlet />
            <Toaster />
        </>
    )
}
