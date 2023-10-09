interface CartInterface {
    items: CartItemInterface[]
}

interface CartItemInterface {
    id: string
    name: string
    qty: number
    imgUrl: string
    price: bigint
}

export type { CartInterface, CartItemInterface }
