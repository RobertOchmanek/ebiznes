export async function updateCart(cart) {

    //TODO: always one request behind
    const cartItems = []

    cart.forEach((cartItem) => {
        cartItems.push(cartItem)
    })

    const data = await fetch('http://localhost:8080/cart', {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        //NOTE: user ID is hardcoded until login is implemented
        body: JSON.stringify({ UserId: 1, CartItems: cartItems })
    });
    
    return await data.json();
}