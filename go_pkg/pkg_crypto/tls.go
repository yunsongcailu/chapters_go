package pkg_crypto

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

func DemoTls() {
	fmt.Println("tls包实现了传输层TLS 1.2")
	testDemoTls()
	fmt.Println("模拟拨号连接 TLS")
	testDemoDialTls()
}

// 模拟TLS 服务 / 客户端
func testDemoTls() {
	// 通过解密网络流量捕获来调试TLS应用程序。

	// 警告：使用KeyLogWriter会危及安全性，并且应该只是
	// 用于调试。
	// 虚假测试HTTP服务器的示例具有不安全的随机输出
	// 重复性。
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.TLS = &tls.Config{
		//Rand: zeroSource{}, // 仅举例来说; 不要这样做。
	}
	server.StartTLS()
	defer server.Close()

	// 通常，日志将转到打开的文件：
	// w, err := os.OpenFile("tls-secrets.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	w := os.Stdout

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				KeyLogWriter: w,

				//Rand:               zeroSource{}, // 用于可重复的输出; 不要这样做。
				InsecureSkipVerify: true, // 测试服务器证书不受信任。
			},
		},
	}
	resp, err := client.Get(server.URL)
	if err != nil {
		log.Fatalf("Failed to get URL: %v", err)
	}
	resp.Body.Close()

	// 生成的文件可以与Wireshark一起使用来解密TLS
	// 通过在SSL协议中设置（Pre）-Master-Secret日志文件名来连接
	// 优先级。
}

