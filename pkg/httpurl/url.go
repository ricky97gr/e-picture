package httpurl

func GeneratePictureUrl(endpoint, bucket, fileName string, ssl bool) string {
	if ssl {
		return "https://" + endpoint + "/" + bucket + "/" + fileName
	}
	return "http://" + endpoint + "/" + bucket + "/" + fileName
}
