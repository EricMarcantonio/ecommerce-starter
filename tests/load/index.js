import http from 'k6/http';

export let options = {
    vus: 500,
    duration: '300s',
};

export default function () {
    var url = 'http://127.0.0.1:58896'
    var payload = JSON.stringify({
        query: '{products {id, name, takenBy }}'
    })
    var params = {
        headers: {
            'Content-Type': 'application/json'
        }
    }
    http.post(url, payload, params)
}