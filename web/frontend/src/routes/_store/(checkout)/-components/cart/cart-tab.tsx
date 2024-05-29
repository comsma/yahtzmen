import { useBasketStore } from '@/lib/store/basket.ts'
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
} from '@/components/card.tsx'
import { Button } from '@/components/ui/button.tsx'
import { Separator } from '@/components/ui/separator.tsx'

export interface CartTabProps {
    onCartStageComplete: () => void
    onCompleted: () => void
}
export function CartTab(props: CartTabProps) {
    const { onCompleted, onCartStageComplete } = props

    const basketItems = useBasketStore((state) => state.items)
    return (
        <Card>
            <CardHeader></CardHeader>
            <CardContent>
                {basketItems.length === 0 && (
                    <p className={'py-20 text-center'}>Your cart is empty</p>
                )}
                {basketItems.map((item) => (
                    <BasketItem item={item} />
                ))}
            </CardContent>
            <CardFooter>
                <Button
                    onClick={() => {
                        onCartStageComplete()
                        onCompleted()
                    }}
                    className={'w-full'}
                >
                    Checkout{' '}
                </Button>
            </CardFooter>
        </Card>
    )
}

function BasketItem({
    item,
}: {
    item: { id: string; name: string; price: number }
}) {
    const removeBasketItem = useBasketStore((state) => state.removeItem)
    return (
        <div className={'flex flex-row items-center gap-4 animate-out'}>
            <div className={''}>
                <img src={''} className={'h-20 w-20'} />
            </div>
            <Separator orientation={'vertical'} />
            <div className={'flex flex-col gap-2'}>
                <div className={'flex flex-row items-baseline gap-4'}>
                    <p>{item.name}</p>
                    <p className={'text-lg font-bold'}>{item.price}</p>
                </div>
                <Separator />
                <div>
                    <Button
                        variant={'outline'}
                        className={'text-destructive'}
                        size={'sm'}
                        onClick={() => removeBasketItem(item.id)}
                        aria-label={`Remove ${item.name} from cart`}
                    >
                        Remove
                    </Button>
                </div>
            </div>
        </div>
    )
}
