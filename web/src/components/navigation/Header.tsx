import { Link } from 'react-router-dom'
import { ShoppingCartIcon } from '@heroicons/react/20/solid'

const navigation = [
    { name: 'Products', href: '/#product' },
    { name: 'Our Mission', href: '/#mission' },
    { name: 'Lets Collaborate', href: '/#collaborate' },
]
export default function NavigationHeader() {
    return (
        <nav
            className="relative w-full bg-blue-800 px-4 py-4 font-lora text-golden-rod sm:px-6 lg:px-8"
            aria-label="Top"
        >
            <div className="grid grid-cols-3 lg:grid-cols-7">
                <div className={'hidden lg:col-span-2 lg:block'}>
                    <Link to={'/'}>
                        <div className={'relative h-8'}>
                            <img
                                className={'h-full w-full object-contain'}
                                src={'/images/common/yahtzmen-name.png'}
                                alt="Yahtzmen YF logo"
                            />
                        </div>
                    </Link>
                </div>
                <div className={'col-span-1 col-start-4'}>
                    <Link to={'/'}>
                        <div className={'relative mx-auto h-10 w-auto'}>
                            <img
                                className={
                                    'mx-auto h-full w-auto object-contain'
                                }
                                src={'/images/common/yahtzmen-badge.png'}
                                alt="Yahtzmen YF logo"
                            />
                        </div>
                    </Link>
                </div>
                <div className={'col-span-2'}>
                    <div className={'flex flex-row justify-center'}>
                        {navigation.map((link) => (
                            <div className={'items-center'} key={link.name}>
                                <Link
                                    to={link.href}
                                    className="hover:text-orange-yellow-crayola ml-5 text-xs font-light md:text-base xl:text-lg"
                                >
                                    {link.name}
                                </Link>
                            </div>
                        ))}
                    </div>
                </div>
                <div className={'col-span-1 col-start-7'}>
                    <Link to={'/Cart'}>
                        <div className={'flex justify-end'}>
                            <div className={'relative inline-block'}>
                                <ShoppingCartIcon className={'h-7 w-7'} />
                                <span
                                    className={
                                        'absolute right-0 top-0 flex h-2 w-2 items-center justify-center rounded-full border-2 border-blue-800 bg-golden-rod p-[5px] text-[7px] font-bold text-blue-800'
                                    }
                                >
                                    {/*{cart.length}*/}
                                </span>
                            </div>
                        </div>
                    </Link>
                </div>
            </div>
        </nav>
    )
}
