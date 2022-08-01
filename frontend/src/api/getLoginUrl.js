export async function getLoginUrl(backendAddress) {

    const loginUrl = await fetch(backendAddress + '/oauth/login');
    return loginUrl.json();
}