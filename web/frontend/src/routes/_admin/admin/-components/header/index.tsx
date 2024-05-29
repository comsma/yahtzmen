import { AppNavigation } from '@/routes/_admin/admin/-components/header/app-navigation.tsx'
import { cn } from '@/lib/utils.ts'

export default function AdminHeader() {
    return (
        <>
            <AppNavigation
                className={cn(
                    'sticky top-0 z-20 w-full border-b bg-secondary text-secondary-foreground md:shadow-sm'
                )}
            />
        </>
    )
}
