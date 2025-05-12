const form = document.getElementById('search-form')
form.addEventListener('submit', async (event) => {
    event.preventDefault()
    const formData = new FormData(event.target)
    const query = formData.get('query')

    const newUrl = `?query=${encodeURIComponent(query)}`
    history.pushState(null, '', newUrl)

    const data = await search(query)
    updateDisplay(data)
})

async function search(query) {
    const response = await fetch("http://127.0.0.1:8080/?query="+query)
    const data = await response.json()
    return data;
}


function updateDisplay(data){
    const display = document.getElementById('results');
    console.log(data)

    const items = Array.isArray(data.data) ? data.data : [];

    items.forEach(item => {
        const page = document.createElement("div");
        const pageTitle = document.createElement("h3");
        const pageURL = document.createElement("span");
        const pageDesc = document.createElement("p")
        pageTitle.textContent = item.title;
        pageURL.textContent = item.url;
        pageDesc.textContent = (item.description).substr(0, 100) + "..."

        page.appendChild(pageTitle);
        page.appendChild(pageURL);
        page.appendChild(pageDesc);
        display.appendChild(page);
    });
}