import React from 'react';
import { getLoginUrl } from '../api/getLoginUrl';
import Header from './Header';

export default function LoginPage(props) {

    const { backendAddress, loggedIn } = props

    const onLogin = () => {
        getLoginUrl(backendAddress).then(loginUrl => window.open(loginUrl, "_self"));
      }

    return(
        <div>
            <Header loggedIn={loggedIn}></Header>
            <div className='login-wrapper'>
                <h2>Please log in using available OAuth2 providers:</h2>
                <div>
                    <button type="submit" onClick={() => onLogin()}>GitHub</button>
                </div>
            </div>
        </div>
    )
}