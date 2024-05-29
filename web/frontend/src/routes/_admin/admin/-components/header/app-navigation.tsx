import {
    AnyRouter,
    Link,
    LinkProps,
    RegisteredRouter,
    RoutePaths,
} from '@tanstack/react-router'
import { LinkComponentProps } from '@/lib/types/router.ts'
import { cn } from '@/lib/utils.ts'

type Props = React.ComponentPropsWithoutRef<'nav'>
export function AppNavigation(props: Props) {
    const { className = '', ...navProps } = props

    return (
        <nav className={cn(className)} {...navProps}>
            <ul className="relative mx-auto -mb-px flex max-w-[1700px] space-x-4 overflow-x-auto px-4 [-ms-overflow-style:none] [scrollbar-width:none] sm:space-x-0.5 md:px-12 [&::-webkit-scrollbar]:hidden">
                <Link className={'mr-8 self-center'} to={'/'}>
                    <img
                        className={'h-6 w-auto object-contain'}
                        width={150}
                        src={'/images/common/yahtzmen-badge.webp'}
                        alt="Yahtzmen YF logo"
                    />
                </Link>
                <AppNavigationLink
                    name="Dashboard"
                    props={{
                        to: '/admin',
                        activeOptions: {
                            exact: true,
                            includeSearch: false,
                        },
                    }}
                />
                <AppNavigationLink
                    name="Products"
                    props={{
                        to: '/admin/product',
                        activeOptions: {
                            exact: true,
                            includeSearch: false,
                        },
                    }}
                />
            </ul>
        </nav>
    )
}

function AppNavigationLink<
    TRouter extends AnyRouter = RegisteredRouter,
    TFrom extends RoutePaths<TRouter['routeTree']> | string = string,
    TTo extends string = '',
    TMaskFrom extends RoutePaths<TRouter['routeTree']> | string = TFrom,
    TMaskTo extends string = '',
>(props: {
    name: string
    props: LinkProps<TRouter, TFrom, TTo, TMaskFrom, TMaskTo> &
        LinkComponentProps<'a'>
}) {
    const { name, props: linkProps } = props
    return (
        <li>
            <Link
                className={cn(
                    'inline-block whitespace-nowrap border-b-4 py-4 leading-none transition-all sm:px-4'
                )}
                activeProps={{
                    className: cn('border-primary font-semibold'),
                }}
                inactiveProps={{
                    className: cn(
                        'border-transparent hover:border-foreground/20'
                    ),
                }}
                {...linkProps}
            >
                {name}
            </Link>
        </li>
    )
}
