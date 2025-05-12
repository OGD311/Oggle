
const form = document.getElementById('search-form')
form.addEventListener('submit', async (event) => {
    event.preventDefault()
    const formData = new FormData(event.target)
    const query = formData.get('query')

    const data = await search(query)
    console.log(data)
})

async function search(query) {
    const response = await fetch("http://127.0.0.1:8080/?query="+query)
    const data = await response.json()
    return data;
}
