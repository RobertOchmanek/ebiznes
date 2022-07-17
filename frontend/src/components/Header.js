import React from 'react';

export default function Header(props) {

    const { numCartItems } = props;

    return (
        <header className='row block center'>
            <div>
                <h1>Robert's Store</h1>
            </div>
            {numCartItems > 0 ? 
                (<div>
                    Items in cart: <button className='badge'>{numCartItems}</button>
                </div>) : 
                (<></>)
            }
        </header>
    );
}