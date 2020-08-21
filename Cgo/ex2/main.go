package main

/*
#include <string.h>
#include <stdio.h>

void go_gmssl_set_cert2(void *cert_file, int cert_file_len, void *cert_file_enc, int cert_file_len_enc, void *ca_file, int ca_file_len, void *password, int password_len) {
        char cert_file_str[1024] = {0,};
        char cert_file_str_enc[1024] = {0,};
        char ca_file_str[1024] = {0,};
        char password_str[1024] = {0,};

        memcpy(cert_file_str, cert_file, cert_file_len >= 1024 ? 1023 : cert_file_len);
        memcpy(cert_file_str_enc, cert_file_enc, cert_file_len_enc >= 1024 ? 1023 : cert_file_len_enc);
        memcpy(ca_file_str, ca_file, ca_file_len >= 1024 ? 1023 : ca_file_len);
        memcpy(password_str, password, password_len >= 1024 ? 1023 : password_len);

		printf(cert_file_str);
		printf(cert_file_str_enc);
		printf(ca_file_str);
		printf(password_str);
}
*/
import "C"
import (
	"unsafe"
)

func main() {
	Xgs_api_set_cert2("1", "22", "333", "4444")
}

func Xgs_api_set_cert2(ca_file string, cert_file string, cert_file_enc string, cert_password string) {
	C.go_gmssl_set_cert2(unsafe.Pointer(&[]byte(cert_file)[0]), C.int(len(cert_file)),
		unsafe.Pointer(&[]byte(cert_file_enc)[0]), C.int(len(cert_file_enc)),
		unsafe.Pointer(&[]byte(ca_file)[0]), C.int(len(ca_file)),
		unsafe.Pointer(&[]byte(cert_password)[0]), C.int(len(cert_password)))

	//fmt.Println(unsafe.Pointer(&[]byte(ca_file)[0]), C.int(len(ca_file)))
	//fmt.Println(unsafe.Pointer(&[]byte(cert_file)[0]), C.int(len(cert_file)))
	//fmt.Println(unsafe.Pointer(&[]byte(cert_file_enc)[0]), C.int(len(cert_file_enc)))
	//fmt.Println(unsafe.Pointer(&[]byte(cert_password)[0]), C.int(len(cert_password)))
}
