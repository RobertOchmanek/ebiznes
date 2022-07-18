export async function getLogoutUrl(userId) {

    const logoutUrl = await fetch('http://localhost:8080/oauth/logout/' + userId);
    return logoutUrl.json();
}