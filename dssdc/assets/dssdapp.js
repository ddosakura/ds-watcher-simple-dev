function GetRequest() {
    var url = location.search
    var theRequest = new Object()
    if (url.indexOf("?") != -1) {
        var str = url.substr(1)
        strs = str.split("&")
        for (var i = 0; i < strs.length; i++) {
            theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1])
        }
    }
    return theRequest
}
const params = GetRequest()

// TODO: remove jquery and redo ui
const app = $("#app")

app.append("<h1>DDoSakura Simple Dev</h1>")

app.append("<h3>EntryPoint:</h3>")
const ep = $("<div></div>")
app.append(ep)

app.append("<h3>Developers:</h3>")
const dev = $("<div></div>")
app.append(dev)

app.append("<h3>Detail:</h3>")
let dl = $("<div>No Choose</div>")
app.append(dl)

function getData(url, el, f) {
    $.getJSON(url, (e) => {
        if (e.code === 0) {
            f(e.data, el)
        } else {
            el.append(e.msg)
        }
    })
}

getData("/entrypoint.action", ep, (data, el) => {
    let i = 0
    for (let url in data) {
        const path = data[url]
        el.append(`<span>${i++}. </span><a href='${url}'>${path}</a><br/>`)
    }
})


getData("/developers.action", dev, (data, el) => {
    for (let i in data) {
        const name = data[i]
        el.append(`<span>${i}. </span><button onclick='getDetail("${name}")'>${name}</button><br/>`)
    }
})


const interval = parseInt(params["interval"]) || 5

function getDetail(name) {
    window.location.search = `?dev=${name}&interval=${interval}`
}

(function (name) {
    if (!name) return
    getData(`/detail.action?name=${name}`, dl, (data, el) => {
        const work = []
        let timer = 0
        for (let i in data) {
            const d = data[i]
            const t = Date.parse(d.ChangeTime)
            if (work[timer]) {
                if (t - work[timer].toTime > 1000 * 60 * interval) {
                    timer++
                    work[timer] = {
                        fromTime: t,
                        toTime: t,
                        saved: [d.File],
                    }
                } else {
                    work[timer].toTime = t
                    work[timer].saved.push(d.File)
                }
            } else {
                work[timer] = {
                    fromTime: t,
                    toTime: t,
                    saved: [d.File],
                }
            }
        }
        // console.log(work)

        const dvs = $("<div></div>")
        for (let i in work) {
            const d = work[i]
            const dv = $("<div></div>")
            dv.append(`<span>ID-${formatInt(i)}. </span>`)
            dv.append(`<span>Work from ${formatDateTime(d.fromTime)} to ${formatDateTime(d.toTime)} continued ${formatTime(d.toTime-d.fromTime)} saved ${d.saved.length} times</span>`)
            dvs.append(dv)
        }
        el.replaceWith(dvs)
        dl = dvs
    })
})(params["dev"])

// TODO: 临时功能
function formatInt(n, l = 4) {
    n = "" + n
    while (n.length < l) {
        n = "0" + n
    }
    return n
}

function formatDateTime(da) {
    da = new Date(da);
    const year = da.getFullYear()
    const month = da.getMonth() + 1
    const date = da.getDate()
    const p1 = [year, month, date].join('-')
    const h = da.getHours()
    const m = da.getMinutes()
    const s = da.getSeconds()
    const p2 = [h, m, s].join(':')
    return p1 + " " + p2
}

function formatTime(t) {
    // t += 60 * 1000
    const h = formatInt(Math.floor(t / 3600 / 1000), 2)
    const m = formatInt(Math.floor(t / 60 / 1000) % 60, 2)
    const s = formatInt(Math.floor(t / 1000) % 60, 2)
    return [h, m, s].join(':')
}