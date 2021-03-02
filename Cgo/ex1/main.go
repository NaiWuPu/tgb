package main

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

extern int cli_send_buf_to_mgd(char **outbuf);
#define sg_ipc_free(d) __sg_ipc_free(__FILE__, __LINE__, d)

static int logtrans_receive_debug(char *outbuf )
{
	int ret = 0;
	char *refund=NULL;
	ret = cli_send_buf_to_mgd(&refund);

	memcpy(outbuf, refund, 4);
	printf("444 %s \n", outbuf);
	return ret;
}

int cli_send_buf_to_mgd(char **outbuf)
{
	int ret = 0;
	char *rcv_msg;

	rcv_msg = "123";
	*outbuf = malloc(4);
	memcpy(*outbuf, rcv_msg, 4);
	printf("333 %s \n", *outbuf);

	if (rcv_msg) {
		rcv_msg = NULL;
	}

	return ret;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	var r string
	fmt.Printf("11 %p \n", &r)

	rs := C.CString(r)
	fmt.Printf("22 %p \n", &rs)
	retVal := C.logtrans_receive_debug(rs)

	fmt.Println(retVal)
	fmt.Println(C.GoString(rs))
}
