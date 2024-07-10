import axios from 'axios'

axios
  .get('http://localhost:8080/sneakers')
  .then((response) => {
    console.log(response.data.length)
  })
  .catch((error) => {
    console.log('ERRRR' + error.message)
  })
