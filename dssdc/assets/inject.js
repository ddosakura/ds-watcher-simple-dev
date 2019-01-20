const jklsdvjeklvkmsldlvs = window.onload
window.onload = function () {
    jklsdvjeklvkmsldlvs && jklsdvjeklvkmsldlvs()
    function connWS() {
        const url = "ws://" + window.location.host + "/fresh"
        const ws = new WebSocket(url)
        ws.onopen = function () {
            console.log('Monitor prepared!')
        }
        var jumpCloseAlert = false
        ws.onclose = function (e) {
            if (jumpCloseAlert) return
            console.error('Monitor Closed!')
            setTimeout(connWS, 3000)
        }
        ws.onmessage = function (e) {
            jumpCloseAlert = true
            console.log(e)
            ws.close()
            window.location.reload()
        }
    }
    connWS()
}