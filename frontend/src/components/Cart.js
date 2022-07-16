import React from 'react';
import { useState } from 'react';

export default function Cart(props) {

    const getDefaultPaymentType = () => {
        const defaultPaymentType = "Blik";
        return defaultPaymentType;
      };
    

    const [paymentType, setPaymentType] = useState(getDefaultPaymentType);

    const { cartItems, onAdd, onRemove, onOrderPlaced } = props;
    const cartValue = cartItems.reduce((accumulator, currentItem) => accumulator + (currentItem.Price * currentItem.Quantity), 0);
    const taxValue = cartValue * 0.23;
    const shipping= cartValue > 2000 ? 0 : 50;
    const totalPrice = cartValue + taxValue + shipping;

    const onPaymentTypeChanged = (e) => {
        setPaymentType(e.target.value);
    }

    return (
        <aside className='block col-1'>
            <h2>Cart items</h2>
            <div>
                {cartItems.length === 0 && <div>Cart is empty</div>}
            </div>
            
            {cartItems.map((cartItem) => (
                <div key={cartItem.ID} className='row'>
                    <div className='col-2'>{cartItem.Name}</div>
                    <div className='col-2'>
                        <button onClick={() => onAdd(cartItem)} className='add'>+</button>
                        <button onClick={() => onRemove(cartItem)} className='remove'>-</button>
                    </div>
                    <div className='col-2 text-right'>
                        {cartItem.Quantity} x ${cartItem.Price.toFixed(2)}
                    </div>
                </div>
            ))}

            {cartItems.length !== 0 && (
                <>
                    <hr></hr>

                    <div className='row'>
                        <div className='col-2'>Cart value:</div>
                        <div className='col-1 text-right'>${cartValue.toFixed(2)}</div>
                    </div>
                    <div className='row'>
                        <div className='col-2'>Tax value:</div>
                        <div className='col-1 text-right'>${taxValue.toFixed(2)}</div>
                    </div>
                    <div className='row'>
                        <div className='col-2'>Shipping:</div>
                        <div className='col-1 text-right'>${shipping.toFixed(2)}</div>
                    </div>
                    <div className='row'>
                        <div className='col-2'><strong>Total price:</strong></div>
                        <div className='col-1 text-right'><strong>${totalPrice.toFixed(2)}</strong></div>
                    </div>

                    <hr></hr>

                    <div className='row'>
                        <div className='col-2'>Payment type:</div>
                        <div className='col-1 text-right'>
                            <select value={paymentType} onChange={onPaymentTypeChanged}>
                                <option value="Blik">Blik</option>
                                <option value="Credit card">Credit card</option>
                                <option value="Bank transfer">Bank transfer</option>
                            </select>
                        </div>
                    </div>

                    <hr></hr>

                    <div className='row'>
                        <button onClick={() => {onOrderPlaced(paymentType); setPaymentType(getDefaultPaymentType())}}>Place an order</button>
                    </div>
                </>
            )}
        </aside>
    );
}