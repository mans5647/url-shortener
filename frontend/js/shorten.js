const API_ENDPOINT = "http://localhost:8080";

async function shortenUrl() {
    
    let url = getLongUrl()

    if (url !== undefined) {

        let result = await getShortUrl(url)
        let shortUrl = createAliasUrl(result)
        setAliasUrl(shortUrl)

    } else {
        alert("Невалидный URL")
    }
}

async function getShortUrl(longUrl) {

    const data = {
        url: longUrl
    }

    let resp = await fetch(API_ENDPOINT + '/api/v1/url', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Content-Type': 'application/json'
        }
    }
    )

    return await resp.json()
}

function createAliasUrl(dict) {
    return window.location.host + "/" + dict['code'];
}

function setAliasUrl(alias) {
    document.getElementById("alias-container")
            .innerText = `Вот сокращенная ссылка: ${alias}`;
}

function getLongUrl() {
    let url = document.getElementById("url-input").value; 

    // Validate 
    if(!isValidUrl(url)) { 
        return undefined; 
    }

    return url;
}

function isValidUrl(url) {

    const re = new RegExp(/(?:http[s]?:\/\/.)?(?:www\.)?[-a-zA-Z0-9@%._\+~#=]{2,256}\.[a-z]{2,6}\b(?:[-a-zA-Z0-9@:%_\+.~#?&\/\/=]*)/gm);
    return re.exec(url) !== null
}