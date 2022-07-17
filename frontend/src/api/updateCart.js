export async function updateCart(cart, userId) {

    const cartItems = []

    cart.forEach((cartItem) => {
        cartItems.push(cartItem)
    })

    const data = await fetch('http://localhost:8080/cart', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ UserId: userId, CartItems: cartItems })
    });
    
    return data.json();
}