const BACKEND_URL = "https://peminjamanruangan.rtln.xyz"

fetch(`${BACKEND_URL}/`)
  .then((response) => response.json())
  .then((data) => console.log(data));
