package main

// #include "app_Main.h"
// #include <stdlib.h>
// static jstring JavaStringFromCString(JNIEnv *env, const char *utf8) {
//   return (*env)->NewStringUTF(env, utf8);
// }
import "C"
import (
	"encoding/json"
	"unsafe"
)

//export Java_app_Main_nativeFetch
func Java_app_Main_nativeFetch(env *C.JNIEnv, _ C.jclass) C.jstring {
	result, err := fetch()
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return JavaString(env, string(data))
}

func JavaString(env *C.JNIEnv, str string) C.jstring {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return C.JavaStringFromCString(env, cstr)
}

func main() {}
