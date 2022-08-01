export async function getCartItems(backendAddress, userId) {

    const data = await fetch(backendAddress + '/cartItems/' + userId);
    return data.json();
}