import Qs from "qs"
//const { objectToString } = require("@vue/shared");

let baseUrl = 'api',

    inital = {
        method: "GET",
        params: null,
        body: null,
        headers: {
            "Content-Type": "application/json",
        },
        credentials: true,
        responseType: "JSON",
        cache: "no-cache"
    };

//校验是否是纯粹的对象
const isPlainObject = function isPlainObject(obj: any) {
    var proto, Ctor;
    if (!obj || typeof obj != "object") {
        return false;
    }
    proto = Object.getPrototypeOf(obj);
    if (!proto) {
        return true;
    }
    Ctor = proto.hasOwnProperty("constructor") && proto.constructor;
    return typeof Ctor == "function" && Ctor == Object;
}

const httprequest = function sendFetch(url: any, config: any, isUpLoadFile?: boolean) {
    (config == null || typeof config != "object") ? config = {} : null;
    if (config.headers && isPlainObject(config.headers)) {
        config.headers = Object.assign({}, inital.headers, config.headers);
    }
    let {
        method,
        params,
        body,
        headers,
        credentials,
        responseType,
        cache
    } = Object.assign({}, inital, config)

    if (typeof url != "string") throw new TypeError(`${url} is not a string!`);
    url = baseUrl + url;
    if (params != null) {
        if (isPlainObject(params)) {
            params = Qs.stringify(params);
        }
        url += `${url.includes('?') ? '$' : '?'}${params}`;
    }
    if (body != null) {
        if (isPlainObject(body)) {
            let contentType = headers["Content-Type"] || "application/json";
            if (contentType.includes("urlencoded")) {
                body = Qs.stringify(body);
            }

            if (contentType.includes("json")) {
                body = JSON.stringify(body);
            }
        }
    }
    credentials = credentials ? 'include' : 'same-origin';
    //TODO: need to fix
    if (isUpLoadFile) {
        headers = {}
    }
    if (sessionStorage.getItem("token")) {
        headers["token"] = sessionStorage.getItem("token");
    }
    // 基于fetch请求数据
    method = method.toUpperCase();
    responseType = responseType.toUpperCase();
    config = {
        method,
        credentials,
        cache,
        headers
    };

    /^(POST|PUT|PATCH|DELETE)$/i.test(method) ? config.body = body : null;
    return fetch(url, config).then(function onfulfilled(response) {
        // 走到这边不一定是成功的：
        // Fetch的特点的是，只要服务器有返回结果，不论状态码是多少，它都认为是成功
        let {
            status,
            statusText
        } = response;
        if (status >= 200 && status < 400) {
            // 真正成功获取数据
            let result;
            switch (responseType) {
                case 'TEXT':
                    result = response.text();
                    break;
                case 'JSON':
                    result = response.json();
                    break;
                case 'BLOB':
                    result = response.blob();
                    break;
                case 'ARRAYBUFFER':
                    result = response.arrayBuffer();
                    break;
            }
            return result;
        }
        // 应该是失败的处理
        return Promise.reject({
            code: 'STATUS ERROR',
            status,
            statusText
        });
    }).catch(function onrejected(reason) {
        // @1:状态码失败
        if (reason && reason.code === "STATUS ERROR") {
            switch (reason.status) {
                case 401:
                    break;
                // ...
            }
        }

        // @2:断网
        if (!navigator.onLine) {
            // ...
        }

        // @3:处理返回数据格式失败
        // ...

        return Promise.reject(reason);
    });
}
export default httprequest;