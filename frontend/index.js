// const BACKEND_URL = "https://peminjamanruangan.rtln.xyz"
const BACKEND_URL = "http://localhost:1118"

const btnRegister = $("#btn-register")
const txtUsername = $("#txt-username").val()
const txtPassword = $("#txt-password").val()
const txtRealName = $("#txt-realname").val()

// fetch(`${BACKEND_URL}/`)
//   .then((Response) => Response.json())
//   .then((data) => console.log(data));

// Example POST method implementation:
async function postData(url = '', data = {}) {
  // Default options are marked with *
  const response = await fetch(url, {
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, *cors, same-origin
    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    credentials: 'same-origin', // include, *same-origin, omit
    headers: {
      'Content-Type': 'application/json'
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    // redirect: 'follow', // manual, *follow, error
    // referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: JSON.stringify(data) // body data type must match "Content-Type" header
  });
  return response.json(); // parses JSON response into native JavaScript objects
}

$("document").ready(() => {
  console.log("script loaded")
})

btnRegister.submit(postData(`${BACKEND_URL}/user/register`, {
  "username": txtUsername,
  "password": txtPassword,
  "real_name": txtRealName,
})
  .then((data) => {
    console.log('User register endpoint called')
    console.log(data); // JSON data parsed by `data.json()` call
  }))

  