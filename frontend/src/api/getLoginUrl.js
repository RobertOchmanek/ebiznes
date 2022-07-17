export async function getLoginUrl(oauthProvider) {

    let loginUrl

    switch(oauthProvider) {
        default:
            loginUrl = await fetch('http://localhost:8080/oauth');
    }

    return loginUrl.json();
}