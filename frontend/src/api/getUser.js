export async function getUser(backendAddress, userToken) {

    const data = await fetch(backendAddress + '/users/' + userToken);
    return data.json();
  }