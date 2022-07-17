export async function getCartItems(userId) {

    const data = await fetch('http://localhost:8080/cartItems/' + userId);
    return data.json();
}