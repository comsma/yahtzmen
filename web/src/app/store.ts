import { configureStore } from '@reduxjs/toolkit'
import { cartSlice } from '../features/cart/cartSlice.ts'

const store = configureStore({
    reducer: {
        cart: cartSlice.reducer,
    },
})

export default store

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
