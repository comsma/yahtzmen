import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_admin/admin/')({
    component: () => <div>Hello /_admin/admin/!</div>,
})
