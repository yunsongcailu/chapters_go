package pkg_crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func DemoRsa() {
	generateRSAKey(2048)
	fmt.Println("使用公钥对数据加密,使用私钥对密文解密")
	//加密
	data := []byte("hello world")
	encrypt, err := encryptRSA(data, "public.pem")
	
	fmt.Println(string(encrypt), err)
	// 解密
	decrypt, err := decryptRSA(encrypt, "private.pem")
	fmt.Println(string(decrypt), err)
}

// generateRSAKey 根据bits证书大小 生成RSA公钥私钥对
func generateRSAKey(bits int) error {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// 保存私钥
	// 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	// 使用pem格式对x509输出的内容进行编码
	// 创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)
	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
	return nil
}

// encryptRSA RSA加密数据
// plainText 要加密的数据
// path 公钥匙文件地址
func encryptRSA(plainText []byte, path string) ([]byte, error) {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	//返回密文
	return cipherText, nil
}

// decryptRSA RSA解密数据
// cipherText 需要解密的byte数据
// path 私钥文件路径
func decryptRSA(cipherText []byte, path string) ([]byte, error) {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText, nil
}
