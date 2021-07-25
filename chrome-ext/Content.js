var searchResult = document.querySelectorAll('div[class=g]')
var resultObj = []

for (let i = 0; i < searchResult.length; i++) {
    let title = searchResult[i].querySelector('h3')
    let desc = searchResult[i].querySelector('em')
    let cite = searchResult[i].querySelector('cite')

    if (!title || !desc || !cite) {
        continue
    }
    
    title = title.innerText
    desc = desc.closest('span').innerText
    cite = cite.innerText

    let url = searchResult[i].querySelector('a').href
    let urlVal = (new URL(url))
    let domain = urlVal.hostname;

    resultObj.push({
        "title": title,
        "desc": desc,
        "cite": cite,
        "url": url,
        "domain": domain,
    })
}

let xhttp = new XMLHttpRequest();
xhttp.onreadystatechange = function() {
    if (this.readyState == 4) {
        var res = this.response;
        if(this.status == 200) {
            console.log(res)
        } else {
            console.log("ERROR: ", res);
        }
    }
};
let url = 'http://localhost:8080/search'
xhttp.responseType = 'json';
xhttp.open("POST", url, true);
xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
var payload = JSON.stringify(resultObj);
xhttp.send(payload);
