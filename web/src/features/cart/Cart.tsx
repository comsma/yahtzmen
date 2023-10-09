import { PropsWithChildren } from 'react'
import {
    MinusCircleIcon,
    PlusCircleIcon,
    TrashIcon,
} from '@heroicons/react/20/solid'
import { useAppDispatch, useAppSelector } from '../../app/hooks.ts'
import {
    decrementProductCount,
    deleteProduct,
    incrementProductCount,
} from './cartSlice.ts'

export default function CartList() {
    const cart = useAppSelector((state) => state.cart)
    return (
        <div className={'col-span-full flex flex-col gap-y-5 lg:col-span-2'}>
            {cart?.products.map((product) => (
                <CartProduct key={product.id} product={product} />
            ))}
        </div>
    )
}

export function CartProduct({
    product,
}: PropsWithChildren<{ product: cartProductI }>) {
    const dispatch = useAppDispatch()
    return (
        <div
            className={
                'grid grid-cols-2 grid-rows-4 gap-x-4 rounded-xl bg-gray-200 p-5 md:grid-rows-2'
            }
            key={product.id}
        >
            <div
                className={
                    'col-span-1 row-span-1 row-start-1 flex flex-col justify-items-start gap-x-4 sm:col-span-4'
                }
            >
                <p className={'text-lg font-bold'}>{product.name}</p>
                <p>${product.price}</p>
            </div>
            <div
                className={
                    'col-span-2 row-start-4 flex flex-row lg:col-span-4 lg:col-start-1 lg:row-start-2'
                }
            >
                <div
                    className={
                        'flow flex flex-1 flex-row place-items-center gap-x-6'
                    }
                >
                    <MinusCircleIcon
                        aria-label={`decrement ${product.name} quantity`}
                        onClick={() =>
                            dispatch(decrementProductCount(product.id))
                        }
                        className={'h-[19px]'}
                    />
                    <p className={'rounded-md bg-golden-rod px-5 text-lg'}>
                        {product.quantity}
                    </p>
                    <PlusCircleIcon
                        aria-label={`increment ${product.name} quantity`}
                        onClick={() =>
                            dispatch(incrementProductCount(product.id))
                        }
                        className={'h-[19px]'}
                    />
                    <TrashIcon
                        aria-label={`remove ${product.name} from cart`}
                        className={'mx-7 ml-auto h-[19px] text-red-500'}
                        onClick={() => dispatch(deleteProduct(product.id))}
                    />
                </div>
                <div className={'col-span-1'}></div>
            </div>

            <div
                className={
                    'relative col-span-2 col-start-1 row-span-2 row-start-2 mx-auto aspect-square h-40 self-center overflow-clip rounded-md lg:col-span-1 lg:col-start-5 lg:row-span-2 lg:row-start-1'
                }
            >
                {/*<Image*/}
                {/*    className={'object-contain'}*/}
                {/*    fill={true}*/}
                {/*    src={item.itemImg}*/}
                {/*    alt={item.name}*/}
                {/*/>*/}
            </div>
        </div>
    )
}
