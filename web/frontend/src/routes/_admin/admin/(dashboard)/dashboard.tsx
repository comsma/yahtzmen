import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_admin/admin/(dashboard)/dashboard')({
  component: () => <div>Hello /_admin/admin/(dashboard)/dashboard!</div>
})