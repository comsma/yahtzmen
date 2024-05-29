import { create } from 'zustand'
import { toast } from 'sonner'
import { createJSONStorage, persist } from 'zustand/middleware'

interface BasketItem {
    id: string
    name: string
    price: number
    quantity: number
}
interface BasketState {
    items: BasketItem[]
    addItem: (item: BasketItem) => void
    removeItem: (itemId: string) => void
    getSubtotal: () => number
    clearItems: () => void
}

export const useBasketStore = create<BasketState>()(
    persist(
        (set, get) => ({
            items: [],
            addItem: (item) =>
                set((state) => {
                    const existingItem = state.items.find(
                        (i) => i.id === item.id
                    )
                    if (existingItem) {
                        return {
                            items: state.items.map((i) => {
                                if (i.id === item.id) {
                                    toast.success(
                                        `Updated quantity of ${item.name} to ${i.quantity + item.quantity}`
                                    )
                                    return {
                                        ...i,
                                        quantity: i.quantity + item.quantity,
                                    }
                                }
                                return i
                            }),
                        }
                    }
                    toast.success(`${item.name} added to basket`)
                    return {
                        items: [...state.items, item],
                    }
                }),
            removeItem: (itemId) =>
                set((state) => ({
                    items: state.items.filter((i) => i.id !== itemId),
                })),
            getSubtotal: () =>
                get().items.reduce(
                    (acc, item) => acc + item.price * item.quantity,
                    0
                ),
            clearItems: () => set({ items: [] }),
        }),
        {
            name: 'basket',
            storage: createJSONStorage(() => sessionStorage),
        }
    )
)
