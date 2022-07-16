export async function getProducts() {

  const data = await fetch('http://localhost:8080/products');
  return data.json();
  }