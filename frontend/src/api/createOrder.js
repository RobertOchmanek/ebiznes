export async function createOrder(cartItems, userId, ammount) {

    const orderItems = []

    cartItems.forEach((cartItem) => {
        const orderItem = {
            ProductId: cartItem.ID,
            Quantity: cartItem.Quantity
        }
        orderItems.push(orderItem)
    })

    const data = await fetch('http://localhost:8080/orders', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ UserId: userId, Ammount: ammount, OrderItems: orderItems })
    });
    
    return data.json();
}