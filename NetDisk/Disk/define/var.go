package define

import "regexp"

var (
	Upattern, _ = regexp.Compile("^[a-zA-Z0-9]{6,20}$")
	Ppattern, _ = regexp.Compile("^[a-zA-Z0-9!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]{6,20}$")
)

var (
	DocSuffix = map[string]bool{
		".doc": true, ".docx": true,
		".xlsx": true, ".xls": true,
		".ppt": true, " pptx": true,
		".pdf": true,
		".txt": true,
	}

	ImageSuffix = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".tif":  true,
		".tiff": true,
		".webp": true,
	}

	AudioSuffix = map[string]bool{
		".mp3":  true,
		".wav":  true,
		".aac":  true,
		".flac": true,
		".ogg":  true,
		".wma":  true,
		".m4a":  true,
		".amr":  true,
		".ape":  true,
	}

	VideoSuffix = map[string]bool{
		".avi":  true,
		".rm":   true,
		".rmvb": true,
		".mpeg": true,
		".mpg":  true,
		".mp4":  true,
		".wmv":  true,
		".mov":  true,
		".flv":  true,
		".mkv":  true,
		".3gp":  true,
		".m4v":  true,
	}
)
