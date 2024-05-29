import { createFileRoute, Link } from '@tanstack/react-router'
import {
    adminDeleteProductMutationOptions,
    adminGetProductQueryOptions,
} from '@/lib/query/admin/product.ts'
import { useMutation, useSuspenseQuery } from '@tanstack/react-query'
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
    Dialog,
    DialogContent,
    DialogDescription,
    DialogTitle,
    DialogTrigger,
    DialogClose,
    DialogFooter,
    DialogHeader,
} from '@/components/ui/dialog.tsx'

export const Route = createFileRoute(
    '/_admin/admin/(product)/product/$productId/_details/view'
)({
    component: ProductViewComponent,
})

function ProductViewComponent() {
    const { viewProductOptions } = Route.useRouteContext()
    const productQuery = useSuspenseQuery(viewProductOptions)
    const product = productQuery.data
    return (
        <Container className={'pt-8'}>
            <Card>
                <CardHeader>
                    <CardTitle>{product?.data?.name}</CardTitle>
                    <CardDescription
                        className={'flex flex-row justify-end gap-4'}
                    >
                        <Link
                            className={buttonVariants({ variant: 'outline' })}
                            to={`/admin/product/$productId/edit`}
                            params={{ productId: product?.data?.id }}
                        >
                            Edit
                        </Link>
                        <DeleteDialog
                            id={product?.data?.id || ''}
                            name={product?.data?.name || ''}
                        />
                    </CardDescription>
                </CardHeader>
                <CardContent className={'flex flex-col gap-4'}>
                    <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                        <dt className="text-sm font-medium leading-6 text-secondary">
                            Description
                        </dt>
                        <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                            {product?.data?.description}
                        </dd>
                    </div>
                    <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                        <dt className="text-sm font-medium leading-6 text-secondary">
                            Features
                        </dt>
                        <dd className="mt-1 whitespace-pre-wrap text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                            {product?.data?.features}
                        </dd>
                    </div>
                    <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                        <dt className="text-sm font-medium leading-6 text-secondary">
                            Notes
                        </dt>
                        <dd className="mt-1 whitespace-pre-wrap text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                            {product?.data?.notes}
                        </dd>
                    </div>
                    <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
                        <dt className="whitespace-pre-wrap text-sm font-medium leading-6 text-secondary">
                            Dimensions
                        </dt>
                        <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                            {product?.data?.dimensions}
                        </dd>
                    </div>
                </CardContent>
            </Card>
        </Container>
    )
}

function DeleteDialog(props: { id: string; name: string }) {
    const adminDeleteProductMutation = useMutation(
        adminDeleteProductMutationOptions()
    )
    return (
        <Dialog>
            <DialogTrigger asChild>
                <Button variant="destructive">Delete</Button>
            </DialogTrigger>
            <DialogContent className="w-11/12 sm:max-w-md">
                <DialogHeader>
                    <DialogTitle>Delete</DialogTitle>
                    <DialogDescription>
                        Are you sure you want to delete.
                    </DialogDescription>
                </DialogHeader>
                <div className="flex items-center space-x-2"></div>
                <Button
                    variant={'destructive'}
                    type="submit"
                    size="sm"
                    className="px-3"
                >
                    Confirm
                </Button>

                <DialogFooter>
                    <DialogClose asChild>
                        <Button
                            type="button"
                            variant="outline"
                            className={'w-full'}
                        >
                            Cancel
                        </Button>
                    </DialogClose>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}
