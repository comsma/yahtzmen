import { createFileRoute } from '@tanstack/react-router'
import { getStoreProductByIdQueryOptions } from '@/hooks/queries.ts'
import { useSuspenseQuery } from '@tanstack/react-query'
import { Button } from '@/components/ui/button.tsx'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/card.tsx'
import { Container } from '@/routes/-components/container.tsx'
import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from '@/components/accordion.tsx'
import {
    Carousel,
    CarouselContent,
    CarouselItem,
    CarouselNext,
    CarouselPrevious,
} from '@/components/carousel.tsx'
import { useBasketStore } from '@/lib/store/basket.ts'

export const Route = createFileRoute('/_store/(product)/product/$productId')({
    parseParams: (params) => ({
        productId: params.productId,
    }),
    loader: (opts) =>
        opts.context.queryClient.ensureQueryData(
            getStoreProductByIdQueryOptions(opts.params.productId)
        ),
    component: ProductById,
})

function ProductById() {
    const addBasketItem = useBasketStore((state) => state.addItem)
    const { productId } = Route.useParams()
    const getStoreProductByIdQuery = useSuspenseQuery(
        getStoreProductByIdQueryOptions(productId)
    )
    const product = getStoreProductByIdQuery.data
    return (
        <Container
            className={'flex flex-col items-center gap-4 pt-8 lg:flex-row'}
        >
            {product?.images && product?.images.length > 0 && (
                <div className={'mx-auto max-w-lg px-12'}>
                    <Carousel>
                        <CarouselContent>
                            {product?.images.map((image) => (
                                <CarouselItem>
                                    <div
                                        className={
                                            'flex aspect-square items-center justify-center overflow-hidden rounded-md'
                                        }
                                    >
                                        <img
                                            src={image.url}
                                            className={'object-center'}
                                        />
                                    </div>
                                </CarouselItem>
                            ))}
                        </CarouselContent>
                        <CarouselPrevious />
                        <CarouselNext />
                    </Carousel>
                </div>
            )}

            <Card className={'max-w-2xl flex-1'}>
                <CardHeader>
                    <CardTitle>{product?.name}</CardTitle>
                </CardHeader>
                <CardContent className={'flex flex-col gap-4'}>
                    <p className="py-3 text-3xl tracking-tight text-gray-900">
                        {product?.price?.toLocaleString('en-US', {
                            style: 'currency',
                            currency: 'USD',
                        })}
                    </p>
                    <Button
                        size={'lg'}
                        className={'w-full'}
                        onClick={() =>
                            addBasketItem({
                                id: product?.id || '',
                                name: product?.name || '',
                                price: product?.price || 0,
                                quantity: 1,
                            })
                        }
                    >
                        Add to Cart
                    </Button>
                    <p className="whitespace-pre-line text-base text-gray-900">
                        {product?.description}
                    </p>
                    <Accordion type={'single'} collapsible>
                        <AccordionItem value={'features'}>
                            <AccordionTrigger>Features</AccordionTrigger>
                            <AccordionContent className={'whitespace-pre-wrap'}>
                                {product?.features}
                            </AccordionContent>
                        </AccordionItem>
                        <AccordionItem value={'dimensions'}>
                            <AccordionTrigger>Dimensions</AccordionTrigger>
                            <AccordionContent className={'whitespace-pre-wrap'}>
                                {product?.dimensions}
                            </AccordionContent>
                        </AccordionItem>
                    </Accordion>
                </CardContent>
            </Card>
        </Container>
    )
}
