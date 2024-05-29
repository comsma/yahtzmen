import { createFileRoute, Link } from '@tanstack/react-router'
import { Container } from '@/routes/-components/container.tsx'
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from '@/components/card.tsx'
import { Button, buttonVariants } from '@/components/ui/button.tsx'
import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from '@/components/ui/table.tsx'
import { useMutation, useQuery } from '@tanstack/react-query'
import { listStoreProductsQueryOptions } from '@/hooks/queries.ts'
import { Skeleton } from '@/components/skeleton.tsx'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu.tsx'
import { MoreHorizontal } from 'lucide-react'
import { adminDeleteProductMutationOptions } from '@/lib/query/admin/product.ts'
import { queryClient } from '@/main.tsx'

export const Route = createFileRoute('/_admin/admin/(product)/product/')({
    component: ProductComponent,
    loader: (opts) =>
        opts.context.queryClient.ensureQueryData(
            listStoreProductsQueryOptions()
        ),
})

function ProductComponent() {
    const productsQuery = useQuery(listStoreProductsQueryOptions())
    const products = productsQuery.data

    if (productsQuery.isLoading) {
        return <div>Loading...</div>
    }
    return (
        <Container className={'pt-8'}>
            <Card>
                <CardHeader>
                    <CardTitle>Products</CardTitle>
                    <CardDescription
                        className={'flex w-full flex-row justify-end'}
                    >
                        <Link
                            to={'/admin/product/new'}
                            search={{}}
                            className={buttonVariants({ variant: 'secondary' })}
                        >
                            Add a product
                        </Link>
                    </CardDescription>
                    <CardContent>
                        {productsQuery.isLoading && (
                            <div className={'flex flex-col gap-8'}>
                                <Skeleton className={'h-[40px] w-full'} />
                                <Skeleton className={'h-[40px] w-full'} />
                                <Skeleton className={'h-[40px] w-full'} />
                                <Skeleton className={'h-[40px] w-full'} />
                            </div>
                        )}
                        <Table>
                            <TableHeader>
                                <TableRow>
                                    <TableHead></TableHead>
                                    <TableHead>Name</TableHead>
                                    <TableHead>Price</TableHead>
                                    <TableHead>More</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {products?.map((product) => (
                                    <TableRow>
                                        <TableCell></TableCell>
                                        <TableCell>
                                            <Link
                                                to={
                                                    '/admin/product/$productId/view'
                                                }
                                                params={{
                                                    productId: product.id,
                                                }}
                                                className={buttonVariants({
                                                    variant: 'link',
                                                    className: 'text-secondary',
                                                })}
                                            >
                                                {product.name}
                                            </Link>
                                        </TableCell>
                                        <TableCell>{product.price}</TableCell>
                                        <TableCell>
                                            <ProductActions id={product.id} />
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </CardContent>
                </CardHeader>
            </Card>
        </Container>
    )
}

function ProductActions(props: { id: string }) {
    const adminDeleteProductMutation = useMutation(
        adminDeleteProductMutationOptions()
    )
    return (
        <DropdownMenu>
            <DropdownMenuTrigger>
                <Button variant={'ghost'} className={'h-8 w-8 p-0'}>
                    <span className={'sr-only'}>Actions</span>
                    <MoreHorizontal className={'h-4 w-4'}></MoreHorizontal>
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align={'end'}>
                <DropdownMenuLabel>Actions</DropdownMenuLabel>
                <DropdownMenuItem>
                    <Button
                        onClick={() => {
                            adminDeleteProductMutation.mutate(props.id)
                            queryClient.invalidateQueries(
                                listStoreProductsQueryOptions()
                            )
                        }}
                        variant={'ghost'}
                        className={'text-destructive'}
                        size={'sm'}
                    >
                        Delete
                    </Button>
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}
