import { createFileRoute } from '@tanstack/react-router'
import { Container } from '@/routes/-components/container.tsx'
import {
    Tabs,
    TabsContent,
    TabsList,
    TabsTrigger,
} from '@/components/ui/tabs.tsx'
import { ReactNode, useMemo, useState } from 'react'
import { CartTab } from '@/routes/_store/(checkout)/-components/cart/cart-tab.tsx'
import { PaymentTab } from '@/routes/_store/(checkout)/-components/payment/payment-tab.tsx'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/card.tsx'

interface TCheckoutCompletionStages {
    cart: boolean
    payment: boolean
}
const defaultCompletionStages: TCheckoutCompletionStages = {
    cart: true,
    payment: false,
}
export const Route = createFileRoute('/_store/(checkout)/checkout')({
    component: Checkout,
})

function Checkout() {
    const StageKeys = {
        cart: 'cart',
        payment: 'payment',
    }
    // const basketItems = useBasketStore((state) => state.items)
    // const subtotal = useBasketStore((state) => state.getSubtotal())
    const [creationStageComplete, setCreationStageComplete] = useState(
        defaultCompletionStages
    )

    const [currentTab, setCurrentTab] = useState(StageKeys.cart)

    const tabsConfig = useMemo(() => {
        const tabs: { id: string; label: string; component: ReactNode }[] = []

        const cartTab = {
            id: StageKeys.cart,
            label: 'Cart',
            component: (
                <CartTab
                    onCartStageComplete={() => {
                        setCreationStageComplete({
                            ...creationStageComplete,
                            payment: true,
                        })
                    }}
                    onCompleted={() => {
                        setCurrentTab(StageKeys.payment)
                    }}
                />
            ),
        }
        const checkoutTab = {
            id: StageKeys.payment,
            label: 'Payment',
            component: <PaymentTab />,
        }
        tabs.push(cartTab)
        tabs.push(checkoutTab)
        return tabs
    }, [creationStageComplete])

    return (
        <Container
            className={'flex flex-col justify-center gap-8 pt-8 lg:flex-row'}
        >
            <Tabs
                value={currentTab}
                onValueChange={setCurrentTab}
                className={'w-full'}
            >
                <TabsList className={'flex w-full flex-row'}>
                    {tabsConfig.map((tab) => (
                        <TabsTrigger
                            disabled={
                                !creationStageComplete[
                                    tab.id as keyof TCheckoutCompletionStages
                                ]
                            }
                            className={'flex-1 flex-grow'}
                            value={tab.id}
                        >
                            {tab.label}
                        </TabsTrigger>
                    ))}
                </TabsList>
                {tabsConfig.map((tab, idx) => (
                    <TabsContent key={`tab-content-${idx}`} value={tab.id}>
                        {tab.component}
                    </TabsContent>
                ))}
            </Tabs>
            <Card className={'flex-'}>
                <CardHeader>
                    <CardTitle>Summary</CardTitle>
                </CardHeader>
                <CardContent></CardContent>
            </Card>
        </Container>
    )
}
