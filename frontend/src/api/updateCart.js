export async function updateCart(backendAddress, cart, userId) {

    const cartItems = []

    cart.forEach((cartItem) => {
        cartItems.push(cartItem)
    })

    const data = await fetch(backendAddress + '/cart', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ UserId: userId, CartItems: cartItems })
    });
    
    return data.json();
}