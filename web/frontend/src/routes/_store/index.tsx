import { createFileRoute, Link } from '@tanstack/react-router'
import { listStoreProductsQueryOptions } from '@/hooks/queries.ts'
import { useQuery } from '@tanstack/react-query'
import thoughtfullyDesigned from '@/assets/thoughtfully-designed.webp'
import thoughtfullyDesignedM from '@/assets/thoughtfully-designed-m.webp'
import handCrafted from '@/assets/hand-crafted.webp'
import handCraftedM from '@/assets/hand-crafted-m.webp'
import ourMission from '@/assets/our-mission.webp'
import ourMissionM from '@/assets/our-mission-m.webp'
import letsCollaborate from '@/assets/lets-collaborate.webp'
import letsCollaborateM from '@/assets/lets-collaborate-m.webp'
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
    CardTitle,
} from '@/components/card.tsx'
import { Skeleton } from '@/components/skeleton.tsx'

export const Route = createFileRoute('/_store/')({
    component: IndexComponent,
    beforeLoad: () => {
        return {
            viewProductOptions: listStoreProductsQueryOptions(),
        }
    },
    loader: ({ context }) => {
        context.queryClient.ensureQueryData(context.viewProductOptions)
    },
})

function IndexComponent() {
    return (
        <>
            <div className={'bg-secondary py-10'}>
                <div
                    className={
                        'mx-auto flex max-w-5xl flex-col gap-10 px-5 font-lora lg:px-0'
                    }
                >
                    <div className={'overflow-hidden rounded-md'}>
                        <div className="aspect-[16/9] max-sm:hidden">
                            <img
                                loading={'lazy'}
                                width={2501}
                                className={'object-contain'}
                                src={thoughtfullyDesigned}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                        <div className="aspect-[11/16] sm:hidden">
                            <img
                                loading={'lazy'}
                                width={1170}
                                className={'object-contain'}
                                src={thoughtfullyDesignedM}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                    </div>
                    <div className={'overflow-hidden rounded-md'}>
                        <div className="aspect-[16/9] max-sm:hidden">
                            <img
                                loading={'lazy'}
                                width={2501}
                                className={'object-contain'}
                                src={handCrafted}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                        <div className="aspect-[11/16] sm:hidden">
                            <img
                                loading={'lazy'}
                                width={1170}
                                className={'object-contain'}
                                src={handCraftedM}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                    </div>

                    <Card id={'products'}>
                        <CardHeader>
                            <CardTitle>Products</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <ProductGrid />
                        </CardContent>
                    </Card>
                    <div
                        className={'overflow-hidden rounded-md'}
                        id={'mission'}
                    >
                        <div className="aspect-[16/9] max-sm:hidden">
                            <img
                                loading={'lazy'}
                                width={2501}
                                className={'object-contain'}
                                src={ourMission}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                        <div className="aspect-[11/16] sm:hidden">
                            <img
                                loading={'lazy'}
                                width={1170}
                                className={'object-contain'}
                                src={ourMissionM}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                    </div>
                    <div
                        className={'overflow-hidden rounded-md'}
                        id={'collaborate'}
                    >
                        <a href={'mailto:info@yahtzmen.com'}>
                            <div className="aspect-[16/9] max-sm:hidden">
                                <img
                                    loading={'lazy'}
                                    width={2501}
                                    className={'object-contain'}
                                    src={letsCollaborate}
                                    alt={
                                        'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                    }
                                />
                            </div>
                            <div className="aspect-[11/16] sm:hidden">
                                <img
                                    loading={'lazy'}
                                    width={1170}
                                    className={'object-contain'}
                                    src={letsCollaborateM}
                                    alt={
                                        'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                    }
                                />
                            </div>
                        </a>
                    </div>
                </div>
            </div>
        </>
    )
}

export default function ProductGrid() {
    const { viewProductOptions } = Route.useRouteContext()
    const productsQuery = useQuery(viewProductOptions)

    if (productsQuery.isPending) {
        return (
            <div
                className={
                    'grid grid-cols-1 place-items-center gap-4 lg:grid-cols-3'
                }
            >
                <Skeleton className={'h-[200px] w-[200px] rounded-xl'} />
                <Skeleton className={'h-[200px] w-[200px] rounded-xl'} />
                <Skeleton className={'h-[200px] w-[200px] rounded-xl'} />
            </div>
        )
    }

    if (productsQuery.isError) {
        return <span>Error: {productsQuery.error?.message}</span>
    }

    const products = productsQuery.data

    return (
        <div
            className={
                'grid grid-cols-1 place-items-center gap-4 lg:grid-cols-3'
            }
        >
            {products?.map((product) => (
                <Link
                    to={`/product/$productId`}
                    search={{}}
                    params={{ productId: product.id }}
                    aria-label={`View ${product.name}`}
                >
                    <Card key={product.id}>
                        <CardHeader>
                            <div className="mx-auto aspect-[1/1] w-[200px] overflow-hidden rounded-md bg-gray-200 group-hover:opacity-75">
                                <img
                                    className={'object-contain'}
                                    src={product.image?.url || ''}
                                    alt={product.name}
                                />
                            </div>
                        </CardHeader>
                        <CardFooter className={'grid grid-cols-2 gap-2'}>
                            <h3 className="col-span-2">{product.name}</h3>
                            <p className="col-start-2 place-self-end text-sm font-bold">
                                ${product.price}
                            </p>
                        </CardFooter>
                    </Card>
                </Link>
            ))}
        </div>
    )
}
