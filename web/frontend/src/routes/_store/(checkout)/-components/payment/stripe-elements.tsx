import {
    loadStripe,
    StripeElementsOptions,
    StripePaymentElementOptions,
} from '@stripe/stripe-js'
import { useEffect, useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { createCheckoutSessionQueryOptions } from '@/hooks/queries.ts'
import {
    Elements,
    PaymentElement,
    useElements,
    useStripe,
} from '@stripe/react-stripe-js'
import { Button } from '@/components/ui/button.tsx'
import { toast } from 'sonner'

export function StripeElements() {
    const stripePromise = loadStripe('')
    const [clientSecret, setClientSecret] = useState('')
    const checkoutSessionQuery = useQuery(createCheckoutSessionQueryOptions())

    useEffect(() => {
        if (
            checkoutSessionQuery.data &&
            checkoutSessionQuery.data.clientSecret
        ) {
            setClientSecret(checkoutSessionQuery.data.clientSecret)
        }
    }, [checkoutSessionQuery.data])

    const options: StripeElementsOptions = {
        clientSecret: clientSecret,
        fonts: [
            {
                cssSrc: 'https://fonts.googleapis.com/css2?family=Lora:ital,wght@0,400..700;1,400..700&display=swap',
            },
        ],
        appearance: {
            theme: 'stripe',
            variables: {
                colorPrimary: '#09142A',
                colorBackground: '#ffffff',
                colorText: '#30313d',
                colorDanger: '#E05D52',
                fontFamily: 'Lora, serif',
                spacingUnit: '4px',
                borderRadius: '4px',
            },
        },
    }

    return (
        <>
            {clientSecret && (
                <Elements stripe={stripePromise} options={options}>
                    <PaymentForm />
                </Elements>
            )}
        </>
    )
}

function PaymentForm() {
    const stripe = useStripe()
    const elements = useElements()

    const [message, setMessage] = useState<string | null>(null)
    const [isLoading, setIsLoading] = useState(false)

    useEffect(() => {
        if (!stripe) {
            return
        }

        const clientSecret = new URLSearchParams(window.location.search).get(
            'payment_intent_client_secret'
        )

        if (!clientSecret) {
            return
        }

        stripe.retrievePaymentIntent(clientSecret).then(({ paymentIntent }) => {
            switch (paymentIntent?.status) {
                case 'succeeded':
                    setMessage('Payment succeeded!')
                    break
                case 'processing':
                    setMessage('Your payment is processing.')
                    break
                case 'requires_payment_method':
                    setMessage(
                        'Your payment was not successful, please try again.'
                    )
                    break
                default:
                    setMessage('Something went wrong.')
                    break
            }
        })
    }, [stripe])

    const handleSubmit = async (e: { preventDefault: () => void }) => {
        e.preventDefault()

        if (!stripe || !elements) {
            // Stripe.js hasn't yet loaded.
            // Make sure to disable form submission until Stripe.js has loaded.
            return
        }

        setIsLoading(true)

        const { error } = await stripe.confirmPayment({
            elements,
            confirmParams: {
                // Make sure to change this to your payment completion page
                return_url: 'http://localhost:3000',
            },
        })

        // This point will only be reached if there is an immediate error when
        // confirming the payment. Otherwise, your customer will be redirected to
        // your `return_url`. For some payment methods like iDEAL, your customer will
        // be redirected to an intermediate site first to authorize the payment, then
        // redirected to the `return_url`.
        if (error.type === 'card_error' || error.type === 'validation_error') {
            toast.error(error?.message || 'An error occurred.')
        } else {
            toast.error('Something went wrong. Please try again.')
        }

        setIsLoading(false)
    }

    const paymentElementOptions: StripePaymentElementOptions = {
        layout: {
            type: 'accordion',
            defaultCollapsed: false,
            radios: false,
            spacedAccordionItems: true,
        },
    }

    return (
        <form
            className={'flex flex-col gap-4'}
            id="payment-form"
            onSubmit={handleSubmit}
        >
            <PaymentElement
                id="payment-element"
                options={paymentElementOptions}
            />
            <Button
                className={'w-full'}
                disabled={isLoading || !stripe || !elements}
                id="submit"
            >
                <span id="button-text">
                    {isLoading ? (
                        <div className="spinner" id="spinner"></div>
                    ) : (
                        'Pay now'
                    )}
                </span>
            </Button>
            {/* Show any error or success messages */}
            {message && <div id="payment-message">{message}</div>}
        </form>
    )
}
