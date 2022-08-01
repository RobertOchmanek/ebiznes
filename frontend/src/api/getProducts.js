export async function getProducts(backendAddress) {

  const data = await fetch(backendAddress + '/products');
  return data.json();
}