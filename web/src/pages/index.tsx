import { PropsWithChildren } from 'react'
import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import axios from 'axios'

interface productI {
    id: string
    name: string
    price: number
}

export default function HomePage() {
    const { data } = useQuery({
        queryKey: ['products'],
        queryFn: () =>
            axios
                .get('http://localhost:8080/api/product')
                .then((res) => res.data),
    })

    console.log(data)

    return (
        <div className={'z-0 bg-blue-800 py-10'}>
            <div className={'mx-auto max-w-5xl px-5 font-lora lg:px-0'}>
                <div className={'overflow-hidden rounded-lg  '}>
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
                            src={'/images/home/thoughtfullyDesigned-mobile.jpg'}
                            alt={
                                'Thoughtfully Designed, Bespoke Furnishings for Modern Living.'
                            }
                        />
                    </div>
                </div>
                <div className={'my-10 overflow-hidden rounded-lg'}>
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
                <div id={'products'} className={'overflow-hidden rounded-lg'}>
                    {/*<ProductGrid />*/}
                </div>
                <div
                    className={'my-10 overflow-hidden rounded-lg  '}
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
                <div className={'mt-10 overflow-hidden rounded-lg'}>
                    <div
                        className={'hover:cursor-pointer hover:opacity-75 '}
                        id={'collaborate'}
                        onClick={() =>
                            (window.location.href = 'mailto:info@yahtzmen.com')
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
    )
}

function ProductGrid({
    products,
}: PropsWithChildren<{ products: [productI] }>) {
    return (
        <div className={'grid grid-cols-1 lg:grid-cols-3'}>
            {products.map((product) => (
                <Product key={product.id} product={product} />
            ))}
        </div>
    )
}

function Product({ product }: PropsWithChildren<{ product: productI }>) {
    return (
        <div className="">
            <Link to={`/product/${product.id}`}>
                <div>
                    <div className="relative aspect-[1/1] overflow-hidden rounded-md bg-gray-200 group-hover:opacity-75">
                        <img src={''} alt={product.name} />
                    </div>
                    <div className="mt-4 flex justify-between">
                        <div>
                            <h3 className="font-book-antiqua text-sm text-gray-700">
                                <span
                                    aria-hidden="true"
                                    className="absolute inset-0"
                                />
                                {product.name}
                            </h3>
                        </div>
                        <p className="text-sm font-medium text-gray-900">
                            ${product.price}
                        </p>
                    </div>
                </div>
            </Link>
        </div>
    )
}
