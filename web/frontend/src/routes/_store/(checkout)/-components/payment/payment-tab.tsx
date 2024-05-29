import { Card, CardContent, CardHeader } from '@/components/card.tsx'
import { StripeElements } from '@/routes/_store/(checkout)/-components/payment/stripe-elements.tsx'

export function PaymentTab() {
    return (
        <Card className={''}>
            <CardHeader></CardHeader>
            <CardContent>
                <StripeElements />
            </CardContent>
        </Card>
    )
}
