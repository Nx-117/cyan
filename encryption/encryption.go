package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"hash"
	"io"
	"os"
)

/**
加密算法相关
*/

/**
MD5
str string 需要加密的字符串
toUpper bool 返回类型 true大写 false小写
*/
func Md5String(str string, toUpper bool) string {
	if toUpper {
		return fmt.Sprintf("%X", md5.Sum([]byte(str)))
	} else {
		return fmt.Sprintf("%x", md5.Sum([]byte(str)))
	}
}

const (
	MD5    = "md5"
	SHA1   = "sha1"
	SHA256 = "sha256"
	SHA512 = "sha512"
)

/**
md5、sha1、sha256、sha512加密算法
src  原始待加密数据
CryptTool  加密类型
isHex 数据是否为16进制串，true表示16进制串，需要进行解析，false表示非16进制串
count int 加密次数
*/
func Hash(src, CryptTool string, isHex bool, count int) string {
	var hash hash.Hash

	switch CryptTool {
	case "md5":
		hash = md5.New()
	case "sha1":
		hash = sha1.New()
	case "sha256":
		hash = sha256.New()
	case "sha512":
		hash = sha512.New()
	}

	if isHex {
		//如果加密串本身是16进制串需要做解析
		csrc, _ := hex.DecodeString(src)
		hash.Write(csrc)
	} else {
		hash.Write([]byte(src))
	}

	//一次加密后的结果
	cryptStr := fmt.Sprintf("%x", hash.Sum(nil))

	for i := 0; i < count; i++ {
		hash.Reset()
		hash.Write([]byte(cryptStr))
		cryptStr = fmt.Sprintf("%x", hash.Sum(nil))
	}

	return cryptStr
}

/**
aes加密算法：秘钥长度可指定（16B、24B、32B)
src 原始待加密数据
*/
func AESEncrypter(src, key string) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length is illegal!")
	}

	//分组块
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	//秘钥长度
	blockSize := block.BlockSize()
	nsrc := padding([]byte(src), blockSize)
	//加密模式
	blockMode := cipher.NewCBCEncrypter(block, []byte(key)[:blockSize])
	dst := make([]byte, len(nsrc))
	//加密
	blockMode.CryptBlocks(dst, nsrc)

	return base64.StdEncoding.EncodeToString(dst), nil
}

/**
aes解密算法：秘钥长度可指定（16B、24B、32B)
src 待解密数据
key 密钥
*/
func AESDecrypter(src, key string) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length is illegal!")
	}

	nsrc, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	//秘钥长度
	blockSize := block.BlockSize()
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(key)[:blockSize])
	dst := make([]byte, len(nsrc))
	//解密
	blockMode.CryptBlocks(dst, nsrc)
	//去掉填充内容
	dst = unPadding(dst)
	return string(dst), nil
}

/**
生成公钥、私钥
keyPath 密钥生成路径 必须以 / 结尾
*/
func GenerateKey(keyPath string) error {
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}

	//序列化私钥
	priKey := x509.MarshalPKCS1PrivateKey(privateKey)

	//将私钥转成pem格式
	priBlock := &pem.Block{
		Type:  "RSA Private Key",
		Bytes: priKey,
	}

	//生成存储私钥的pem文件
	priKeyFp, err := os.Create(keyPath + "privateKey.pem")
	if err != nil {
		return err
	}

	defer priKeyFp.Close()

	//pem格式数据写入文件
	err = pem.Encode(priKeyFp, priBlock)
	if err != nil {
		return err
	}

	//生成公钥
	publicKey := &privateKey.PublicKey

	//序列化公钥
	pubKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	//将公钥转成pem格式
	pubBlock := &pem.Block{
		Type:  "RSA Public Key",
		Bytes: pubKey,
	}

	//生成存储公钥的pem文件
	pubKeyFp, err := os.Create(keyPath + "publicKey.pem")
	if err != nil {
		return err
	}

	defer pubKeyFp.Close()

	//pem格式数据写入文件
	err = pem.Encode(pubKeyFp, pubBlock)
	if err != nil {
		return err
	}
	return nil
}

/**
RSA加密算法：公钥加密
@param src [string] 待加密数据串
@return [string] RSA加密数据串
*/
func RSAEncrypter(src, publicKeyStr string) (string, error) {

	//解析pem格式公钥
	pubBlock, _ := pem.Decode([]byte(publicKeyStr))
	if pubBlock == nil {
		panic("pem格式公钥解析失败")
	}
	//解析公钥
	pub, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return "", err
	}
	//类型断言，转化类型
	publicKey := pub.(*rsa.PublicKey)
	//公钥加密
	nsrc, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(src))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(nsrc), nil
}

/**
RSA解密算法：私钥解密
@param src 待解密数据串
privateKeyStr  私钥字符串
*/
func RSADecrypter(src, privateKeyStr string) (string, error) {

	//解析pem格式私钥
	priBlock, _ := pem.Decode([]byte(privateKeyStr))
	//priBlock, _ := pem.Decode([]byte(priKeyPem))
	//解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		return "", err
	}
	nsrc, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	//私钥解密
	nsrc, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, nsrc)
	if err != nil {
		return "", err
	}
	return string(nsrc), nil
}

/**
去掉填充串
*/
func unPadding(src []byte) []byte {
	//填充长度
	length := int(src[len(src)-1])
	src = src[:(len(src) - length)]

	return src
}

/**
filePath 证书文件路径
密钥证书转换成字符串
*/
func GetKeyStringByFile(filePath string) (string, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	tmp := make([]byte, 1024)
	//pem格式私钥
	var priKeyPem string
	for {
		n, err := fp.Read(tmp)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		if n == 0 {
			break
		}
		priKeyPem += string(tmp[:n])
	}
	return priKeyPem, nil
}

/**
填充原串
*/
func padding(src []byte, blockSize int) []byte {
	//需要补齐填充的长度
	length := blockSize - len(src)%blockSize
	//用来补齐填充的字符切片，里面的内容清一色是补齐长度的byte类型值
	paddingStr := bytes.Repeat([]byte{byte(length)}, length)
	nsrc := append(src, paddingStr...)

	return nsrc
}
