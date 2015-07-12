/*

H264 decoder/encoder ffmpeg wrapper

	var img *image.YCbCr
	d, err = h264.NewEncoder(640, 480)
	img, err = d.Encode(img)

	d, err = h264.NewDecoder(pps)
	img, err = d.Decode(nal)
*/
package h264

import (
	"reflect"
	"unsafe"

	/*
		#cgo LDFLAGS: -lavformat -lavutil -lavcodec

		#include <libavutil/avutil.h>
		#include <libavformat/avformat.h>

		static void ffmpeg_init() {
			av_register_all();
			av_log_set_level(AV_LOG_DEBUG);
		}
	*/
	"C"
)

func init() {
	C.ffmpeg_init()
}

func fromCPtr(buf unsafe.Pointer, size int) (ret []uint8) {
	hdr := (*reflect.SliceHeader)((unsafe.Pointer(&ret)))
	hdr.Cap = size
	hdr.Len = size
	hdr.Data = uintptr(buf)
	return
}
