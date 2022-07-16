import React from 'react';

export default function Header(props) {

    const { numCartItems } = props;

    return (
        <header className='row block center'>
            <div>
                <a href='#/'>
                    <h1>Robert's Store</h1>
                </a>
            </div>
            <div>
                <a href='#/cart'>

                    Cart { ' ' }
                    {numCartItems ? (<button className='badge'>{numCartItems}</button>) : ('')}

                </a> <a href='#/signIn'>Sign In</a>
            </div>
        </header>
    );
}