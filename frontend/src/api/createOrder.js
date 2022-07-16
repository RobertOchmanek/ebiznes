export async function createOrder(cartItems, paymentType) {

    const orderItems = []

    cartItems.forEach((cartItem) => {
        const orderItem = {
            ProductId: cartItem.ID,
            Quantity: cartItem.Quantity
        }
        orderItems.push(orderItem)
    })

    let paymentTypeValue

    switch(paymentType) {
        case "Blik":
          paymentTypeValue = 0;
          break;
        case "Credit card":
          paymentTypeValue = 1;
          break;
        case "Bank transfer":
           paymentTypeValue = 2;
           break;
        default:
          paymentTypeValue = 0;
      }

    const payemnt = {
        Accepted: true,
        PaymentType: paymentTypeValue
    }

    const data = await fetch('http://localhost:8080/orders', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        //NOTE: user ID is hardcoded until login is implemented
        body: JSON.stringify({ UserId: 1, Payment: payemnt, OrderItems: orderItems})
    });
    
    return data.json();
}