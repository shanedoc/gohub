package mail

type Driver interface {
	//检查验证码
	Send(email Email)
}
