interface ProductImage {
    id: number
    url: string
    order: number
}

interface ProductImageWithUrl extends ProductImage {
    url: string
}
