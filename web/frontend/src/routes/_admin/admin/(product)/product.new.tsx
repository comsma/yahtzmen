import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from '@/components/ui/form.tsx'

import { Button } from '@/components/ui/button.tsx'
import { Input } from '@/components/ui/input.tsx'
import { Container } from '@/routes/-components/container.tsx'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/card.tsx'
import { Textarea } from '@/components/ui/textarea.tsx'
import { useMutation } from '@tanstack/react-query'
import { adminCreateProductMutationOptions } from '@/lib/query/admin/product.ts'

const formSchema = z.object({
    name: z.string().min(2).max(50),
    description: z.string().min(2).max(1000),
    price: z.coerce.number().min(0),
    features: z.string().max(1000),
    notes: z.string().max(1000),
    dimensions: z.string().max(500),
    isLive: z.boolean(),
})

export const Route = createFileRoute('/_admin/admin/(product)/product/new')({
    component: NewProductComponent,
})

function NewProductComponent() {
    return (
        <Container className={'pt-8'}>
            <Card>
                <CardHeader>
                    <CardTitle>Create a new product</CardTitle>
                </CardHeader>
                <CardContent>
                    <NewProductForm />
                </CardContent>
            </Card>
        </Container>
    )
}

function NewProductForm() {
    const createProductMutation = useMutation(
        adminCreateProductMutationOptions()
    )
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: '',
            description: '',
            price: 0,
            features: '',
            notes: '',
            dimensions: '',
            isLive: true,
        },
    })

    function onSubmit(values: z.infer<typeof formSchema>) {
        createProductMutation.mutate({
            name: values.name,
            description: values.description,
            price: values.price,
            features: values.features,
            notes: values.notes,
            dimensions: values.dimensions,
            isLive: values.isLive,
        })
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                <FormField
                    control={form.control}
                    name="name"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Name</FormLabel>
                            <FormControl>
                                <Input {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="description"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Description</FormLabel>
                            <FormControl>
                                <Textarea {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="price"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Price</FormLabel>
                            <FormControl>
                                <Input type="number" {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="notes"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Notes</FormLabel>
                            <FormControl>
                                <Textarea {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="features"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Features</FormLabel>
                            <FormControl>
                                <Textarea {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="dimensions"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Dimensions</FormLabel>
                            <FormControl>
                                <Textarea {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <div className={'flex flex-row justify-between '}>
                    <Button
                        variant="outline"
                        onClick={() => {
                            form.setValue('isLive', false)
                        }}
                        type="submit"
                    >
                        Save as draft
                    </Button>
                    <Button
                        type="submit"
                        onClick={() => {
                            form.setValue('isLive', true)
                        }}
                    >
                        Publish Product
                    </Button>
                </div>
            </form>
        </Form>
    )
}
