export async function getUser(userToken) {

    const data = await fetch('http://localhost:8080/users/' + userToken);
    return data.json();
  }