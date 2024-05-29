import { Link, useNavigate } from '@tanstack/react-router'
import { LucideMenu, LucideShoppingCart } from 'lucide-react'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuItem,
    DropdownMenuLinkItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button, buttonVariants } from '@/components/ui/button.tsx'
import { useBasketStore } from '@/lib/store/basket.ts'
import { useState } from 'react'

const navigation = [
    { name: 'Products', href: '/', hash: 'products' },
    { name: 'Our Mission', href: '/', hash: 'mission' },
    { name: 'Lets Collaborate', href: '/', hash: 'collaborate' },
]

export function Header() {
    const basketItems = useBasketStore((state) => state.items)
    const navigator = useNavigate()
    const [open, setOpen] = useState(false)
    return (
        <nav
            className="w-full bg-secondary px-4 py-6 [-ms-overflow-style:none] [scrollbar-width:none] sm:space-x-0.5 sm:px-6 md:px-12 lg:px-8 [&::-webkit-scrollbar]:hidden"
            aria-label="Top"
        >
            <div className="mx-auto flex max-w-7xl flex-row items-center justify-between px-8">
                <div className={'max-lg:hidden'}>
                    <Link to={'/'}>
                        <div className={''}>
                            <img
                                width={2949}
                                className={'h-6 w-auto object-contain'}
                                src={'/images/common/yahtzmen-name.webp'}
                                alt="Yahtzmen YF logo"
                            />
                        </div>
                    </Link>
                </div>
                <div className={'justify-self-center'}>
                    <img
                        className={'h-6 w-auto object-contain'}
                        width={150}
                        src={'/images/common/yahtzmen-badge.webp'}
                        alt="Yahtzmen YF logo"
                    />
                </div>
                <div className={'flex flex-row items-center justify-between'}>
                    <div className={'flex flex-row max-lg:hidden '}>
                        {navigation.map((link) => (
                            <div className={'items-center'} key={link.name}>
                                <Link
                                    to={link.href}
                                    className={buttonVariants({
                                        variant: 'link',
                                    })}
                                >
                                    {link.name}
                                </Link>
                            </div>
                        ))}
                    </div>

                    <Link
                        className={buttonVariants({
                            variant: 'link',
                            size: 'icon',
                            className: 'ml-6 hidden lg:inline-flex',
                        })}
                        to={'/checkout'}
                    >
                        <div className={'relative inline-flex'}>
                            <LucideShoppingCart className={'h-7 w-7'} />

                            <span
                                className={
                                    'absolute -right-2 -top-2 flex h-4 w-4 items-center justify-center rounded-full border-2 border-primary  bg-secondary p-2 text-xs font-bold text-primary'
                                }
                            >
                                {basketItems.length}
                            </span>
                        </div>
                    </Link>
                </div>
                <div className={'text-primary lg:hidden'}>
                    <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                            <LucideMenu className={'h-6 w-6'} />
                        </DropdownMenuTrigger>
                        <DropdownMenuContent
                            className={'flex w-full flex-col'}
                            align={'end'}
                        >
                            {navigation.map((link) => (
                                <DropdownMenuItem
                                    key={link.href}
                                    className={buttonVariants({
                                        variant: 'link',

                                        className: 'text-secondary',
                                    })}
                                    asChild
                                >
                                    <Link to={link.href} hash={link.hash}>
                                        {link.name}
                                    </Link>
                                </DropdownMenuItem>
                            ))}

                            <DropdownMenuSeparator />

                            <DropdownMenuItem>
                                cart
                                {/*<Link*/}
                                {/*    search={{}}*/}
                                {/*    to={'/checkout'}*/}
                                {/*    className={buttonVariants({*/}
                                {/*        variant: 'link',*/}
                                {/*        className: 'text-secondary',*/}
                                {/*    })}*/}
                                {/*    aria-label={`Cart with ${basketItems.length} items`}*/}
                                {/*>*/}
                                {/*    Cart*/}
                                {/*</Link>*/}
                            </DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                </div>
            </div>
        </nav>
    )
}
