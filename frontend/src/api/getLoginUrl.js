export async function getLoginUrl() {

    const loginUrl = await fetch('http://localhost:8080/oauth');
    return loginUrl.json();
}