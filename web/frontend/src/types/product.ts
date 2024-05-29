interface Product {
    id: number
    name: string
    price: number
    description: string
    features: string
    dimensions: string
    productImages: ProductImageWithUrl[]
}

interface ProductWithImage extends Product {
    primaryImage: string
}