// 模拟拨号连接 TLS
func testDemoDialTls() {
	// 使用自定义根证书集连接。

	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----`

	// 首先，创建一组根证书。 对于这个例子我们只
	// 有一个。 也可以省略这个以便使用
	// 当前操作系统的默认根集。
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	conn, err := tls.Dial("tcp", "mail.163.com:443", &tls.Config{
		RootCAs: roots,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	conn.Close()
}

/*
type Config struct {
	// Rand为nonce和RSA致盲提供了熵的来源。
	// 如果Rand为零，则TLS在包中使用加密随机读取器
	// crypto/rand。
	// 阅读器必须安全使用多个goroutines。
	Rand io.Reader

	// 时间将当前时间作为自纪元以来的秒数返回。
	// 如果Time为nil，则TLS使用time.Now。
	Time func() time.Time

	// 证书包含一个或多个要呈现的证书链
	// 连接的另一面。 服务器配置必须包含
	// 至少一个证书或设置GetCertificate。 客户端正在
	// 客户端身份验证可以设置证书或
	// GetClientCertificate。
	Certificates []Certificate

	// NameToCertificate从证书名称映射到元素
	// 证书。 请注意，证书名称可以是表单
	// '* .example.com'因此不必是域名。
	// 请参见Config.BuildNameToCertificate
	// nil值导致使用证书的第一个元素
	// 对于所有连接。
	NameToCertificate map[string]*Certificate

	// GetCertificate根据给定的值返回证书
	// ClientHelloInfo。 只有在客户提供SNI时才会调用
	// 信息或证书是否为空。
	//
	// 如果GetCertificate为nil或返回nil，则证书为
	// 从NameToCertificate检索。 如果NameToCertificate是nil，那么
	// 将使用证书的第一个元素。
	GetCertificate func(*ClientHelloInfo) (*Certificate, error)

	// GetClientCertificate，如果不是nil，则在服务器请求时调用
	// 来自客户的证书。 如果设置，证书的内容将
	// 被忽略。
	//
	// 如果GetClientCertificate返回错误，则握手（handshake）将是
	// 中止，将返回该错误。 除此以外
	// GetClientCertificate必须返回非零证书。 如果
	// Certificate.Certificate为空，则不会发送任何证书
	// 服务器。 如果这对服务器来说是不可接受的，那么它可能会中止
	// 握手。
	//
	// 可以多次调用GetClientCertificate
	// 如果发生重新协商或者正在使用TLS 1.3，则连接。
	GetClientCertificate func(*CertificateRequestInfo) (*Certificate, error)

	// GetConfigForClient，如果不是nil，则在ClientHello之后调用
	// 从客户端那里收到 它可能会返回一个非零配置
	// 更改将用于处理此连接的Config。 如果
	// 返回的Config为nil，将使用原始配置。该
	// 此回调返回的配置可能不会随后被修改。
	//
	// 如果GetConfigForClient为nil，则传递给Server()的Config将为
	// 用于所有连接。
	//
	// 对于返回的Config中的字段，会话票证密钥是唯一的
	// 如果未设置，将从原始配置中复制。
	// 具体来说，如果在原始上调用了SetSessionTicketKeys
	// 配置但不在返回的配置上，然后从票证键
	// 原始配置将在使用前复制到新配置中。
	// 否则，如果在原始配置中设置了SessionTicketKey但是
	// 不在返回的配置中然后它将被复制到返回的
	// 使用前配置。 如果这两种情况都不适用那么关键
	// 返回的配置中的材料将用于会话票证。
	GetConfigForClient func(*ClientHelloInfo) (*Config, error)

	// VerifyPeerCertificate，如果不是nil，则在正常情况下被调用
	// 由TLS客户端或服务器验证证书。 它
	// 接收对等方提供的原始ASN.1证书
	// 正常处理发现的任何经过验证的链。 如果它返回一个
	// 非零错误，中止握手并产生错误。
	//
	// 如果正常验证失败，那么握手将在之前中止
	// 考虑这个回调。 如果禁用正常验证
	// 设置InsecureSkipVerify然后会考虑这个回调但是
	// verifiedChains参数将始终为零。
	VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

	// RootCAs定义了根证书授权的集合
	// 客户端在验证服务器证书时使用。
	// 如果RootCAs为nil，则TLS使用主机的根CA集。
	RootCAs *x509.CertPool

	// NextProtos是受支持的应用程序级协议列表。
	NextProtos []string

	// ServerName用于验证返回的主机名
	// 证书，除非给出InsecureSkipVerify。 它也包括在内
	// 在客户端的握手中支持虚拟主机，除非它是
	// 一个IP地址。
	ServerName string

	// ClientAuth确定服务器的策略
	// TLS客户端身份验证。 默认值为NoClientCert。
	ClientAuth ClientAuthType

	// ClientCAs定义了根证书颁发机构的集合
	// 服务器使用，如果需要验证客户端证书
	// 通过ClientAuth中的策略。
	ClientCAs *x509.CertPool

	// InsecureSkipVerify控制客户端是否验证
	// 服务器的证书链和主机名。
	// 如果InsecureSkipVerify为true，则TLS接受任何证书
	// 由服务器和该证书中的任何主机名提供。
	// 在这种模式下，TLS容易受到man-in-the-middle攻击。
	// 这应该仅用于测试。
	InsecureSkipVerify bool

	// CipherSuites是受支持的密码套件列表。 如果是CipherSuites
	// 是零，TLS使用实现支持的套件列表。
	CipherSuites []uint16

	// PreferServerCipherSuites控制服务器是否选择
	// 客户端最喜欢的密码套件，或服务器最优选的密码套件
	// 密码套件。 如果为true则为服务器的首选项，如表所示
	// 使用CipherSuites中元素的顺序。
	PreferServerCipherSuites bool

	// 可以将SessionTicketsDisabled设置为true以禁用会话票证
	// （恢复）支持。
	SessionTicketsDisabled bool

	// TLS服务器使用SessionTicketKey来提供会话
	// 恢复。 请参阅RFC 5077.如果为零，则将填充
	// 第一次服务器握手之前的随机数据。
	//
	// 如果多个服务器正在终止同一主机的连接
	// 他们都应该拥有相同的SessionTicketKey。 如果
	// SessionTicketKey泄漏，以前记录和未来的TLS
	// 使用该密钥的连接受到损害。
	SessionTicketKey [32]byte

	// SessionCache是TLS会话的ClientSessionState条目的缓存
	// 恢复。
	ClientSessionCache ClientSessionCache

	// MinVersion包含可接受的最小SSL/TLS版本。
	// 如果为零，则将TLS 1.0作为最小值。
	MinVersion uint16

	// MaxVersion包含可接受的最大SSL/TLS版本。
	// 如果为零，则使用此程序包支持的最大版本，
	// 目前是TLS 1.2。
	MaxVersion uint16

	// CurvePreferences包含将在其中使用的椭圆曲线
	// ECDHE握手，按优先顺序排列。 如果为空，则默认为
	// 被使用。
	CurvePreferences []CurveID

	// DynamicRecordSizingDisabled禁用TLS记录的自适应大小调整。
	// 如果为true，则始终使用最大可能的TLS记录大小。 当
	// false时，可以调整TLS记录的大小
	// 改善延迟。
	DynamicRecordSizingDisabled bool

	// 重新协商控制支持哪种类型的重新协商。
	// 对于绝大多数应用程序，默认值none都是正确的。
	Renegotiation RenegotiationSupport

	// KeyLogWriter可选择指定TLS主机密的目标
	// 在NSS密钥日志格式中，可用于允许外部程序
	// 比如Wireshark来解密TLS连接。
	// 请参阅https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format。
	// 使用KeyLogWriter会危及安全性，应该只是
	// 用于调试。
	KeyLogWriter io.Writer
	// 包含已过滤或未导出的字段
}
*/
