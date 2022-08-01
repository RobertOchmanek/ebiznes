export async function getLogoutUrl(backendAddress, userId) {

    const logoutUrl = await fetch(backendAddress + '/oauth/logout/' + userId);
    return logoutUrl.json();
}