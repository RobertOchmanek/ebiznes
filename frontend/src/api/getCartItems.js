export async function getCartItems() {

    //NOTE: user ID is hardcoded until login is implemented
    const data = await fetch('http://localhost:8080/cartItems/1');
    return await data.json();
}