package main

import (
	"archive/zip"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	downLoad(w, r)
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "cert/cacert.pem"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err: ", err)
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8088",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:          pool,
			ClientAuth:         tls.RequireAndVerifyClientCert,
			InsecureSkipVerify: true,
		},
	}
	fmt.Println("listen...")

	//http.HandleFunc("/", downLoad)
	//err = http.ListenAndServe(":8088", nil)
	err = s.ListenAndServeTLS("cert/server.pem", "cert/server.key")

	if err != nil {
		fmt.Println(err)
	}
}

func downLoad(rw http.ResponseWriter, r *http.Request) {
	zipName := "111.zip"
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))
	rw.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	err := getZip(rw)
	if err != nil {
		fmt.Println(err)
	}
}

func getZip(w io.Writer) error {
	// 该方法创建一个zip.Writer，用于向zip文件中写入内容，即打包的文件
	// 返回值为一个zip.Writer，最后的zip内容都会写入这个zip.Writer，而最终当然是写入了参数的io.Writer中，也就是我们的http.ResponseWriter中
	zipW := zip.NewWriter(w)
	defer zipW.Close()

	for i := 0; i < 5; i++ {
		// 该方法向zip.Writer中添加一个文件，也就是说向zip文件中添加一个文件
		// 参数字为字符串，会作为写入zip中的文件的文件名
		f, err := zipW.Create(strconv.Itoa(i) + ".txt")
		if err != nil {
			return err
		}
		// 第一个返回值为一个io.Writer，用于向其中，也就是向我们添加到zip的文件中，写入文件内容，
		// 即如_, err = f.Write([]byte(fmt.Sprintf("Hello file %d", i)))代码所示，我们向文件中写入了简单的字符串
		_, err = f.Write([]byte(fmt.Sprintf("Hello file %d", i)))
		if err != nil {
			return err
		}
	}
	return nil
}
