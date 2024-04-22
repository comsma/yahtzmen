import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'

export default function Home() {
    const [products, setProducts] = useState<ProductWithImage[]>([])

    useEffect(() => {
        const fetchProducts = async () => {
            const response = await fetch('/api/v1/products')
            const data = await response.json()
            setProducts(data)
        }

        fetchProducts()
    }, [])
    return (
        <>
            <div className={'z-0 bg-blue-800 py-10'}>
                <div className={'mx-auto max-w-5xl px-5 font-lora lg:px-0'}>
                    <div className={'overflow-hidden rounded-md'}>
                        <div className="relative hidden aspect-[16/9] sm:flex">
                            <img
                                src={'/images/home/thoughtfullyDesigned.jpg'}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                        <div className="visible relative aspect-[11/16] sm:hidden">
                            <img
                                src={
                                    '/images/home/thoughtfullyDesigned-mobile.jpg'
                                }
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                    </div>
                    <div className={'my-10 overflow-hidden rounded-md'}>
                        <div className="relative hidden aspect-[16/9] sm:flex">
                            <img
                                src={'/images/home/handCrafted.jpg'}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                        <div className="visible relative aspect-[11/16] sm:hidden">
                            <img
                                src={'/images/home/handCrafted-mobile.jpg'}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                    </div>
                    <div id={'products'}>
                        <ProductGrid products={products} />
                    </div>
                    <div
                        className={'my-10 overflow-hidden rounded-md  '}
                        id={'mission'}
                    >
                        <div className="relative hidden aspect-[16/9] sm:block">
                            <img
                                src={'OurMission'}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                        <div className="visible relative aspect-[11/16] sm:hidden">
                            <img
                                src={'OurMissionMobile'}
                                alt={
                                    'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                }
                            />
                        </div>
                    </div>
                    <div className={'mt-10 overflow-hidden rounded-md'}>
                        <div
                            className={'hover:cursor-pointer hover:opacity-75 '}
                            id={'collaborate'}
                            onClick={() =>
                                (window.location.href =
                                    'mailto:info@yahtzmen.com')
                            }
                        >
                            <div className="relative hidden aspect-[16/9] sm:block">
                                <img
                                    src={'LetCollaborate'}
                                    alt={
                                        'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                    }
                                />
                            </div>
                            <div className="visible relative aspect-[8/16] sm:hidden">
                                <img
                                    src={'LetCollaborateMobile'}
                                    alt={
                                        'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                                    }
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

function ProductGrid({ products }: { products: ProductWithImage[] }) {
    return (
        <div className={'my-20 rounded-md bg-white px-10 py-10'}>
            <h2 className={'py-10 text-3xl'}>Products</h2>
            <div className={'grid grid-cols-1 gap-5 lg:grid-cols-3 '}>
                {products.map((product) => (
                    <div key={product.id}>
                        <Link to={`/product/${product.id}`}>
                            <div>
                                <div className="relative aspect-[1/1] overflow-hidden rounded-md bg-gray-200 group-hover:opacity-75">
                                    <img
                                        src={product.primaryImage}
                                        alt={product.name}
                                    />
                                </div>
                                <div className="flex items-baseline justify-between py-4">
                                    <h3 className="text-base text-gray-700">
                                        {product.name}
                                    </h3>
                                    <p className="text-sm font-medium text-gray-900">
                                        ${product.price}
                                    </p>
                                </div>
                            </div>
                        </Link>
                    </div>
                ))}
            </div>
        </div>
    )
}
