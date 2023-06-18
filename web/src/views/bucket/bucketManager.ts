import { Bucket } from "../../types/interfaces"
import errorHandle from "../../util/errHandle"
import httprequest from "../../util/http"
import successHandle from "../../util/successHandle"





export default class BucketManager {
    createBucket(bucket: Bucket) {
        console.log(typeof (bucket.totalNumber))
        httprequest("/bucket/create", {
            method: "POST",
            body: {
                name: bucket.name,
                status: "inuse",
                totalNumber: bucket.totalNumber,
                totalSize: bucket.totalSize,
                capacityMode: bucket.capacityMode
            }
        }).then(data => {
            if (errorHandle(data, "")) {
                return
            }
            successHandle("创建成功")
            // listBucket()
        })
    }

}



