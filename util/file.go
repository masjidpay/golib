package util

func ValidateImage(ext string) bool {

	if ext == ".jpg" ||
		ext == ".jpeg" ||
		ext == ".png" ||
		ext == ".JPG" ||
		ext == ".JPEG" ||
		ext == ".PNG" {
		return true
	}
	return false

}
