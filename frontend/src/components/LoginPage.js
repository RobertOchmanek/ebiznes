import React from 'react';
import { getLoginUrl } from '../api/getLoginUrl';
import Header from './Header';

export default function LoginPage() {

    const onLogin = (oauthProvider) => {
        getLoginUrl(oauthProvider).then(loginUrl => window.open(loginUrl, "_self"));
      }

    return(
        <div>
            <Header></Header>
            <div className='login-wrapper'>
                <h2>Please login using available OAuth2 providers:</h2>
                <div>
                    <button type="submit" onClick={() => onLogin("GitHub")}>GitHub</button>
                </div>
            </div>
        </div>
    )
}