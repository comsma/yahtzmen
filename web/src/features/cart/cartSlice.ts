import { createSlice, PayloadAction } from '@reduxjs/toolkit'

const initialState: cartI = {
    products: [],
}
export const cartSlice = createSlice({
    name: 'cartSlice',
    initialState,
    reducers: {
        createProduct: (state, action: PayloadAction<cartProductI>) => {
            //check if product exists
            const index = state.products.findIndex(
                (obj) => obj.id === action.payload.id
            )
            if (index > -1) {
                alert('Item is already in the cart')
            } else {
                state.products.push(action.payload)
            }
        },
        deleteProduct: (state, action: PayloadAction<String>) => {
            //get index of product
            const index = state.products.findIndex(
                (obj) => obj.id === action.payload
            )
            state.products.splice(index, 1)
        },
        incrementProductCount: (state, action: PayloadAction<String>) => {
            //get index of product
            const index = state.products.findIndex(
                (obj) => obj.id === action.payload
            )
            //increments product's quantity
            state.products[index].quantity += 1
        },
        decrementProductCount: (state, action: PayloadAction<String>) => {
            //get index of product
            const index = state.products.findIndex(
                (obj) => obj.id === action.payload
            )
            //decrements product quantity
            if (state.products[index].quantity <= 1) {
                state.products.splice(index, 1)
            } else {
                state.products[index].quantity -= 1
            }
        },
        emptyCart: (state) => {
            state.products.length = 0
        },
    },
})
export const {
    createProduct,
    deleteProduct,
    incrementProductCount,
    decrementProductCount,
    emptyCart,
} = cartSlice.actions

export default cartSlice.reducer
