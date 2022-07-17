import { useEffect, useState } from 'react';
import Cart from './components/Cart';
import Header from './components/Header';
import LoginPage from './components/LoginPage'
import Products from './components/Products';
import { createOrder } from './api/createOrder';
import { getCartItems } from './api/getCartItems';
import { getProducts } from './api/getProducts';
import { getUser } from './api/getUser';
import { updateCart } from './api/updateCart';

function App() {

  const [userToken, setUserToken] = useState();
  const [user, setUser] = useState();
  const [products, setProducts] = useState([]);
  const [cartItems, setCartItems] = useState([]);

  useEffect(() => {

    let mounted = true;

    if (!userToken) {

      const urlParams = new URLSearchParams(window.location.search);
      setUserToken(urlParams.get('user_token'))

    } else {

      getUser(userToken).then(fetchedUser => {
        if (mounted) {
          setUser(fetchedUser)
          return fetchedUser.ID
        }
      }).then(userId => {
        getCartItems(userId).then(fetchedCartItems => {
          if (mounted) {
            setCartItems(fetchedCartItems)
          }
        });
      });

      getProducts().then(fetchedProducts => {
        if (mounted) {
          setProducts(fetchedProducts)
        }
      });

    }

    return () => mounted = false;
  }, [userToken]);

  const onAdd = (product) => {

    let cartItemsCopy = [...cartItems]
    const existingItem = cartItemsCopy.find(cartItem => cartItem.ID === product.ID)

    if (existingItem) {
      cartItemsCopy = cartItemsCopy.map((cartItem) => 
          cartItem.ID === product.ID ? {...existingItem, Quantity: existingItem.Quantity + 1} : cartItem
      );
    } else {
      cartItemsCopy.push({...product, Quantity: 1})
    }

    setCartItems(cartItemsCopy)
    updateCart(cartItemsCopy, user.ID);
  };

  const onRemove = (product) => {

    let cartItemsCopy = [...cartItems]
    const existingItem = cartItemsCopy.find((cartItem) => cartItem.ID === product.ID);

    if (existingItem.Quantity === 1) {
      cartItemsCopy = cartItemsCopy.filter((cartItem) => cartItem.ID !== product.ID);
    } else {
        cartItemsCopy = cartItemsCopy.map((cartItem) => 
          cartItem.ID === product.ID ? {...existingItem, Quantity: existingItem.Quantity - 1} : cartItem
      );
    }

    setCartItems(cartItemsCopy)
    updateCart(cartItemsCopy, user.ID);
  };

  const onOrderPlaced = (paymentType) => {

    createOrder(cartItems, paymentType, user.ID);
    setCartItems([]);
    alert("Order successfully placed!")
  }

  if(!userToken) {

    return (
      <div className='App'>
        <LoginPage></LoginPage>
      </div>
    );

  } else {

    return(
      <div className='App'>
        <Header numCartItems={cartItems.length} showBadge={true} setUserToken={setUserToken}></Header>
        <div className='row'>
          <Products onAdd={onAdd} products={products}></Products>
          <Cart onAdd={onAdd} onRemove={onRemove} onOrderPlaced={onOrderPlaced} cartItems={cartItems}></Cart>
        </div>
      </div>
    );

  }
}

export default App;
