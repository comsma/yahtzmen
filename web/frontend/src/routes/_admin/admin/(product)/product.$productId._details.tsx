import { createFileRoute, Link, Outlet } from '@tanstack/react-router'
import {
    Breadcrumb,
    BreadcrumbItem,
    BreadcrumbLink,
    BreadcrumbList,
    BreadcrumbPage,
    BreadcrumbSeparator,
} from '@/components/ui/breadcrumbs.tsx'
import { Container } from '@/routes/-components/container.tsx'
import {
    adminGetProductImagesQueryOptions,
    adminGetProductQueryOptions,
} from '@/lib/query/admin/product.ts'
import { useSuspenseQuery } from '@tanstack/react-query'

export const Route = createFileRoute(
    '/_admin/admin/(product)/product/$productId/_details'
)({
    component: ProductDetailsLayout,
    beforeLoad: ({ params: { productId } }) => {
        return {
            viewProductOptions: adminGetProductQueryOptions(productId),
            viewProductImagesOptions:
                adminGetProductImagesQueryOptions(productId),
        }
    },
    loader: ({ context }) => {
        context.queryClient.ensureQueryData(context.viewProductOptions)
        context.queryClient.ensureQueryData(context.viewProductImagesOptions)
    },
})

function ProductDetailsLayout() {
    const { viewProductOptions } = Route.useRouteContext()
    const productQuery = useSuspenseQuery(viewProductOptions)
    const product = productQuery.data
    return (
        <Container className={'pt-8'}>
            <Breadcrumb>
                <BreadcrumbList>
                    <BreadcrumbItem>
                        <BreadcrumbLink asChild>
                            <Link to={'/admin/product'}>Products</Link>
                        </BreadcrumbLink>
                    </BreadcrumbItem>
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                        <BreadcrumbPage>{product?.data?.name}</BreadcrumbPage>
                    </BreadcrumbItem>
                </BreadcrumbList>
            </Breadcrumb>
            <Outlet />
        </Container>
    )
}
