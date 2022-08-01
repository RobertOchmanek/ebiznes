import React from 'react';
import { getLogoutUrl } from '../api/getLogoutUrl';

export default function Header(props) {

    const { backendAddress, numCartItems, loggedIn, user } = props;

    const onLogout = () => {
        getLogoutUrl(backendAddress, user.ID).then(logoutUrl => window.open(logoutUrl, "_self"));
    }

    return (
        <header className='row block center'>
            <div>
                <h1>Robert's Store</h1>
                {loggedIn ? (<button onClick={() => onLogout()}>Log out</button>) : (<></>)}
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